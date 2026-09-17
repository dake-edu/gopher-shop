"""Check only the approved marketing illustration shipped in the sample."""
import hashlib
import json

def audit(root):
    folder = root/'assets/screenshots'
    capture = json.loads((folder/'capture.json').read_text())
    assert capture['kind'] == 'approved-sample-images'
    assert {i['file'] for i in capture['screenshots']} == {p.name for p in folder.glob('*.png')}
    for item in capture['screenshots']:
        assert hashlib.sha256((folder/item['file']).read_bytes()).hexdigest() == item['sha256']
    return len(capture['screenshots'])
