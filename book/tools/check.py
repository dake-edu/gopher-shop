#!/usr/bin/env python3
"""Run checkpoints, compare output, vet and test; render includes for reading."""
import os as _os
import sys as _sys
if not _sys.flags.utf8_mode:
    _os.environ['PYTHONUTF8'] = '1'
    _os.execv(_sys.executable, [_sys.executable, '-X', 'utf8', *_sys.argv])

from pathlib import Path
import os
import re
import subprocess
import tempfile
import shutil
import json
import hashlib
import platform
from datetime import datetime, timezone
from reader_audit import audit as audit_reader_commands

ROOT = Path(__file__).resolve().parents[1]
ENV = dict(os.environ, GOWORK='off')
records = []
def source_hashes():
    return {f.relative_to(ROOT).as_posix(): hashlib.sha256(f.read_bytes()).hexdigest()
            for folder in ('manuscript', 'examples')
            for f in sorted((ROOT/folder).rglob('*')) if f.is_file()}

initial_sources = source_hashes()
initial_registry = (ROOT/'book.json').read_bytes()
metadata = json.loads((ROOT/'book.json').read_text())
registered = {c['source'] for c in metadata['chapters']}
actual = {f.relative_to(ROOT).as_posix() for f in (ROOT/'manuscript').glob('*.md')}
if registered != actual:
    raise ValueError('book.json and manuscript files differ')
expected_checkpoints = {c['checkpoint'] for c in metadata['chapters'] if 'checkpoint' in c}
actual_checkpoints = {f.parent.relative_to(ROOT).as_posix() for f in (ROOT/'examples').glob('*/go.mod')}
if expected_checkpoints != actual_checkpoints:
    raise ValueError('book.json and checkpoints differ')
for chapter in metadata['chapters']:
    if 'checkpoint' in chapter:
        module = (ROOT/chapter['checkpoint']/'go.mod').read_text().splitlines()[0]
        if module != 'module '+chapter['module']:
            raise ValueError('Module does not match registry: '+chapter['id'])

def run(args, cwd, expected=0):
    result = subprocess.run(args, cwd=cwd, env=ENV, capture_output=True, text=True, encoding="utf-8", timeout=120)
    records.append({'directory':cwd.relative_to(ROOT).as_posix() if cwd.is_relative_to(ROOT) else 'temporary scenario',
                    'command':args, 'exit_code':result.returncode,
                    'stdout':result.stdout, 'stderr':result.stderr})
    if expected == 0 and result.returncode != 0:
        raise RuntimeError(f'{args}: {result.stdout}\n{result.stderr}')
    if expected != 0 and result.returncode == 0:
        raise RuntimeError(f'Expected failure: {args}')
    return result

version = run(['go', 'version'], ROOT).stdout.strip()
if 'go'+metadata['go_version']+' ' not in version:
    raise RuntimeError('Use Go '+metadata['go_version']+' for this edition; got '+version)
print('Checking reader commands and navigation', flush=True)
reader_commands = audit_reader_commands(ROOT, metadata, run)
# The scanner uses only the standard library, independent of the legacy root module.
with tempfile.TemporaryDirectory(prefix='book-syntax-') as tmp:
    directory = Path(tmp)
    (directory/'go.mod').write_text('module book-syntax\n\ngo '+metadata['go_version']+'\n')
    scanner = directory/('syntaxaudit.exe' if os.name == 'nt' else 'syntaxaudit')
    run(['go', 'build', '-o', str(scanner), str(ROOT/'tools/syntaxaudit.go')], directory)
    syntax=json.loads(run([str(scanner)],ROOT).stdout)
policy=json.loads((ROOT/'editorial/syntax-policy.json').read_text())
for occurrence in syntax:
    rule=policy.get(occurrence['token'])
    if not rule or int(rule['introduced_in']) > int(occurrence['chapter']):
        raise ValueError('New syntax needs an earlier explanation: '+str(occurrence))
(ROOT/'research').mkdir(exist_ok=True)
(ROOT/'research/syntax-inventory.json').write_text(json.dumps(syntax,ensure_ascii=False,indent=2)+'\n')

