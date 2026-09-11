#!/usr/bin/env python3
"""Inventory source editions without modifying original manuscripts."""
import os as _os
import sys as _sys
if not _sys.flags.utf8_mode:
    _os.environ['PYTHONUTF8'] = '1'
    _os.execv(_sys.executable, [_sys.executable, '-X', 'utf8', *_sys.argv])

from pathlib import Path
from zipfile import ZipFile
from collections import Counter
from difflib import SequenceMatcher
import xml.etree.ElementTree as ET
import json
import re

BOOK = Path(__file__).resolve().parents[1]
SOURCE = Path('/Users/dake/go/src/The Gopher Shop')
NS = {'w': 'http://schemas.openxmlformats.org/wordprocessingml/2006/main'}
texts = {}
for index, path in enumerate(sorted(SOURCE.glob('Adaptive*.docx')), 1):
    with ZipFile(path) as archive:
        root = ET.fromstring(archive.read('word/document.xml'))
    paragraphs = [''.join(t.text or '' for t in p.findall('.//w:t', NS))
                  for p in root.findall('.//w:p', NS)]
    text = '\n'.join(paragraphs)
    texts[path.name] = text
    (BOOK / 'research/extracted' / f'edition-{index}.txt').write_text(text)
report = ['# Сравнение редакций The Gopher Shop', '',
          'Извлечён текст абзацев DOCX, включая абзацы таблиц. Рисунки, сноски, '
          'комментарии и визуальная вёрстка этим сравнением не проверяются.', '',
          '| Редакция | Символы | Непустые строки | Повторные строки длиннее 80 символов |',
          '|---|---:|---:|---:|']
for name, text in texts.items():
    lines = [line.strip() for line in text.splitlines() if line.strip()]
    counts = Counter(line for line in lines if len(line) > 80)
    repeats = sum(n - 1 for n in counts.values() if n > 1)
    report.append(f'| {name} | {len(text)} | {len(lines)} | {repeats} |')
report += ['', '## Сходство последовательностей абзацев', '',
           'Коэффициент ниже описывает сходство текста, а не качество редакции.', '']
names = list(texts)
for i, left in enumerate(names):
    for right in names[i+1:]:
        ratio = SequenceMatcher(None, texts[left].splitlines(), texts[right].splitlines(), autojunk=False).ratio()
        report.append(f'- {left} ↔ {right}: {ratio:.3f}.')
report += ['', '## Решение', '',
           'Основной материал для последовательного чтения: book_chapters/chapter_02.txt–chapter_21.txt '
           '(главы 1–20). chapter_01.txt — вступление. assignments.md читать отдельно. '
           'DOCX использовать как дополнительные редакции, проверяя уникальные фрагменты. '
           'Самый большой файл не назначать эталоном только по размеру.', '',
           'В старой рукописи два заголовка «Глава 19» встречаются в блоках плана '
           '(Dockerfile и docker-compose); при переносе отделять план от основного оглавления.', '',
           '## Покрытие локального курса', '']
lessons = Path('/Users/dake/go/src/shanraq.org/course/lessons/go')
for path in sorted(lessons.glob('*.md')):
    if path.stem.endswith(('-en', '-kz')):
        continue
    text = path.read_text()
    headings = re.findall(r'^## (.+)$', text, re.M)
    report.append(f'- `{path.name}`: {text.splitlines()[0][2:]}. Разделов: {len(headings)}.')
report += ['', 'Этот документ — инвентаризация, не утверждение о полном техническом аудите каждой строки.']
(BOOK / 'editorial/edition-comparison.md').write_text('\n'.join(report) + '\n')
print('\n'.join(report[:20]))
