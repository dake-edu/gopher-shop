#!/usr/bin/env python3
"""Build a clearly labelled review edition in PDF, EPUB and offline HTML.
Requires requirements-build.txt. Run check.py first to resolve source includes.
"""
import os as _os
import sys as _sys
if not _sys.flags.utf8_mode:
    _os.environ['PYTHONUTF8'] = '1'
    _os.execv(_sys.executable, [_sys.executable, '-X', 'utf8', *_sys.argv])

from pathlib import Path
import hashlib
import html
import json
import shutil
import re
import zipfile
import xml.etree.ElementTree as ET
from datetime import datetime, timezone
import markdown
from reportlab.pdfbase import pdfmetrics
from reportlab.pdfbase.ttfonts import TTFont
from reportlab.lib.styles import ParagraphStyle
from reportlab.lib import colors
from reportlab.platypus import SimpleDocTemplate, Paragraph, Spacer, PageBreak, XPreformatted, Table, TableStyle, KeepTogether, Image
from reportlab.lib.enums import TA_LEFT
from reportlab.graphics import renderPDF
from reportlab.platypus import Flowable
from diagrams import render_svg, drawing
from highlighting import PALETTE, highlight_tree, css as highlighting_css
from transitions import render_transitions
from pdf_text import escaped_text
from frontmatter import front_pages
from screenshot_audit import audit as audit_screenshots
from sample_boundary import audit as audit_boundary

ROOT = Path(__file__).resolve().parents[1]
audit_screenshots(ROOT)
audit_boundary(ROOT)
OUT = ROOT/'build/preview'
OUT.mkdir(parents=True, exist_ok=True)
WEB = OUT/'html'
META = json.loads((ROOT/'book.json').read_text())
verification = json.loads((ROOT/'research/verification.json').read_text())
current_sources = {
    file.relative_to(ROOT).as_posix(): hashlib.sha256(file.read_bytes()).hexdigest()
    for folder in ('manuscript', 'examples')
    for file in sorted((ROOT/folder).rglob('*')) if file.is_file()
}
if verification.get('status') != 'passed' or verification.get('source_sha256') != current_sources:
    raise RuntimeError('Sources changed: run book/tools/check.py before building')

TITLE = META['title']
PUBLICATION = META['publication']
COVER_NAME = Path(PUBLICATION['cover']).name
COVER_MIME = {'.jpg':'image/jpeg', '.jpeg':'image/jpeg', '.png':'image/png'}[Path(COVER_NAME).suffix.lower()]
if WEB.exists(): shutil.rmtree(WEB)
WEB.mkdir()