if 'go'+metadata['go_version']+' ' not in version:
    raise RuntimeError('Use Go '+metadata['go_version']+' for this edition')
for directory in sorted((ROOT/'examples').iterdir()):
    if not (directory/'go.mod').exists(): continue
    print('Checking '+directory.name, flush=True)
    dependencies=run(['go','list','-deps','-test','-f','{{if and (not .Standard) .Module}}{{.Module.Path}}@{{.Module.Version}}{{end}}','./...'],directory).stdout
    own_module=(directory/'go.mod').read_text().splitlines()[0].removeprefix('module ')
    chapter = next(c for c in metadata['chapters'] if c.get('checkpoint') == directory.relative_to(ROOT).as_posix())
    allowed = {own_module+'@'} | {path+'@'+version for path, version in chapter.get('dependencies', {}).items()}
    assert set(dependencies.split()) <= allowed, 'Unreviewed dependency or version in checkpoint: '+directory.name
    formatting = run(['gofmt', '-l', '.'], directory).stdout
    if formatting.strip(): raise RuntimeError(f'Unformatted Go files: {formatting}')
    chapter = next(c for c in metadata['chapters'] if c.get('checkpoint') == directory.relative_to(ROOT).as_posix())
    output = run(['go', 'run', chapter.get('run_target', '.'), *chapter.get('run_args', [])], directory).stdout
    assert output == (directory/'stdout.txt').read_text(), f'Output mismatch: {directory.name}'
    run(['go', 'vet', './...'], directory)
    run(['go', 'test', '-count=1', './...'], directory)
    actual_labs = {p.parent.relative_to(directory).as_posix() for p in (directory/'labs').glob('*/main.go')}
    registered_labs = set(chapter.get('learning_labs', []))
    if actual_labs != registered_labs:
        raise ValueError('Learning lab registry differs from files: '+directory.name)
    for lab in chapter.get('learning_labs', []):
        target = './'+lab
        if target not in chapter.get('additional_run_targets', []):
            raise ValueError('Learning lab is not a documented run target: '+target)
        output = run(['go', 'run', target], directory).stdout
        if output != (directory/lab/'stdout.txt').read_text():
            raise ValueError('Learning lab output differs: '+directory.name+'/'+lab)

# Exercise solutions are checked as changes to temporary checkpoints.
scenarios = [
 ('02-first-program', 'fmt.Println("Форматы: PDF, EPUB, HTML")',
  'fmt.Println("Форматы: PDF, EPUB, HTML")\n\tfmt.Println("Статус: готовится")',
  'Go: от первой строки до интернет-магазина\nФорматы: PDF, EPUB, HTML\nСтатус: готовится\n'),
 ('03-values', '249900', '350050',
  'Go: от первой строки до интернет-магазина\nЦена: 3500.50 KZT\nОпубликована: false\nОпубликована: true\n'),
 ('04-conditions', 'published := true', 'published := false', 'Книга готовится\n'),
 ('04-conditions', 'edition <= 3', 'edition < 3',
  'Доступно издание 1\nИздание 2 снято с продажи\n'),
]
for name, old, new, expected in scenarios:
    with tempfile.TemporaryDirectory(prefix='book-exercise-') as tmp:
        directory = Path(tmp)/'example';shutil.copytree(ROOT/'examples'/name, directory)
        file = directory/'main.go';file.write_text(file.read_text().replace(old,new))
        assert run(['go','run','.'],directory).stdout == expected, name
with tempfile.TemporaryDirectory(prefix='book-test-mutation-') as tmp:
    directory=Path(tmp)/'example';shutil.copytree(ROOT/'examples/05-functions',directory)
    file=directory/'main.go';original=file.read_text()
    file.write_text(original.replace('priceMinor - discountMinor','priceMinor + discountMinor'))
    failure=run(['go','test','-count=1','./...'],directory,expected=1)
    assert 'TestPriceAfterDiscount' in failure.stdout and 'FAIL' in failure.stdout
    file.write_text(original)
    (directory/'exercise_test.go').write_text('''package main
import "testing"
func TestSmallDiscount(t *testing.T) {
    got, ok := priceAfterDiscount(105, 1)
    if !ok || got != 104 { t.Fatalf("got (%d, %t), want (104, true)", got, ok) }
}
''')
    run(['go','test','-count=1','./...'],directory)
