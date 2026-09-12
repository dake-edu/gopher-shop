"""Static lexical highlighting shared by HTML, EPUB and PDF.

Names deliberately share a colour: a lexer cannot infer variable/function roles
from every short fragment. Highlighting never validates a program.
"""
import xml.etree.ElementTree as ET
import re
from pygments.lexers import get_lexer_by_name
from pygments.token import Keyword, Name, String, Number, Comment, Operator, Punctuation

PALETTE = {
    'keyword': '#402B56',
    'name': '#00758D',
    'string': '#166534',
    'number': '#9C3B10',
    'comment': '#555759',
    'symbol': '#555759',
}

def category(token):
    # Go predeclared types/constants are identifiers, not reserved keywords.
    if token in Keyword.Type or token in Keyword.Constant:
        return 'name'
    for family, name in ((Comment, 'comment'), (String, 'string'),
                         (Number, 'number'), (Keyword, 'keyword'),
                         (Name, 'name'), (Operator, 'symbol'),
                         (Punctuation, 'symbol')):
        if token in family:
            return name
    return None

def fragments(source, language='go'):
    lexer = get_lexer_by_name(language, stripnl=False, ensurenl=False)
    # The unprocessed API preserves tabs, newlines and incomplete inline fragments.
    result = [(category(token), value)
              for _, token, value in lexer.get_tokens_unprocessed(source)]
    if ''.join(value for _, value in result) != source:
        raise ValueError('Highlighting must preserve the source exactly')
    return result

def highlight_tree(root):
    block_codes = {node for pre in root.iter('pre') for node in pre.iter('code')}
    for node in list(root.iter('code')):
        source = ''.join(node.itertext())
        if node in block_codes:
            language = node.get('class', '').removeprefix('language-')
            if language not in ('go', 'sh', 'bash', 'powershell'):
                continue  # Output and plain text are not source code.
        else:
            language = 'go'
            if source.startswith(('go ', 'gofmt ', './', '.\\')):
                language = 'sh'
            elif source.startswith(('examples/', 'book/', 'manuscript/', 'research/', 'https://', '../')) or re.fullmatch(r'[\w.-]+\.(?:go|mod|sum|md|txt|exe|json)', source):
                continue  # Paths and URLs are not Go expressions.
        node.text = None
        for child in list(node):
            node.remove(child)
        for kind, value in fragments(source, language):
            span = ET.SubElement(node, 'span')
            if kind:
                span.set('class', 'tok-' + kind)
            span.text = value
    return root

def css():
    return '\n'.join('.tok-' + kind + '{color:' + colour + '}'
                     for kind, colour in PALETTE.items()) + '\n'
