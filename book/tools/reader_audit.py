"""Check runnable commands printed in the manuscript, without executing shell text."""
from pathlib import Path
import re
import shlex


def commands(text):
    """Fenced commands plus complete inline test commands; prose is not a shell script."""
    found = []
    for block in re.finditer(r'```(?:sh|shell|bash|powershell)\n(.*?)```', text, re.S):
        for line in block[1].splitlines():
            if line.strip() and not line.lstrip().startswith('#'):
                found.append(line.strip())
    # Inline negative examples such as `go run .` after moving main are intentional.
    # Inline test commands, unlike such experiments, promise actual selected tests.
    for match in re.finditer(r'(?<!`)`(go test [^`\n]+)`(?!`)', text):
        if re.search(r'(?:^|\s)\.(?:/[^\s]+)?(?:$|\s)', match[1]):
            found.append(match[1])
    return list(dict.fromkeys(found))


def flag_value(argv, name):
    for index, word in enumerate(argv):
        if word.startswith(name+'='):
            return word.split('=', 1)[1]
        if word == name:
            if index+1 == len(argv) or argv[index+1].startswith('-'):
                raise ValueError('Missing test flag value')
            return argv[index+1]
    return None


def test_arguments(argv):
    packages = []
    values = ('-run', '-count', '-bench', '-benchtime', '-fuzz', '-fuzztime',
              '-parallel', '-cpuprofile', '-memprofile', '-trace')
    index = 2
    while index < len(argv):
        word = argv[index]
        name = word.split('=', 1)[0]
        if word in ('-v', '-race', '-benchmem'):
            pass
        elif name in values:
            flag_value(argv, name)
            if '=' not in word:
                index += 1
        elif word.startswith('-'):
            raise ValueError('Explain and support new test flag: '+word)
        else:
            packages.append(word)
        index += 1
    return packages or ['.'], flag_value(argv, '-run')


def package_path(directory, target):
    if target == './...':
        return
    path = (directory / target).resolve()
    if not path.is_relative_to(directory.resolve()) or not path.is_dir():
        raise ValueError('Missing local package: '+target)
    if not any(path.glob('*.go')):
        raise ValueError('No Go files in package: '+target)


def audit_menu(directory):
    """Check literal navigation URLs against this checkpoint's literal GET routes."""
    header = directory/'web/templates/partials/header.html'
    if not header.exists():
        return
    source = '\n'.join(p.read_text(encoding='utf-8')
                       for p in (directory/'internal/web').glob('*.go')
                       if not p.name.endswith('_test.go'))
    routes = {p.replace('{$}', '') for p in re.findall(r'HandleFunc\("GET ([^"\n]+)"', source)}
    for href in re.findall(r'href="(/[^"]*)"', header.read_text(encoding='utf-8')):
        if href not in routes:
            raise ValueError('Menu links to unavailable chapter route: '+str(directory)+': '+href)