with tempfile.TemporaryDirectory(prefix='book-loop-exercise-') as tmp:
    directory=Path(tmp)/'example';shutil.copytree(ROOT/'examples/04-conditions',directory)
    file=directory/'main.go'
    file.write_text(file.read_text().replace('edition <= 3', 'edition <= 4').replace('edition == 2', 'edition == 3'))
    assert run(['go','run','.'],directory).stdout == 'Доступно издание 1\nДоступно издание 2\nИздание 3 снято с продажи\nДоступно издание 4\n'
# Regression: the new Book checkpoint retains earlier business rules.
for chapter in ('07-catalog', '08-books', '09-errors'):
    for name in ('title.go', 'title_test.go'):
        assert (ROOT/'examples/06-strings'/name).read_bytes() == (ROOT/'examples'/chapter/name).read_bytes()
old_pricing=(ROOT/'examples/05-functions/main.go').read_text().split('func priceAfterDiscount',1)[1].split('\nfunc main()',1)[0].strip()
new_pricing=(ROOT/'examples/08-books/pricing.go').read_text().split('func priceAfterDiscount',1)[1].strip()
assert old_pricing == new_pricing
with tempfile.TemporaryDirectory(prefix='book-new-exercises-') as tmp:
    directory=Path(tmp)/'strings';shutil.copytree(ROOT/'examples/06-strings',directory)
    (directory/'exercise_test.go').write_text(r'''package main
import "testing"
func TestInternalSpaces(t *testing.T) {
    got, ok := normalizeTitle("  Go  дүкені  ")
    if !ok || got != "Go  дүкені" { t.Fatalf("got (%q, %t)", got, ok) }
}
''')
    run(['go','test','-count=1','./...'],directory)
    file=directory/'title.go'
    file.write_text(file.read_text().replace('utf8.RuneCountInString(title) > 80','len(title) > 80'))
    failed=run(['go','test','-count=1','./...'],directory,expected=1)
    assert 'TestTitleRuneBoundary' in failed.stdout
    directory=Path(tmp)/'catalog';shutil.copytree(ROOT/'examples/07-catalog',directory)
    file=directory/'main.go';original=file.read_text()
    changed=original.replace('ids = append(ids, "sql-notes")','ids = append(ids, "sql-notes", "go-tests")').replace('"sql-notes": "Заметки о SQL",','"sql-notes": "Заметки о SQL",\n        "go-tests": "Тесты на Go",')
    file.write_text(changed)
    result=run(['go','run','.'],directory)
    assert result.stdout == 'Форматы: [PDF EPUB HTML]\n1 Go: от первой строки до интернет-магазина\n2 Заметки о SQL\n3 Тесты на Go\nНеизвестная книга найдена: false\n'
    file.write_text(changed.replace('"go-tests": "Тесты на Go",',''))
    assert run(['go','run','.'],directory).stdout == 'Форматы: [PDF EPUB HTML]\nКаталог неполон\n'
    directory=Path(tmp)/'books';shutil.copytree(ROOT/'examples/08-books',directory)
    (directory/'exercise_test.go').write_text(r'''package main
import "testing"
func TestFreeBook(t *testing.T) {
    book := Book{ID:"free", Title:"Go", Currency:"KZT"}
    if !book.Publish() || !book.Published { t.Fatal("free book not published") }
}
''')
    run(['go','test','-count=1','./...'],directory)
# Chapter 9: retained model and executable exercise solutions.
for name in ('book.go','pricing.go','title.go'):
    assert (ROOT/'examples/08-books'/name).read_bytes() == (ROOT/'examples/09-errors'/name).read_bytes()
