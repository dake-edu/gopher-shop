from pathlib import Path
import unittest
from reportlab.pdfbase import pdfmetrics
from reportlab.pdfbase.ttfonts import TTFont

from pdf_text import escaped_text


class PDFFontTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        directory = Path(__file__).resolve().parents[1]/'assets/fonts'
        for name, file in [('Noto', 'NotoSans-Regular.ttf'), ('NotoBold', 'NotoSans-Bold.ttf'),
                           ('NotoMono', 'NotoSansMono-Regular.ttf')]:
            pdfmetrics.registerFont(TTFont(name, str(directory/file)))

    def test_arrows_use_an_available_glyph(self):
        self.assertNotIn(ord('→'), pdfmetrics.getFont('Noto').face.charToGlyph)
        self.assertIn(ord('→'), pdfmetrics.getFont('NotoMono').face.charToGlyph)
        self.assertEqual(escaped_text('A → B'), 'A <font name="NotoMono">→</font> B')

    def test_fallback_preserves_markup_escaping(self):
        self.assertEqual(escaped_text('<script>&'), '&lt;script&gt;&amp;')

    def test_missing_glyph_is_not_silently_dropped(self):
        with self.assertRaisesRegex(ValueError, 'PDF fonts do not contain'):
            escaped_text('\U0010ffff')


if __name__ == '__main__':
    unittest.main()