FRONT = front_pages(META)
EDITION = 'Ознакомительный фрагмент '+META['edition']+' · вступление и главы 1–'+str(max(int(c['id']) for c in META['chapters']))
CSS = '''@font-face{font-family:Noto;src:url(fonts/NotoSans-Regular.ttf)}
@font-face{font-family:Noto;src:url(fonts/NotoSans-Bold.ttf);font-weight:700}
@font-face{font-family:NotoMono;src:url(fonts/NotoSansMono-Regular.ttf)}
:root{color-scheme:light}*,*::before,*::after{box-sizing:border-box}body{overflow-wrap:anywhere;font-family:Noto,sans-serif;line-height:1.7;margin:auto;max-width:52rem;padding:1.5rem;color:#182f35;background:#ffffff}
h1,h2,h3{line-height:1.25;color:#00758D}h1{font-size:2rem}h2{margin-top:2rem}
a{color:#00758D;text-underline-offset:.15em}nav{border-bottom:1px solid #b9cdcb;padding-bottom:1rem;margin-bottom:2rem}
pre{background:#ffffff;border:1px solid #DBD9D6;padding:1rem;border-left:3px solid #00ADD8;overflow:auto;line-height:1.5}
code{font-family:NotoMono,monospace;font-size:.88em}pre code{font-size:.8rem}
img{max-width:100%;height:auto}.figure-caption{font-size:.85em;color:#526669;margin-top:.4rem}table{border-collapse:collapse;width:100%;font-size:.9em}td,th{padding:.55rem;border:1px solid #aabfbc;text-align:left;vertical-align:top;overflow-wrap:anywhere}
blockquote{margin-left:0;padding-left:1rem;border-left:3px solid #b08242}.status{color:#64543b;font-size:.9em}.skip{display:block}
input{font:inherit;width:95%;padding:.5rem}li{margin-bottom:.5rem}

@media print{nav,.search,.skip{display:none}body{color:#111;background:white}pre{white-space:pre-wrap}h1,h2{break-after:avoid}}
'''
CSS += highlighting_css()
CSS += """
.front-cover{max-width:42rem;padding:0 1rem;text-align:center}
.front-cover .book-cover{display:block;width:100%;height:auto;margin:0 auto}
.titlepage{text-align:center;max-width:42rem;padding:3rem 1.5rem}
.titlepage main{min-height:46rem;min-height:calc(100svh - 6rem);display:flex;flex-direction:column;gap:3rem}
.titlepage .author{font-size:.725rem;margin:0}
.titlepage .title-block{margin:auto 0}
.titlepage h1{font-size:clamp(1.7rem,5vw,2.5rem)}.titlepage .subtitle{font-size:1.2rem}
.titlepage .levels{margin-top:2rem}.titlepage .imprint{margin:0}
@media(max-height:35rem){.titlepage main{min-height:40rem}}
.copyright-page{max-width:42rem}.copyright-page p{font-size:.95rem}
.front-nav{text-align:center;margin-top:2rem}
"""
CSS += '\n.sample-actions{display:flex;flex-wrap:wrap;gap:.7rem;margin:1rem 0}.sample-actions a{display:inline-block;padding:.65rem 1rem;border-radius:.65rem;background:#006078;color:white;text-decoration:none}\n'
(WEB/'book.css').write_text(CSS)
shutil.copytree(ROOT/'assets/fonts', WEB/'fonts', dirs_exist_ok=True)
(WEB/'cover').mkdir()
shutil.copy2(ROOT/PUBLICATION['cover'], WEB/'cover'/COVER_NAME)
(WEB/'diagrams').mkdir(exist_ok=True)
for scene in (ROOT/'assets/diagrams').glob('*.json'):
    svg=render_svg(json.loads(scene.read_text()))
    (scene.with_suffix('.svg')).write_text(svg)
    (WEB/'diagrams'/scene.with_suffix('.svg').name).write_text(svg)
shutil.copytree(ROOT/'assets/screenshots', WEB/'screenshots', dirs_exist_ok=True)
chapters=[]
for chapter in META['chapters']:
    file=ROOT/'build/reading'/Path(chapter['source']).name
    text=file.read_text()
    if '<!-- include:' in text: raise RuntimeError('Unresolved source includes')
    title=text.splitlines()[0].removeprefix('# ')
    text=re.sub(r'(?:\.\./)+assets/diagrams/', 'diagrams/', text)
    text=re.sub(r'(?:\.\./)+assets/screenshots/', 'screenshots/', text)
    body=markdown.markdown(text,extensions=['fenced_code','tables','toc'],output_format='xhtml')
    # Parsing now catches malformed XHTML before packaging EPUB.
    tree=highlight_tree(ET.fromstring('<root>'+body+'</root>'))
    for element in list(tree):
        picture=element.find('img') if element.tag=='p' else None
        if picture is not None and picture.get('src','').startswith('screenshots/'):
            caption=ET.Element('p', {'class':'figure-caption'})
            caption.text=picture.get('alt','')
            tree.insert(list(tree).index(element)+1,caption)
    body=''.join(ET.tostring(child,encoding='unicode',method='xml') for child in tree)
    chapters.append(dict(stem=file.stem,title=title,body=body,text=text))