def audit(root, metadata, run):
    rules = {
        'version': ('01', 'Установите Go'),
        'mod init': ('01', 'Дайте проекту имя'),
        'run': ('02', 'Запуск и сборка'),
        'build': ('02', 'Запуск и сборка'),
        'test': ('05', 'Проверка, которая остаётся'),
        'vet': ('10', 'Проверка без догадок'),
        'doc': ('10', 'Проверка без догадок'),
        'tool pprof': ('36', 'Профиль показывает, где искать'),
        'tool trace': ('37', 'Трасса показывает ожидание'),
        'mod verify': ('39', 'Зависимости и инструменты сборки'),
        'install': ('39', 'Уязвимость модуля и достижимый вызов'),
        'get': ('16', 'Почему здесь появляется зависимость'),
        'mod tidy': ('16', 'Почему здесь появляется зависимость'),
    }
    by_id = {c['id']: c for c in metadata['chapters']}
    report = []
    for chapter in metadata['chapters']:
        # The preface previews the route and explicitly assigns commands to future chapters.
        if chapter['id'] == '00':
            continue
        text = (root/chapter['source']).read_text(encoding='utf-8')
        directory = root/chapter.get('checkpoint', '.')
        if 'checkpoint' in chapter:
            audit_menu(directory)
            for reference in set(re.findall(r'`((?:internal|cmd|web)/[\w./-]+\.(?:go|html|css))`', text)):
                if not (directory/reference).is_file():
                    raise ValueError('Missing file in chapter '+chapter['id']+': '+reference)
        for line in commands(text):
            argv = shlex.split(line)
            if not argv or argv[0] != 'go':
                continue
            action = ' '.join(argv[1:3]) if argv[1] in ('mod', 'tool') else argv[1]
            if action not in rules:
                raise ValueError('New command needs an introduction: '+line)
            introduced, heading = rules[action]
            introduction = (root/by_id[introduced]['source']).read_text(encoding='utf-8')
            if chapter['id'] < introduced or '\n## '+heading+'\n' not in introduction:
                raise ValueError('Command precedes its introduction in '+chapter['id']+': '+line)
            if action == 'test':
                targets, pattern = test_arguments(argv)
                for target in targets:
                    package_path(directory, target)
                for flag, first in [('-run', '09'), ('-v', '09'), ('-race', '30'),
                                    ('-fuzz', '35'), ('-fuzztime', '35'), ('-parallel', '35'),
                                    ('-bench', '36'), ('-benchmem', '36'), ('-benchtime', '36'),
                                    ('-cpuprofile', '36'), ('-memprofile', '36'), ('-trace', '37')]:
                    if any(a.split('=')[0] == flag for a in argv) and chapter['id'] < first:
                        raise ValueError('Test flag precedes its explanation: '+line)
                special = [(name, flag_value(argv, name)) for name in ('-bench', '-fuzz')]
                intentional_skip = pattern == '^$' and any(value for _, value in special)
                if pattern is not None and not intentional_skip:
                    # Go's own regexp engine decides which tests match; no tests are executed here.
                    result = run(['go', 'test', '-list', pattern, *targets], directory)
                    if not any(s.startswith(('Test', 'Example', 'Fuzz')) for s in result.stdout.splitlines()):
                        raise ValueError('Command selects no tests: '+line)
                for name, selected in special:
                    if selected is not None:
                        result = run(['go', 'test', '-list', selected, *targets], directory)
                        prefix = 'Benchmark' if name == '-bench' else 'Fuzz'
                        matches = [s for s in result.stdout.splitlines() if s.startswith(prefix)]
                        if not matches or (name == '-fuzz' and len(matches) != 1):
                            raise ValueError('Command selects no unique matching target: '+line)
            elif action == 'install':
                if argv[2:] != ['golang.org/x/vuln/cmd/govulncheck@v1.8.0']:
                    raise ValueError('Unreviewed tool installation: '+line)
            elif action == 'tool trace':
                if argv[3:] != ['trace.out']:
                    raise ValueError('Unexpected trace input: '+line)
            elif action == 'tool pprof':
                if argv[3:-1] not in (['-top'], ['-top', '-alloc_space']) or argv[-1] not in ('cpu.out', 'mem.out'):
                    raise ValueError('Explain and support profiling arguments: '+line)
            elif action == 'run':
                target = argv[2]
                package_path(directory, target)
                if target not in [chapter.get('run_target', '.'), *chapter.get('additional_run_targets', [])]:
                    raise ValueError('Run target differs from checkpoint: '+line)
                main = '\n'.join(p.read_text(encoding='utf-8') for p in (directory/target).glob('*.go'))
                flags = set(re.findall(r'flag\.(?:Bool|String|Int)\("([^"]+)"', main)) | {'h', 'help'}
                for argument in argv[3:]:
                    if argument.startswith('-') and argument.lstrip('-').split('=')[0] not in flags:
                        raise ValueError('Unknown shop flag in manuscript: '+line)
            elif action == 'build':
                package_path(directory, argv[-1])
            elif action in ('vet', 'doc'):
                package_path(directory, argv[2])
            elif action == 'get':
                allowed = {key+'@'+value for key, value in chapter.get('dependencies', {}).items()}
                if not set(argv[2:]) <= allowed:
                    raise ValueError('Dependency command differs from registry: '+line)
            report.append({'chapter': chapter['id'], 'command': line})
    return report
