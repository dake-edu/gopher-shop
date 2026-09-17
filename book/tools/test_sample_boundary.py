"""Ensure an accidentally copied paid chapter cannot enter the public sample."""
import json
from pathlib import Path
import tempfile
import unittest
from sample_boundary import audit

class SampleBoundaryTests(unittest.TestCase):
    def test_extra_manuscript_is_rejected(self):
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            (root/'manuscript').mkdir()
            (root/'examples').mkdir()
            chapters = [{'id': f'{i:02}', 'source': f'manuscript/{i:02}-chapter.md'} for i in range(13)]
            for c in chapters:
                (root/c['source']).write_text('Sample')
            (root/'book.json').write_text(json.dumps({'sample': {'last_chapter':12}, 'chapters':chapters}))
            audit(root)
            (root/'manuscript/13-paid.md').write_text('Must not ship')
            with self.assertRaises(AssertionError):
                audit(root)