def page(title, body, extra=''):
    return f'''<!DOCTYPE html>
<html lang="ru"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>{html.escape(title)}</title><link rel="stylesheet" href="book.css"></head><body><a class="skip" href="#content">К содержимому страницы</a><nav aria-label="Навигация книги"><a href="toc.html">Оглавление книги</a></nav><p class="status">{EDITION}. Бесплатный фрагмент: главы 1–12.</p><main id="content">{body}</main>{extra}</body></html>'''
for i,c in enumerate(FRONT):
    previous = FRONT[i-1]['stem']+'.html' if i else None
    following = FRONT[i+1]['stem']+'.html' if i+1<len(FRONT) else 'toc.html'
    links = ([f'<a href="{previous}">Назад</a>'] if previous else [])
    links.append(f'<a href="{following}">Далее</a>')
    body_class = 'front-cover' if c['kind']=='cover' else c['kind']
    document = f'''<!DOCTYPE html><html lang="ru"><head><meta charset="utf-8"/><meta name="viewport" content="width=device-width, initial-scale=1"/><title>{html.escape(c['title'])} — {html.escape(TITLE)}</title><link rel="stylesheet" href="book.css"/></head><body class="{body_class}"><main id="content">{c['body']}</main><nav class="front-nav" aria-label="Переход между страницами">{' · '.join(links)} · <a href="toc.html">Оглавление</a></nav></body></html>'''
    (WEB/(c['stem']+'.html')).write_text(document)
for i,c in enumerate(chapters):
    links=[]
    if i: links.append(f'<a href="{chapters[i-1]["stem"]}.html">← Предыдущая глава</a>')
    else: links.append('<a href="toc.html">← Оглавление</a>')
    if i+1<len(chapters):links.append(f'<a href="{chapters[i+1]["stem"]}.html">Следующая глава →</a>')
    (WEB/(c['stem']+'.html')).write_text(page(c['title'],c['body'],'<nav aria-label="Переход между главами">'+' · '.join(links)+'</nav>'))
index='<h1>'+html.escape(TITLE)+'</h1><p>Предисловие и главы 1–12 доступны бесплатно. Продолжение: <a href="https://shanraq.org/shop/go-book">полная книга на shanraq.org</a>. Примеры этих глав входят в бесплатный архив кода.</p><ol>'
index+=''.join(f'<li><a href="{c["stem"]}.html">{html.escape(c["title"])}</a></li>' for c in FRONT+chapters)+'</ol>'
index+='''<section class="search" aria-labelledby="search-title"><h2 id="search-title">Поиск по написанным главам</h2><label for="query">Слово или фраза</label><input id="query" type="search" autocomplete="off"><p id="search-status" role="status" aria-live="polite"></p><ul id="results"></ul><noscript>Для поиска включите JavaScript. Оглавление и главы работают без него.</noscript></section>'''
(WEB/'toc.html').write_text(page(TITLE,index,'<script src="search-data.js"></script><script src="search.js"></script>'))
data=[dict(title=c['title'],url=c['stem']+'.html',text=' '.join(ET.fromstring('<root>'+c['body']+'</root>').itertext())) for c in chapters]
(WEB/'search-data.js').write_text('window.BOOK_SEARCH = '+json.dumps(data,ensure_ascii=False).replace('<','\\u003c')+';\n')
(WEB/'search.js').write_text('''"use strict";
const query=document.getElementById("query");
query.addEventListener("input",()=>{
 const list=document.getElementById("results");list.replaceChildren();
 const term=query.value.trim().toLocaleLowerCase("ru");
 const hits=term?window.BOOK_SEARCH.filter(x=>x.text.toLocaleLowerCase("ru").includes(term)):[];
 for(const hit of hits){const li=document.createElement("li");const a=document.createElement("a");a.href=hit.url;a.textContent=hit.title;li.append(a);list.append(li);}
 document.getElementById("search-status").textContent=term?"Найдено глав: "+hits.length:"";
});
''')
with zipfile.ZipFile(OUT/'go-book-preview-html.zip','w',zipfile.ZIP_DEFLATED) as archive:
    for file in sorted(WEB.rglob('*')):
        if file.is_file():archive.write(file,file.relative_to(WEB).as_posix())

