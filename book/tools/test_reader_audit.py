"""Regression cases for gaps missed by merely testing complete checkpoints."""
from pathlib import Path
import subprocess
import tempfile
import unittest

from reader_audit import audit, audit_menu


class ReaderCommandTests(unittest.TestCase):
    def test_menu_does_not_promise_a_future_chapter(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            (root/'web/templates/partials').mkdir(parents=True)
            (root/'internal/web').mkdir(parents=True)
            header = root/'web/templates/partials/header.html'
            header.write_text('<a href="/">Home</a><a href="/login">Login</a>', encoding='utf-8')
            source = root/'internal/web/web.go'
            source.write_text('mux.HandleFunc("GET /{$}", home)', encoding='utf-8')
            with self.assertRaisesRegex(ValueError, 'unavailable chapter route'):
                audit_menu(root)
            source.write_text(source.read_text()+'\nmux.HandleFunc("GET /login", login)', encoding='utf-8')
            audit_menu(root)

    def fixture(self, root, command):
        (root/'intro.md').write_text('Title\n## Проверка, которая остаётся\nExplanation\n', encoding='utf-8')
        (root/'lesson.md').write_text('```sh\n'+command+'\n```\n', encoding='utf-8')
        package = root/'example/cmd/shop'
        package.mkdir(parents=True)
        (package/'main.go').write_text('package main\n', encoding='utf-8')
        return {'chapters': [
            {'id': '05', 'source': 'intro.md'},
            {'id': '24', 'source': 'lesson.md', 'checkpoint': 'example'}]}

    def test_extra_command_requires_explicit_registration(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            meta = self.fixture(root, 'go run ./cmd/package -edition v1')
            (root/'intro.md').write_text('Title\n## Запуск и сборка\nExplanation\n', encoding='utf-8')
            meta['chapters'][0]['id'] = '02'
            package = root/'example/cmd/package'
            package.mkdir()
            (package/'main.go').write_text('package main\nvar edition = flag.String("edition", "", "edition")\n', encoding='utf-8')
            with self.assertRaisesRegex(ValueError, 'Run target differs'):
                audit(root, meta, lambda *_: None)
            meta['chapters'][1]['additional_run_targets'] = ['./cmd/package']
            self.assertEqual(len(audit(root, meta, lambda *_: None)), 1)

    def test_old_root_target_is_rejected(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            meta = self.fixture(root, 'go test . -run TestNotificationSender -v')
            with self.assertRaisesRegex(ValueError, 'No Go files'):
                audit(root, meta, lambda *_: self.fail('must reject before executing Go'))

    def test_empty_filter_is_not_a_success(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            meta = self.fixture(root, 'go test ./cmd/shop -run MisspelledTest -v')
            with self.assertRaisesRegex(ValueError, 'selects no tests'):
                audit(root, meta, lambda *_: subprocess.CompletedProcess([], 0, 'ok example [no tests to run]\n'))

    def test_actual_selected_test_is_accepted(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            meta = self.fixture(root, "go test ./cmd/shop -run 'TestOne|TestTwo' -v")
            calls = []
            def run(argv, cwd):
                calls.append(argv)
                return subprocess.CompletedProcess(argv, 0, 'TestOne\nTestTwo\nok example\n')
            self.assertEqual(len(audit(root, meta, run)), 1)
            self.assertEqual(calls, [['go', 'test', '-list', 'TestOne|TestTwo', './cmd/shop']])

    def test_missing_fuzz_target_is_rejected(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            meta = self.fixture(root, "go test ./cmd/shop -run '^$' -fuzz FuzzTypo -fuzztime 10s -parallel 2")
            meta['chapters'][-1]['id'] = '35'
            with self.assertRaisesRegex(ValueError, 'matching target'):
                audit(root, meta, lambda *_: subprocess.CompletedProcess([], 0, 'ok example\n'))

    def test_benchmark_can_skip_regular_tests(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            meta = self.fixture(root, "go test ./cmd/shop -run '^$' -bench '^BenchmarkWork$' -benchmem -count 5 -benchtime 100ms")
            meta['chapters'][-1]['id'] = '36'
            calls = []
            def run(argv, cwd):
                calls.append(argv)
                return subprocess.CompletedProcess(argv, 0, 'BenchmarkWork\nok example\n')
            self.assertEqual(len(audit(root, meta, run)), 1)
            self.assertEqual(calls, [['go', 'test', '-list', '^BenchmarkWork$', './cmd/shop']])


if __name__ == '__main__':
    unittest.main()
