// Publisher-only offline book audit; Node >=22 and a local Chrome executable are needed.
import { spawn } from 'node:child_process';
import { mkdtemp, readdir, writeFile, rm } from 'node:fs/promises';
import { tmpdir } from 'node:os';
import { join, resolve } from 'node:path';
import { pathToFileURL } from 'node:url';

const directory = resolve(process.argv[2] || '');
const executable = process.env.BOOK_CHROME;
if (!executable || !process.argv[2]) throw new Error('Set BOOK_CHROME and pass the exported fixture directory');
const profile = await mkdtemp(join(tmpdir(), 'book-chrome-'));
const chrome = spawn(executable, [
    '--headless', '--disable-gpu', '--disable-background-networking',
    '--no-first-run', '--no-default-browser-check', '--remote-debugging-port=0',
    `--user-data-dir=${profile}`, 'about:blank',
], { stdio: ['ignore', 'ignore', 'pipe'] });
let socket;
try {
    const endpoint = await new Promise((resolve, reject) => {
        let output = '';
        const timer = setTimeout(() => reject(new Error('Chrome startup timeout')), 20000);
        chrome.on('error', reject);
        chrome.stderr.on('data', (chunk) => {
            output += chunk;
            const match = output.match(/DevTools listening on (ws:\/\/[^\s]+)/);
            if (match) { clearTimeout(timer); resolve(match[1]); }
        });
        chrome.on('exit', (code) => reject(new Error(`Chrome exited ${code}`)));
    });
    socket = new WebSocket(endpoint);
    await new Promise((resolve, reject) => {
        socket.addEventListener('open', resolve, { once: true });
        socket.addEventListener('error', reject, { once: true });
    });
    let sequence = 0;
    const pending = new Map();
    const listeners = new Map();
    socket.addEventListener('message', ({ data }) => {
        const message = JSON.parse(data);
        if (message.id && pending.has(message.id)) {
            const { resolve, reject, timer } = pending.get(message.id);
            clearTimeout(timer);
            pending.delete(message.id);
            if (message.error) reject(new Error(JSON.stringify(message.error)));
            else resolve(message.result);
        }
        if (listeners.has(message.method)) listeners.get(message.method)(message.params);
    });
    const send = (method, params = {}, sessionId) => new Promise((resolve, reject) => {
        const id = ++sequence;
        const timer = setTimeout(() => reject(new Error(`CDP timeout: ${method}`)), 15000);
        pending.set(id, { resolve, reject, timer });
        socket.send(JSON.stringify({ id, method, params, sessionId }));
    });
    const { targetId } = await send('Target.createTarget', { url: 'about:blank' });
    const { sessionId } = await send('Target.attachToTarget', { targetId, flatten: true });
    const call = (method, params) => send(method, params, sessionId);
    await call('Page.enable');
    await call('Network.enable');
    await call('Network.emulateNetworkConditions', {offline:true,latency:0,downloadThroughput:0,uploadThroughput:0});
    const pages=(await readdir(directory)).filter(name=>name.endsWith('.html')).sort();
    const report=[];
    for(const [width,height] of [[320,568],[844,390],[1366,768]]){
        await call('Emulation.setDeviceMetricsOverride',{width,height,deviceScaleFactor:1,mobile:width<900});
        for(const page of pages){
            const loaded=new Promise((resolve,reject)=>{
                const timer=setTimeout(()=>reject(new Error('Page load timeout')),15000);
                listeners.set('Page.loadEventFired',()=>{clearTimeout(timer);resolve()});
            });
            await call('Page.navigate',{url:pathToFileURL(join(directory,page)).href});await loaded;
            const result=await call('Runtime.evaluate',{awaitPromise:true,returnByValue:true,expression:`(async()=>{
                await document.fonts.ready;
                await Promise.all([...document.images].map(i=>i.decode().catch(()=>{})));
                const width=document.documentElement.clientWidth;
                const overflow=document.documentElement.scrollWidth>width+1;
                const brokenImages=[...document.images].filter(i=>!i.naturalWidth).map(i=>i.getAttribute('src'));
                let searchOK=true;
                const query=document.getElementById('query');
                if(query){query.value='шаблоны';query.dispatchEvent(new Event('input'));
                    searchOK=!!document.querySelector('#results a[href="12-templates.html"]');
                    query.value='<script>bad</script>';query.dispatchEvent(new Event('input'));
                    searchOK=searchOK&&!document.querySelector('#results script');
                }
                const outside=[...document.querySelectorAll('p,table,h1,h2,h3')].filter(e=>e.getBoundingClientRect().right>width+1).map(e=>e.textContent.slice(0,100));
                return {overflow,brokenImages,searchOK,outside};
            })()`});
            if(result.exceptionDetails)throw new Error(JSON.stringify(result.exceptionDetails));
            report.push({page,width,height,...result.result.value});
        }
    }
    await writeFile(join(directory,'../html-browser-report.json'),JSON.stringify(report,null,2)+'\n');
    const failures=report.filter(r=>r.overflow||r.brokenImages.length||!r.searchOK||r.outside.length);
    if(failures.length)throw new Error(JSON.stringify(failures));
    console.log(`PASS: ${report.length} offline HTML layouts, all images loaded, search finds chapter 12, no page overflow`);
} finally {
    socket?.close();chrome.kill('SIGTERM');
    await new Promise(resolve=>{if(chrome.exitCode!==null||chrome.signalCode!==null)resolve();else chrome.once('exit',resolve)});
    await rm(profile,{recursive:true,force:true,maxRetries:10,retryDelay:200});
}
