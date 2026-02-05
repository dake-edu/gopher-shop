import{_ as a,o as n,c as i,ae as e}from"./chunks/framework.BD9pLfSD.js";const k=JSON.parse('{"title":"Database Abstraction & Repository Pattern","description":"","frontmatter":{},"headers":[],"relativePath":"database.md","filePath":"database.md"}'),t={name:"database.md"};function l(p,s,o,r,h,c){return n(),i("div",null,[...s[0]||(s[0]=[e(`<h1 id="database-abstraction-repository-pattern" tabindex="-1">Database Abstraction &amp; Repository Pattern <a class="header-anchor" href="#database-abstraction-repository-pattern" aria-label="Permalink to &quot;Database Abstraction &amp; Repository Pattern&quot;">​</a></h1><h2 id="the-power-of-interfaces" tabindex="-1">The Power of Interfaces <a class="header-anchor" href="#the-power-of-interfaces" aria-label="Permalink to &quot;The Power of Interfaces&quot;">​</a></h2><div class="tip custom-block github-alert"><p class="custom-block-title">TIP</p><p>⚓ <strong>Visual Anchor:</strong> The Universal Plug</p></div><p>Think of the <code>BookRepository</code> interface as a wall socket.</p><ul><li>Your Application is the lamp.</li><li>The Electricity Source can be <strong>Memory</strong>, <strong>PostgreSQL</strong>, or <strong>File System</strong>.</li></ul><p>The lamp doesn&#39;t care where the electricity comes from, as long as it fits the socket.</p><div class="language-mermaid vp-adaptive-theme"><button title="Copy Code" class="copy"></button><span class="lang">mermaid</span><pre class="shiki shiki-themes github-light github-dark vp-code" tabindex="0"><code><span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">classDiagram</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    class Application {</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        +Uses Repository</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    }</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    class BookRepository {</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        &lt;&lt;interface&gt;&gt;</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        +GetAll()</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        +GetByID()</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    }</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    class InMemoryStore {</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        +GetAll()</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        +GetByID()</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    }</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    class PostgresStore {</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        +GetAll()</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">        +GetByID()</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    }</span></span>
<span class="line"></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    Application --&gt; BookRepository</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    InMemoryStore ..|&gt; BookRepository</span></span>
<span class="line"><span style="--shiki-light:#24292E;--shiki-dark:#E1E4E8;">    PostgresStore ..|&gt; BookRepository</span></span></code></pre></div><h2 id="configuration" tabindex="-1">Configuration <a class="header-anchor" href="#configuration" aria-label="Permalink to &quot;Configuration&quot;">​</a></h2><p>The application now reads from Environment Variables to connect to the database. Check <code>.env.example</code> and ensure your <code>docker-compose.yml</code> is running.</p>`,9)])])}const d=a(t,[["render",l]]);export{k as __pageData,d as default};