# EPUB 3, script-free reading order and XHTML navigation.
epub_items={}
for c in FRONT:
    epub_items[c['stem']+'.xhtml']=f'''<?xml version="1.0" encoding="utf-8"?><html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" lang="ru" xml:lang="ru"><head><title>{html.escape(c['title'])}</title><link rel="stylesheet" type="text/css" href="book.css"/></head><body class="{'front-cover' if c['kind']=='cover' else c['kind']}"><section epub:type="{c['kind']}">{c['body']}</section></body></html>'''
for c in chapters:
    epub_items[c['stem']+'.xhtml']=f'''<?xml version="1.0" encoding="utf-8"?>
<html xmlns="http://www.w3.org/1999/xhtml" lang="ru" xml:lang="ru"><head><title>{html.escape(c['title'])}</title><link rel="stylesheet" type="text/css" href="book.css"/></head><body><p>{EDITION}. Полный черновик книги.</p>{c['body']}</body></html>'''
nav=''.join(f'<li><a href="{c["stem"]}.xhtml">{html.escape(c["title"])}</a></li>' for c in FRONT+chapters)
epub_items['nav.xhtml']=f'''<?xml version="1.0" encoding="utf-8"?><html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" lang="ru" xml:lang="ru"><head><title>Оглавление</title><link rel="stylesheet" type="text/css" href="book.css"/></head><body><nav epub:type="toc" id="toc"><h1>Оглавление</h1><ol>{nav}</ol></nav></body></html>'''
manifest=''.join(f'<item id="c{i}" href="{c["stem"]}.xhtml" media-type="application/xhtml+xml"/>' for i,c in enumerate(chapters))
manifest+=''.join(f'<item id="front{i}" href="{c["stem"]}.xhtml" media-type="application/xhtml+xml"/>' for i,c in enumerate(FRONT))
manifest+=f'<item id="cover-image" href="cover/{COVER_NAME}" media-type="{COVER_MIME}" properties="cover-image"/>'
spine=''.join(f'<itemref idref="c{i}"/>' for i in range(len(chapters)))
spine=''.join(f'<itemref idref="front{i}"/>' for i in range(len(FRONT)))+'<itemref idref="nav"/>'+spine
fonts=list((ROOT/'assets/fonts').glob('*.ttf'))
manifest+=''.join(f'<item id="font{i}" href="fonts/{f.name}" media-type="font/ttf"/>' for i,f in enumerate(fonts))
manifest+=''.join(f'<item id="diagram{i}" href="diagrams/{file.name}" media-type="image/svg+xml"/>' for i,file in enumerate(sorted((WEB/'diagrams').glob('*.svg'))))
manifest+=''.join(f'<item id="screenshot{i}" href="screenshots/{file.name}" media-type="image/png"/>' for i,file in enumerate(sorted((WEB/'screenshots').glob('*.png'))))
manifest+='<item id="font-license" href="fonts/OFL.txt" media-type="text/plain"/>'
modified=datetime.now(timezone.utc).strftime('%Y-%m-%dT%H:%M:%SZ')
opf=f'''<?xml version="1.0" encoding="utf-8"?><package xmlns="http://www.idpf.org/2007/opf" version="3.0" unique-identifier="book-id" xml:lang="ru"><metadata xmlns:dc="http://purl.org/dc/elements/1.1/"><dc:identifier id="book-id">urn:shanraq:go-book:preview:{META['edition']}</dc:identifier><dc:title>{html.escape(TITLE)} — рабочая редакция</dc:title><dc:language>ru</dc:language><dc:creator>{html.escape(PUBLICATION["author"])}</dc:creator><dc:date>{PUBLICATION["date"]}</dc:date><dc:publisher>{html.escape(PUBLICATION["publisher"])}</dc:publisher><dc:description>{EDITION}. Полный черновик для редакционной проверки.</dc:description><meta property="dcterms:modified">{modified}</meta></metadata><manifest><item id="nav" href="nav.xhtml" media-type="application/xhtml+xml" properties="nav"/><item id="css" href="book.css" media-type="text/css"/>{manifest}</manifest><spine>{spine}</spine></package>'''
ET.fromstring(opf)
with zipfile.ZipFile(OUT/'go-book-preview.epub','w') as archive:
    archive.writestr('mimetype','application/epub+zip',compress_type=zipfile.ZIP_STORED)
    archive.writestr('META-INF/container.xml','''<?xml version="1.0"?><container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container"><rootfiles><rootfile full-path="EPUB/package.opf" media-type="application/oebps-package+xml"/></rootfiles></container>''')
    archive.writestr('EPUB/package.opf',opf,compress_type=zipfile.ZIP_DEFLATED)
    archive.write(ROOT/PUBLICATION['cover'],'EPUB/cover/'+COVER_NAME,compress_type=zipfile.ZIP_DEFLATED)
    archive.writestr('EPUB/book.css',CSS.replace('overflow:auto','white-space:pre-wrap;overflow-wrap:anywhere'),compress_type=zipfile.ZIP_DEFLATED)
    for path,text in epub_items.items():
        ET.fromstring(text)
        archive.writestr('EPUB/'+path,text,compress_type=zipfile.ZIP_DEFLATED)
    for font in fonts:archive.write(font,'EPUB/fonts/'+font.name,compress_type=zipfile.ZIP_DEFLATED)
    archive.write(ROOT/'assets/fonts/OFL.txt','EPUB/fonts/OFL.txt',compress_type=zipfile.ZIP_DEFLATED)
    for file in (WEB/'diagrams').glob('*.svg'):
        archive.write(file,'EPUB/diagrams/'+file.name,compress_type=zipfile.ZIP_DEFLATED)

    for file in (WEB/'screenshots').glob('*.png'):
        archive.write(file,'EPUB/screenshots/'+file.name,compress_type=zipfile.ZIP_DEFLATED)

