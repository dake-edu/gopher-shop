#!/usr/bin/env python3
"""Structural review checks. This is not EPUBCheck or an accessibility audit."""
import os as _os
import sys as _sys
if not _sys.flags.utf8_mode:
    _os.environ['PYTHONUTF8'] = '1'
    _os.execv(_sys.executable, [_sys.executable, '-X', 'utf8', *_sys.argv])

from pathlib import Path
from html.parser import HTMLParser
from urllib.parse import urlparse, unquote
import hashlib
import json
import zipfile
import xml.etree.ElementTree as ET

ROOT = Path(__file__).resolve().parents[1]
OUT = ROOT/'build/preview'
META=json.loads((ROOT/'book.json').read_text())
class Links(HTMLParser):
    def __init__(self):
        super().__init__(); self.links=[]; self.ids=set()
    def handle_starttag(self, tag, attrs):
        attributes=dict(attrs)
        if tag == 'img' and not attributes.get('alt'):
            raise ValueError('Image is missing meaningful alt text')
        if 'id' in attributes:
            if attributes['id'] in self.ids: raise ValueError('Duplicate HTML id')
            self.ids.add(attributes['id'])
        for key in ('href','src'):
            if key in attributes: self.links.append(attributes[key])

pages={}
for file in (OUT/'html').glob('*.html'):
    parser=Links(); parser.feed(file.read_text()); pages[file.name]=parser
count=0
for name, parser in pages.items():
    for link in parser.links:
        url=urlparse(link)
        if url.scheme: continue
        target=(OUT/'html'/unquote(url.path)) if url.path else (OUT/'html'/name)
        if not target.exists(): raise ValueError((name,link))
        if url.fragment and target.suffix=='.html':
            if unquote(url.fragment) not in pages[target.name].ids: raise ValueError((name,link))
        count+=1
with zipfile.ZipFile(OUT/'go-book-preview.epub') as archive:
    assert archive.infolist()[0].filename=='mimetype'
    assert archive.infolist()[0].compress_type==zipfile.ZIP_STORED
    assert archive.read('mimetype')==b'application/epub+zip'
    ns={'o':'http://www.idpf.org/2007/opf'}
    opf=ET.fromstring(archive.read('EPUB/package.opf'))
    items={item.get('id'):item for item in opf.findall('o:manifest/o:item',ns)}
    for item in items.values(): assert 'EPUB/'+item.get('href') in archive.namelist()
    for item in opf.findall('o:spine/o:itemref',ns): assert item.get('idref') in items
    order=[items[item.get('idref')].get('href') for item in opf.findall('o:spine/o:itemref',ns)]
    assert order[:5]==['index.xhtml','title-page.xhtml','publication-details.xhtml','nav.xhtml','00-preface.xhtml']
    covers=[item for item in items.values() if 'cover-image' in item.get('properties','').split()]
    assert len(covers)==1
    assert archive.read('EPUB/'+covers[0].get('href'))==(ROOT/META['publication']['cover']).read_bytes()
    creator=opf.find('o:metadata/{http://purl.org/dc/elements/1.1/}creator',ns)
    assert creator is not None and creator.text==META['publication']['author']
    for name in archive.namelist():
        if name.endswith(('.xhtml','.xml','.opf','.svg')): ET.fromstring(archive.read(name))
assert (OUT/'html/cover/go-book-cover.jpg').read_bytes()==(ROOT/META['publication']['cover']).read_bytes()
assert 'title-page.html' in pages['index.html'].links
assert 'publication-details.html' in pages['title-page.html'].links
assert 'toc.html' in pages['publication-details.html'].links
with zipfile.ZipFile(OUT/'go-book-preview-html.zip') as archive:
    assert 'index.html' in archive.namelist()
    for file in (OUT/'html').rglob('*'):
        if file.is_file(): assert archive.read(file.relative_to(OUT/'html').as_posix())==file.read_bytes()
for line in (OUT/'SHA256SUMS.txt').read_text().splitlines():
    digest,name=line.split('  ',1)
    assert hashlib.sha256((OUT/name).read_bytes()).hexdigest()==digest
assert (OUT/'go-book-preview.pdf').read_bytes().startswith(b'%PDF-')
report={'status':'passed','html_pages':len(pages),'local_links':count,
        'checks':['HTML links and unique IDs','EPUB XML, manifest, spine and mimetype','front matter order, author metadata and cover image','HTML ZIP matches directory','artifact SHA-256'],
        'not_checked':['EPUBCheck','screen reader reading order','full browser/ereader compatibility','PDF tagging/accessibility']}

# The reader bundle must contain exactly the source used for this edition.
with zipfile.ZipFile(OUT/'go-book-preview-code.zip') as archive:
    expected = {'README.txt', 'LICENSE', 'book/book.json'}
    for folder in ('examples', 'research/reproductions/map-value'):
        for file in sorted((ROOT/folder).rglob('*')):
            if file.is_file() and file.suffix in ('.go', '.mod', '.txt'):
                name = 'book/'+file.relative_to(ROOT).as_posix()
                expected.add(name)
                assert archive.read(name) == file.read_bytes(), name
    assert set(archive.namelist()) == expected
    assert archive.read('book/book.json') == (ROOT/'book.json').read_bytes()
print('PASS: reader source bundle matches the current checkpoints')

report['checks'].append('Reader code ZIP matches current sources')
(OUT/'validation.json').write_text(json.dumps(report,ensure_ascii=False,indent=2)+'\n')
print(f'PASS: {len(pages)} HTML pages, {count} local links, EPUB structure and artifact hashes')
