"""Install the checked public sample and reviewable shop links into shanraq.org."""
import argparse
import hashlib
import json
from pathlib import Path
import shutil
from sample_boundary import audit_packages

parser = argparse.ArgumentParser()
parser.add_argument('site', type=Path)
parser.add_argument('--apply', action='store_true')
args = parser.parse_args()
root = Path(__file__).resolve().parents[1]
audit_packages(root)
source = root/'build/sample-site'
meta = json.loads((source/'sample.json').read_text())
assert meta['chapters'] == [f'{n:02}' for n in range(13)]
for name, digest in meta['files'].items():
    assert hashlib.sha256((source/name).read_bytes()).hexdigest() == digest
prefix = '/static/shop/go-book-sample/'+meta['edition']
destination = args.site/'web'/prefix.lstrip('/')
template = args.site/'pkg/modules/shop/templates/shop.html'
s = template.read_text()
old = '''        {{ if .P.PreviewURL }}
        <p class="shop-offer__acts">'''
new = '''        {{ if eq .P.Slug "go-book" }}
        <p class="shop-offer__acts">
          <a class="btn btn--outline" href="PREFIX/read/toc.html" data-track="shop_preview_online">{{ t .Lang "shop.sample_online" }}</a>
          <a class="btn btn--outline" href="PREFIX/go-book-preview.pdf" download data-track="shop_preview_pdf">PDF</a>
          <a class="btn btn--outline" href="PREFIX/go-book-preview.epub" download data-track="shop_preview_epub">EPUB</a>
          <a class="btn btn--outline" href="PREFIX/go-book-preview-html.zip" download data-track="shop_preview_html">HTML ZIP</a>
        </p>
        <p class="hint">{{ t .Lang "shop.sample_note" }} <a href="PREFIX/go-book-preview-code.zip" download>{{ t .Lang "shop.sample_code" }}</a></p>
        {{ else if .P.PreviewURL }}
        <p class="shop-offer__acts">'''.replace('PREFIX', prefix)
assert s.count(old) == 1, 'Shop template changed; review the patch before applying'
translation = args.site/'pkg/site/i18n.go'
t = translation.read_text()
needle = '\t"shop.preview":'
assert needle in t and '"shop.sample_online":' not in t
pos = t.index(needle)
t = t[:pos]+'''\t"shop.sample_online": {"kz": "12 тарауды онлайн тегін оқу", "ru": "Читать 12 глав онлайн бесплатно", "en": "Read 12 chapters online for free"},
\t"shop.sample_note": {"kz": "Орыс тіліндегі кіріспе және 1–12 тараулар. Тіркелусіз оқыңыз немесе офлайн оқу үшін жүктеңіз.", "ru": "Предисловие и главы 1–12 на русском языке. Читайте без регистрации или скачайте для чтения без интернета.", "en": "The preface and chapters 1–12 in Russian. No sign-up: read online or download for offline reading."},
\t"shop.sample_code": {"kz": "Мысалдар кодын жүктеу", "ru": "Скачать код примеров", "en": "Download the example code"},
'''+t[pos:]
changes = {template:s.replace(old,new), translation:t}
document = args.site/'docs/go-book-sample.md'
changes[document] = '''# Free Go book sample

The product page `/shop/go-book` links to the free preface and chapters 1–12 in Russian. Readers can open the HTML book without signing in or download PDF, EPUB, offline HTML and the example code. The last chapter links back to the full product.

The checked static package is embedded from `web/static/shop/go-book-sample/0.19.0-sample/`. Its `sample.json` declares the chapter boundary and file hashes. Chapters 13–42 and the full paid book are not included. Build the application again after updating these static files.

In the shop admin, use `/static/shop/go-book-sample/0.19.0-sample/index.html` for the preview URL and `/static/shop/go-book-sample/0.19.0-sample/read/cover/go-book-cover-v3.png` for the cover. The current Russian title is “Go: от первой строки до интернет-магазина”. The product remains in its existing sale state; a free sample does not enable payment or announce availability.

To regenerate, use the public gopher-shop repository: run the book checks, build and validate the sample, then `python book/tools/prepare_sample_site.py`. Inspect the website ZIP before replacing this version. Keep each new edition in a separate directory so cached files from different editions do not mix.
'''
changes[args.site/'web/sample_test.go'] = '''package web

import (
    "encoding/json"
    "net/http/httptest"
    "strings"
    "testing"
)

func TestBookSamplePublicReadingAndDownloads(t *testing.T) {
    prefix := "/shop/go-book-sample/0.19.0-sample/"
    for _, path := range []string{"read/toc.html", "read/12-templates.html", "go-book-preview.pdf", "go-book-preview.epub", "go-book-preview-html.zip", "go-book-preview-code.zip"} {
        response := httptest.NewRecorder()
        StaticHandler().ServeHTTP(response, httptest.NewRequest("GET", prefix+path, nil))
        if response.Code != 200 || response.Body.Len() == 0 {
            t.Fatalf("public sample %s: status %d, bytes %d", path, response.Code, response.Body.Len())
        }
    }
    response := httptest.NewRecorder()
    StaticHandler().ServeHTTP(response, httptest.NewRequest("GET", prefix+"read/13-forms.html", nil))
    if response.Code != 404 { t.Fatalf("paid chapter returned %d", response.Code) }
    data, err := staticFiles.ReadFile("static"+prefix+"sample.json")
    if err != nil { t.Fatal(err) }
    var sample struct { Chapters []string `json:"chapters"` }
    if err := json.Unmarshal(data, &sample); err != nil { t.Fatal(err) }
    if strings.Join(sample.Chapters, ",") != "00,01,02,03,04,05,06,07,08,09,10,11,12" { t.Fatal("unexpected sample chapter boundary") }
}
'''
for path, body in changes.items():
    print(('WRITE ' if args.apply else 'PREVIEW ')+str(path.relative_to(args.site))+f' ({len(body.encode())} bytes)')
print(('COPY ' if args.apply else 'PREVIEW COPY ')+str(destination.relative_to(args.site))+f' ({sum(p.is_file() for p in source.rglob("*"))} files)')
if args.apply:
    if destination.exists():
        raise SystemExit('Versioned sample directory already exists; do not overwrite published assets')
    shutil.copytree(source, destination)
    for path, body in changes.items():
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(body)
