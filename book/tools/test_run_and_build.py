"""Execute the beginner's experiment with saved source and a stale binary."""
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


class RunAndBuildTests(unittest.TestCase):
    def test_saved_edit_requires_new_build(self):
        example = Path(__file__).resolve().parents[1]/'examples/02-first-program'
        with tempfile.TemporaryDirectory(prefix='book-run-build-') as tmp:
            directory = Path(tmp)/'example'
            shutil.copytree(example, directory)
            def run(argv):
                result = subprocess.run(argv, cwd=directory, env=dict(os.environ, GOWORK='off'),
                                        capture_output=True, text=True, encoding='utf-8', timeout=120)
                self.assertEqual(result.returncode, 0, result.stdout+result.stderr)
                return result.stdout
            name = 'bookshop.exe' if os.name == 'nt' else 'bookshop'
            executable = str(directory/name)
            run(['go', 'build', '-o', name, '.'])
            original = run([executable])
            source = directory/'main.go'
            source.write_text(source.read_text(encoding='utf-8').replace('Форматы:', 'Доступно:'), encoding='utf-8')
            self.assertEqual(run([executable]), original)
            edited = run(['go', 'run', '.'])
            self.assertNotEqual(edited, original)
            self.assertIn('Доступно:', edited)
            run(['go', 'build', '-o', name, '.'])
            self.assertEqual(run([executable]), edited)


if __name__ == '__main__':
    unittest.main()
