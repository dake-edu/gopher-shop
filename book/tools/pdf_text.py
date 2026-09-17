"""Keep instructional symbols visible when the body font lacks their glyphs."""
import html
from reportlab.pdfbase import pdfmetrics


def escaped_text(text):
    regular = pdfmetrics.getFont('Noto').face.charToGlyph
    bold = pdfmetrics.getFont('NotoBold').face.charToGlyph
    fallback = pdfmetrics.getFont('NotoMono').face.charToGlyph
    parts = []
    for character in text or '':
        escaped = html.escape(character)
        point = ord(character)
        if character.isspace() or point in regular and point in bold:
            parts.append(escaped)
        elif point in fallback:
            parts.append('<font name="NotoMono">'+escaped+'</font>')
        else:
            raise ValueError(f'PDF fonts do not contain U+{point:04X}: {character!r}')
    return ''.join(parts)