# PDF uses the same parsed XHTML; preserve tables and text code.
for name,file in [('Noto','NotoSans-Regular.ttf'),('NotoBold','NotoSans-Bold.ttf'),('NotoMono','NotoSansMono-Regular.ttf')]:
    pdfmetrics.registerFont(TTFont(name,str(ROOT/'assets/fonts'/file)))
pdfmetrics.registerFontFamily('Noto',normal='Noto',bold='NotoBold',italic='Noto',boldItalic='NotoBold')
styles={
 'body':ParagraphStyle('body',fontName='Noto',fontSize=10,leading=15,spaceAfter=8),
 'h1':ParagraphStyle('h1',fontName='NotoBold',fontSize=22,leading=28,spaceAfter=18,textColor=colors.HexColor('#00758D'),keepWithNext=True),
 'h2':ParagraphStyle('h2',textColor=colors.HexColor('#00758D'),fontName='NotoBold',fontSize=14,leading=19,spaceBefore=15,spaceAfter=8,keepWithNext=True),
 'h3':ParagraphStyle('h3',textColor=colors.HexColor('#00758D'),fontName='NotoBold',fontSize=11,leading=16,spaceBefore=10,spaceAfter=6,keepWithNext=True),
 'small':ParagraphStyle('small',fontName='Noto',fontSize=8,leading=11,spaceAfter=5),
}

