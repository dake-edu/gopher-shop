"""Fail closed if a public build gains chapters or source beyond chapter 12."""
import json
from pathlib import Path
import re
import zipfile


def audit(root):
    meta = json.loads((root/'book.json').read_text())
    assert meta['sample']['last_chapter'] == 12, 'Not an authorized sample'
    assert [c['id'] for c in meta['chapters']] == [f'{n:02}' for n in range(13)]
    assert {p.relative_to(root).as_posix() for p in (root/'manuscript').glob('*.md')} == {c['source'] for c in meta['chapters']}
    assert {p.name for p in (root/'examples').iterdir() if p.is_dir()} == {Path(c['checkpoint']).name for c in meta['chapters'] if 'checkpoint' in c}
    for folder in ['manuscript', 'examples']:
        for p in (root/folder).rglob('*'):
            assert not p.is_symlink(), 'Symlink is not allowed in sample: '+str(p)
    return meta


def audit_packages(root):
    meta = audit(root)
    allowed = {Path(c['source']).stem for c in meta['chapters']}
    expected_pages = allowed | {'index', 'toc', 'title-page', 'publication-details'}
    out = root/'build/preview'
    assert {p.stem for p in (out/'html').glob('*.html')} == expected_pages
    for name in ['go-book-preview.epub', 'go-book-preview-html.zip', 'go-book-preview-code.zip']:
        with zipfile.ZipFile(out/name) as archive:
            for path in archive.namelist():
                assert '.private-book' not in path
                for part in Path(path).parts:
                    chapter = re.match(r'^(\d{2})-', part)
                    assert not chapter or int(chapter[1]) <= 12, 'Paid chapter in sample: '+path
            if name.endswith('code.zip'):
                assert json.loads(archive.read('book/book.json')) == meta
    print('PASS: sample contains only the preface and chapters 1–12')


if __name__ == '__main__':
    audit_packages(Path(__file__).resolve().parents[1])
