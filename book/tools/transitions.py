"""Derive exact checkpoint file changes for the reader's source bundle."""
import re


def snapshot(directory):
    return {p.relative_to(directory).as_posix(): p.read_bytes()
            for p in directory.rglob('*') if p.is_file()}


def normalized(data):
    return re.sub(rb'github.com/dake-edu/gopher-shop/book/examples/\d\d-[a-z-]+',
                  b'CHECKPOINT_MODULE', data)


def render_transitions(root, metadata):
    lines = ['# Точные изменения контрольных точек', '',
             'Пути ниже считаются от папки соответствующей главы в book/examples.',
             'Сначала прочитайте объяснение главы. Это перечень различий исходников, а не команда удаления вашей работы.',
             'Сохраните предыдущую папку; задания выполняйте в отдельной копии. Базы, ключи и покупки не переносите автоматически.',
             'В каждой новой контрольной точке свой путь module в go.mod и такое же начало локальных импортов.',
             'Изменения только этого префикса не перечисляются как изменения поведения. Остальные файлы перечислены полностью.', '']
    previous = {}
    for chapter in metadata['chapters']:
        if 'checkpoint' not in chapter:
            continue
        current = snapshot(root/chapter['checkpoint'])
        lines += ['## '+chapter['title'], '', '`book/'+chapter['checkpoint']+'`', '']
        groups = [('Создать', sorted(current.keys()-previous.keys())),
                  ('Заменить содержимое', sorted(k for k in current.keys() & previous.keys()
                                                if normalized(current[k]) != normalized(previous[k]))),
                  ('Не переносить из прежней папки', sorted(previous.keys()-current.keys()))]
        for label, paths in groups:
            if paths:
                lines += [label+':', ''] + ['- `'+path+'`' for path in paths] + ['']
        previous = current
    return '\n'.join(lines)+'\n'