def inline(element):
    text=escaped_text(element.text)
    for child in element:
        content=inline(child)
        if child.tag=='code':text+='<font name="NotoMono">'+content+'</font>'
        elif child.tag=='span' and child.get('class','').startswith('tok-'):
            colour=PALETTE[child.get('class')[4:]]
            text+='<font color="'+colour+'">'+content+'</font>'
        elif child.tag in ('strong','b'):text+='<b>'+content+'</b>'
        elif child.tag=='a':text+='<a href="'+html.escape(child.get('href',''),quote=True)+'" color="#00758D">'+content+'</a>'
        elif child.tag=='br':text+='<br/>'
        else:text+=content
        text+=escaped_text(child.tail)
    return text

title_styles={
 'author':ParagraphStyle('title-author',fontName='Noto',fontSize=9,leading=12.5,alignment=1,spaceAfter=0),
 'title':ParagraphStyle('title-title',fontName='NotoBold',fontSize=28,leading=36,alignment=1,textColor=colors.HexColor('#00758D'),spaceAfter=25),
 'subtitle':ParagraphStyle('title-subtitle',fontName='Noto',fontSize=13,leading=20,alignment=1,spaceAfter=24),
 'imprint':ParagraphStyle('title-imprint',fontName='Noto',fontSize=11,leading=17,alignment=1,spaceAfter=15),
}
class TitlePage(Flowable):
    """Place author, central title block and imprint independently on one page."""
    def __init__(self, body):
        super().__init__()
        self.width = 473
        self.height = 720
        self.root = ET.fromstring('<root>' + body + '</root>')

    def paragraph(self, element, style):
        paragraph = Paragraph(inline(element), style)
        _, height = paragraph.wrap(self.width, self.height)
        return paragraph, height

    def draw(self):
        author, _ = self.paragraph(self.root.find("p[@class='author']"), title_styles['author'])
        author.drawOn(self.canv, 0, self.height - 42)
        entries = []
        for element in self.root.find("div[@class='title-block']"):
            role = element.get('class', '')
            style = title_styles['title' if element.tag == 'h1' else 'subtitle' if role == 'subtitle' else 'imprint']
            paragraph, height = self.paragraph(element, style)
            entries.append((paragraph, height, style.spaceAfter))
        total = sum(height + gap for _, height, gap in entries) - entries[-1][2]
        top = self.height / 2 + total / 2
        for paragraph, height, gap in entries:
            paragraph.drawOn(self.canv, 0, top - height)
            top -= height + gap
        imprint, _ = self.paragraph(self.root.find("p[@class='imprint']"), title_styles['imprint'])
        imprint.drawOn(self.canv, 0, 10)

# Cover is drawn by the page callback; the next two pages have no printed folios.
story = [Spacer(1, 1), PageBreak(), TitlePage(FRONT[1]['body']), PageBreak()]
for element in ET.fromstring('<root>'+FRONT[2]['body']+'</root>'):
    story.append(Paragraph(inline(element),styles['h2'] if element.tag=='h1' else styles['body']))
story.extend([PageBreak(),Paragraph('Оглавление',styles['h1'])])
for i,c in enumerate(chapters):story.append(Paragraph(f'<a href="#chapter-{i}">{html.escape(c["title"])}</a>',styles['body']))

class Diagram(Flowable):
    def __init__(self, scene):
        super().__init__()
        self.scene=scene
        self.width=480
        self.height=480*scene['height']/scene['width']
    def draw(self):
        scene=drawing(self.scene)
        scale=480/self.scene['width']
        self.canv.saveState()
        self.canv.scale(scale,scale)
        renderPDF.draw(scene,self.canv,0,0)
        self.canv.restoreState()

