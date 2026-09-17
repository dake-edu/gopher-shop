"""Build a self-contained, public website package from the checked 12-chapter sample."""
import hashlib
import html
import json
from pathlib import Path
import shutil
import zipfile
from sample_boundary import audit_packages

root = Path(__file__).resolve().parents[1]
audit_packages(root)
meta = json.loads((root/'book.json').read_text())
preview = root/'build/preview'
site = root/'build/sample-site'
if site.exists():
    shutil.rmtree(site)
site.mkdir(parents=True)
shutil.copytree(preview/'html', site/'read')
files = ['go-book-preview.pdf', 'go-book-preview.epub', 'go-book-preview-html.zip', 'go-book-preview-code.zip']
for name in files:
    shutil.copy2(preview/name, site/name)
shutil.copy2(preview/'SHA256SUMS.txt', site/'SHA256SUMS.txt')
(site/'index.html').write_text('''<!doctype html>
<html lang="ru"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>Читать 12 глав бесплатно — Go: от первой строки до интернет-магазина</title>
<link rel="stylesheet" href="read/book.css">
<style>.sample-actions{display:flex;flex-wrap:wrap;gap:.75rem}.sample-actions a{padding:.7rem 1rem;border-radius:.7rem;background:#006078;color:white;text-decoration:none}main{padding:1rem 0}footer{margin-top:3rem}a:focus-visible{outline:3px solid #b08242;outline-offset:4px}</style>
</head><body><main>
<p>shanraq.org · Бесплатный ознакомительный фрагмент</p>
<h1>Go: от первой строки до интернет-магазина</h1>
<p>Предисловие и главы 1–12: от первой программы до HTTP-каталога с HTML-шаблонами. Читайте на сайте или скачайте для чтения без интернета. Регистрация и покупка не требуются.</p>
<div class="sample-actions"><a href="read/toc.html">Читать 12 глав онлайн</a><a href="go-book-preview.pdf" download>Скачать PDF</a><a href="go-book-preview.epub" download>Скачать EPUB</a><a href="go-book-preview-html.zip" download>Скачать HTML</a></div>
<p>HTML ZIP распакуйте и откройте index.html. Для EPUB нужна программа чтения электронных книг.</p>
<p><a href="go-book-preview-code.zip" download>Скачать примеры глав 2–12</a> · <a href="SHA256SUMS.txt">Контрольные суммы файлов</a></p>
<h2>Что будет дальше</h2><p>В полной книге — 42 главы: хранение данных, вход и регистрация, заказы, тестовая оплата, выдача книг, административная часть и эксплуатация приложения. Учебная оплата не списывает реальные деньги.</p>
<div class="sample-actions"><a href="https://shanraq.org/shop/go-book">Посмотреть полную книгу и условия покупки</a></div>
</main><footer><p>Ознакомительный фрагмент '''+html.escape(meta['edition'])+''' · © 2026 Баймурзин Даулет Абузарович</p></footer></body></html>''', encoding='utf-8')
(site/'sample.json').write_text(json.dumps({'edition': meta['edition'], 'chapters': [c['id'] for c in meta['chapters']], 'files': {p.name: hashlib.sha256(p.read_bytes()).hexdigest() for p in (site/n for n in files)}}, ensure_ascii=False, indent=2)+'\n')
archive = preview/'go-book-sample-site.zip'
with zipfile.ZipFile(archive, 'w', zipfile.ZIP_DEFLATED) as output:
    for p in sorted(site.rglob('*')):
        if p.is_file():
            output.write(p, p.relative_to(site).as_posix())
# Recompute, rather than append, so repeated builds remain reproducible in composition.
(preview/'SHA256SUMS.txt').write_text(''.join(hashlib.sha256((preview/n).read_bytes()).hexdigest()+'  '+n+'\n' for n in [*files, archive.name]))
print('PASS: website package with online reading and four free downloads: '+str(archive))