with tempfile.TemporaryDirectory(prefix='book-errors-exercise-') as tmp:
    directory=Path(tmp)/'errors';shutil.copytree(ROOT/'examples/09-errors',directory)
    (directory/'exercise_test.go').write_text(r'''package main
import (
    "bytes"
    "errors"
    "fmt"
    "strings"
    "testing"
)
func TestInvalidTitleAndCloseFailure(t *testing.T) {
    closeErr := errors.New("close failed")
    reader := &trackedReader{reader: strings.NewReader(""), closeErr: closeErr}
    got, err := readTitleAndClose(reader)
    if got != "" || !errors.Is(err, ErrInvalidTitle) || !errors.Is(err, closeErr) || reader.closes != 1 {
        t.Fatalf("got (%q, %v), closes=%d", got, err, reader.closes)
    }
}
func TestDeferredOrderAndArguments(t *testing.T) {
    var out bytes.Buffer
    func() {
        defer fmt.Fprintln(&out, "первый")
        defer fmt.Fprintln(&out, "второй")
        fmt.Fprintln(&out, "работа")
    }()
    if out.String() != "работа\nвторой\nпервый\n" { t.Fatal(out.String()) }
    out.Reset()
    func() {
        n := 1
        defer fmt.Fprintln(&out, n)
        n = 2
    }()
    if out.String() != "1\n" { t.Fatal(out.String()) }
    out.Reset()
    func() {
        n := 1
        defer func() { fmt.Fprintln(&out, n) }()
        n = 2
    }()
    if out.String() != "2\n" { t.Fatal(out.String()) }
}
''')
    run(['go','test','-count=1','./...'],directory)
    file=directory/'file.go'
    file.write_text(file.read_text().replace('errors.Join(err, closeErr)','closeErr'))
    failure=run(['go','test','-count=1','./...'],directory,expected=1)
    assert 'TestInvalidTitleAndCloseFailure' in failure.stdout

repro=run(['go','run','.'],ROOT/'research/reproductions/map-value')
assert repro.stdout == 'Go\nЗамена поля сохранилась: false\n'

# A single source for every included listing; reject invalid references.
pattern=re.compile(r'<!-- include:([^ >]+) -->')
def include(match):
    ref=match[1];path, _, region=ref.partition('#')
    file=(ROOT/path).resolve()
    if not file.is_relative_to(ROOT): raise ValueError(ref)
    code=file.read_text().rstrip()
    if region:
        if region != 'first-test': raise ValueError(region)
        code=code.split('\nfunc TestNoDiscount',1)[0].rstrip()
    language={'.go':'go', '.html':'html', '.css':'css', '.json':'json'}.get(file.suffix, 'text')
    return f'```{language}\n{code}\n```'
rendered=ROOT/'build/reading';rendered.mkdir(parents=True,exist_ok=True)
for file in sorted((ROOT/'manuscript').glob('*.md')):
    text = file.read_text()
    if any(ord(c) < 32 and c not in '\n\t\r' for c in text):
        raise ValueError(f'Unexpected control character: {file.name}')
    (rendered/file.name).write_text(pattern.sub(include,text).replace('../assets/', '../../assets/'))
if source_hashes() != initial_sources or (ROOT/'book.json').read_bytes() != initial_registry:
    raise RuntimeError('Sources changed during verification; rerun on a stable working tree')
report={'checked_at':datetime.now(timezone.utc).isoformat(),'go_version':version,
        'status':'passed','commands':records,
        'platform':platform.platform(),
        'checked_chapters':[c['id'] for c in metadata['chapters']],
        'source_sha256':initial_sources, 'reader_commands':reader_commands,
        'limits':['This report records only the current platform; remote CI must be confirmed separately.',
                  'Checks cover all registered chapters; they do not prove pedagogical effectiveness or production readiness.']}
(ROOT/'research').mkdir(exist_ok=True)
(ROOT/'research/verification.json').write_text(json.dumps(report,ensure_ascii=False,indent=2)+'\n')
print(f'PASS: {version}; {len(records)} commands; rendered {len(list(rendered.glob("*.md")))} chapters')