def append_elements(root):
    for element in root:
        tag=element.tag
        if tag in ('h1','h2','h3'):story.append(Paragraph(inline(element),styles[tag]))
        elif tag=='pre':
            code=''.join(element.itertext()).expandtabs(4).rstrip()
            max_width=max((pdfmetrics.stringWidth(line,'NotoMono',8) for line in code.splitlines()),default=1)
            size=min(8,8*475/max_width)
            if size<6:raise RuntimeError('Code line too long for legible PDF; edit source')
            code_style=ParagraphStyle('code',fontName='NotoMono',fontSize=size,leading=size*1.5,spaceBefore=5,spaceAfter=10,leftIndent=5)
            story.append(XPreformatted(inline(element).replace('\t','    ').rstrip(),code_style))
        elif tag=='table':
            rows=[]
            for row in element.findall('.//tr'):
                rows.append([Paragraph(inline(cell),styles['small']) for cell in row])
            if rows:
                table=Table(rows,colWidths=[480/len(rows[0])]*len(rows[0]),repeatRows=1,hAlign='LEFT')
                table.setStyle(TableStyle([('GRID',(0,0),(-1,-1),.4,colors.HexColor('#aabfbc')),('BACKGROUND',(0,0),(-1,0),colors.HexColor('#ffffff')),('VALIGN',(0,0),(-1,-1),'TOP'),('LEFTPADDING',(0,0),(-1,-1),7),('RIGHTPADDING',(0,0),(-1,-1),7),('TOPPADDING',(0,0),(-1,-1),6),('BOTTOMPADDING',(0,0),(-1,-1),6)]))
                story.extend([table,Spacer(1,10)])
        elif tag in ('ul','ol'):
            for i,li in enumerate(element):
                story.append(Paragraph((str(i+1)+'. ' if tag=='ol' else '• ')+inline(li),styles['body']))
        elif tag=='blockquote':append_elements(element)
        elif tag=='p' and element.find('img') is not None:
            img=element.find('img')
            if img.get('src','').startswith('screenshots/'):
                picture=Image(str(WEB/img.get('src')))
                scale=min(485/picture.imageWidth, 540/picture.imageHeight)
                picture.drawWidth=picture.imageWidth*scale
                picture.drawHeight=picture.imageHeight*scale
                story.append(KeepTogether([picture, Spacer(1,6), Paragraph(html.escape(img.get('alt','')),styles['small'])]))
            else:
                scene_path=ROOT/'assets/diagrams'/(Path(img.get('src')).stem+'.json')
                story.append(KeepTogether([Diagram(json.loads(scene_path.read_text())),Paragraph(html.escape(img.get('alt','')),styles['small'])]))
        elif tag=='p' and element.get('class')=='figure-caption':pass
        elif tag=='p':story.append(Paragraph(inline(element),styles['body']))
        else:raise RuntimeError('Unhandled PDF element: '+tag)
for i,c in enumerate(chapters):
    story.extend([PageBreak(),Paragraph(f'<a name="chapter-{i}"/>',styles['small'])])
    append_elements(ET.fromstring('<root>'+c['body']+'</root>'))

def footer(canvas,doc):
    canvas.saveState()
    canvas.setFillColor(colors.white)
    canvas.rect(0,0,doc.pagesize[0],doc.pagesize[1],fill=1,stroke=0)
    canvas.restoreState()
    if doc.page == 1:
        from reportlab.lib.utils import ImageReader
        cover=ImageReader(str(ROOT/PUBLICATION['cover']))
        width,height=cover.getSize()
        scale=min(doc.pagesize[0]/width,doc.pagesize[1]/height)
        canvas.drawImage(cover,(doc.pagesize[0]-width*scale)/2,(doc.pagesize[1]-height*scale)/2,width=width*scale,height=height*scale)
        canvas.bookmarkPage('cover')
        canvas.addOutlineEntry('Обложка','cover',0)
    if doc.page in (2,3):
        key,label=('title-page','Титульный лист') if doc.page==2 else ('publication-details','Выходные сведения')
        canvas.bookmarkPage(key)
        canvas.addOutlineEntry(label,key,0)
    if doc.page <= 3:
        return
    canvas.setFont('Noto',8);canvas.setFillColor(colors.HexColor('#526669'))
    canvas.drawString(55,28,'shanraq.org · '+META['edition'])
    canvas.drawRightString(540,28,str(doc.page))

