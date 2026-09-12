"""Reader-facing guarantees: source fidelity, consistent colours and contrast."""
import unittest
import xml.etree.ElementTree as ET
from highlighting import PALETTE, fragments, highlight_tree

class HighlightingTests(unittest.TestCase):
    def test_keyword_and_string_have_same_colours_inline_and_in_listing(self):
        root = ET.fromstring('<root><p><code>import "fmt"</code></p>'
                             '<pre><code class="language-go">import "fmt"</code></pre></root>')
        highlight_tree(root)
        inline, block = list(root.iter('code'))
        self.assertEqual(ET.tostring(inline).replace(b'<code>', b'<code class="language-go">'),
                         ET.tostring(block))
        self.assertEqual([(s.get('class'), s.text) for s in inline],
                         [('tok-import', 'import'), (None, ' '), ('tok-string', '"fmt"')])

    def test_go_keywords_and_braces_have_distinct_colours(self):
        expected = [('package', 'package'), ('import', 'import'), ('func', 'func'),
                    ('keyword', 'return'), ('brace', '{'), ('brace', '}'),
                    ('bracket', '('), ('bracket', ')')]
        actual = [(kind, value) for kind, value in fragments('package import func return {}()')
                  if value.strip()]
        self.assertEqual(actual, expected)
        self.assertEqual(len({PALETTE[kind] for kind in ('package', 'import', 'func', 'keyword', 'brace')}), 5)

    def test_keyword_spelling_in_strings_and_comments_keeps_its_context(self):
        self.assertTrue(all(kind == 'string' for kind, value in fragments('"package import func {}"')))
        self.assertTrue(all(kind == 'comment' for kind, value in fragments('// package import func {}')))

    def test_unicode_tabs_comments_escapes_and_incomplete_fragments_survive(self):
        for source in ['\t// ӘGo <>&\nfmt.Println("ӘGo\\n")\n', 'x := 10\n', '"',
                       'func main() {', 'price < limit && ok', '/* comment\n */', '']:
            with self.subTest(source=source):
                self.assertEqual(''.join(value for _, value in fragments(source)), source)

    def test_output_and_paths_are_not_presented_as_go(self):
        root = ET.fromstring('<root><pre><code class="language-text">false 42 import</code></pre>'
                             '<p><code>examples/08-books</code><code>go.mod</code><code>main.go</code></p></root>')
        highlight_tree(root)
        self.assertFalse(list(root.iter('span')))

    def test_predeclared_go_names_are_not_reserved_keywords(self):
        parts = fragments('int64 true false nil len string any comparable')
        self.assertTrue(all(kind == 'name' for kind, value in parts if value.strip()))

    def test_text_colours_are_legible_on_white(self):
        for colour in PALETTE.values():
            channels = [int(colour[i:i+2], 16) / 255 for i in (1, 3, 5)]
            linear = [c/12.92 if c <= .04045 else ((c+.055)/1.055)**2.4 for c in channels]
            luminance = sum(c*w for c, w in zip(linear, (.2126, .7152, .0722)))
            self.assertGreaterEqual(1.05/(luminance+.05), 4.5, colour)

if __name__ == '__main__':
    unittest.main()