class BookDoc(SimpleDocTemplate):
    def afterFlowable(self, flowable):
        if self.page > 3 and isinstance(flowable, Paragraph) and flowable.style.name in ('h1', 'h2'):
            key='outline-'+str(getattr(self, 'outline_count', 0))
            self.outline_count=getattr(self, 'outline_count', 0)+1
            self.canv.bookmarkPage(key)
            self.canv.addOutlineEntry(flowable.getPlainText(), key, 0 if flowable.style.name=='h1' else 1, closed=False)

doc=BookDoc(str(OUT/'go-book-preview.pdf'),pagesize=(595.28,841.89),rightMargin=55,leftMargin=55,topMargin=50,bottomMargin=50,title=TITLE+' — рабочая редакция',author=PUBLICATION['author'],subject=PUBLICATION['subtitle'])
doc.build(story,onFirstPage=footer,onLaterPages=footer)
code_archive = OUT/'go-book-preview-code.zip'
with zipfile.ZipFile(code_archive, 'w', zipfile.ZIP_DEFLATED) as archive:
    for folder in ('examples', 'research/reproductions/map-value'):
        for file in sorted((ROOT/folder).rglob('*')):
            if file.is_file() and (file.name in ('.gitignore', '.gitkeep') or file.suffix in ('.go', '.mod', '.sum', '.txt', '.html', '.css', '.sql', '.json', '.png', '.ico', '.yml', '.conf', '.service')):
                archive.write(file, 'book/'+file.relative_to(ROOT).as_posix())
    archive.write(ROOT/'book.json', 'book/book.json')
    archive.write(ROOT.parent/'LICENSE', 'LICENSE')
    archive.write(ROOT/'assets/brand/NOTICE.txt', 'ASSET-NOTICES.txt')
    archive.write(ROOT.parent/'.editorconfig', '.editorconfig')
    archive.writestr('book/TRANSITIONS.md', render_transitions(ROOT, META))
    archive.writestr('README.txt', 'Код рабочей книги '+META['edition']+'\n'
        'Требуется Go '+META['go_version']+'. Распакуйте архив.\n'
        'Откройте терминал в book/examples/02-first-program и выполните go run .\n'
        'Для другой главы откройте её папку. С главы 10 запускайте go run ./cmd/shop. Начиная с главы 5: go test ./...\n'
        'У каждой папки свой go.mod: повторный go mod init не нужен.\n'
        'Точный список новых, изменённых и удалённых файлов каждой главы: book/TRANSITIONS.md.\n'
        'Сохраняйте задания в копиях папок. Python, npm, Docker не нужны.\n'
        'Состав глав указан в book/book.json; это рабочая редакция.\n')
paths=[OUT/'go-book-preview.pdf',OUT/'go-book-preview.epub',OUT/'go-book-preview-html.zip', code_archive]
(OUT/'SHA256SUMS.txt').write_text(''.join(hashlib.sha256(f.read_bytes()).hexdigest()+'  '+f.name+'\n' for f in paths))
(OUT/'README.txt').write_text(EDITION+'\nPDF и EPUB открываются в соответствующей программе. HTML ZIP распакуйте и откройте index.html.\nЭто ознакомительный фрагмент: предисловие и главы 1–12. Полная книга: https://shanraq.org/shop/go-book. Полная проверка доступности и совместимости со всеми читалками ещё не завершена.\nШрифты Noto используются по SIL OFL 1.1; лицензия включена в EPUB/HTML и доступна в assets/fonts/OFL.txt исходного комплекта.\n')
print('Built PDF, EPUB, offline HTML and SHA-256 manifest')
