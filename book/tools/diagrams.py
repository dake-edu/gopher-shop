"""One scene definition rendered as SVG and as native PDF vector shapes."""
import html
import math
from reportlab.graphics.shapes import Drawing, Rect, Line, Polygon, String
from reportlab.lib import colors

INK='#125d63'
PALE='#eaf0ef'
def render_svg(scene):
    width,height=scene['width'],scene['height']
    title=html.escape(scene['title'])
    parts=[f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}" role="img" aria-labelledby="title desc">',
           f'<title id="title">{title}</title><desc id="desc">{html.escape(scene["note"])}</desc>',
           '<defs><marker id="arrow" markerWidth="8" markerHeight="8" refX="7" refY="4" orient="auto"><path d="M0 0 L8 4 L0 8 Z" fill="#125d63"/></marker></defs>',
           f'<rect width="{width}" height="{height}" fill="#faf9f5"/>',
           f'<text x="{width/2}" y="30" text-anchor="middle" font-family="sans-serif" font-size="20" font-weight="bold" fill="{INK}">{title}</text>']
    for box in scene['boxes']:
        x,y,w,h=box['x'],box['y'],box['w'],box['h']
        parts.append(f'<rect x="{x}" y="{y}" width="{w}" height="{h}" rx="8" fill="{PALE}" stroke="{INK}" stroke-width="2"/>')
        start=y+h/2-(len(box['lines'])-1)*11
        for i,text in enumerate(box['lines']):
            parts.append(f'<text x="{x+w/2}" y="{start+i*22}" text-anchor="middle" dominant-baseline="middle" font-family="sans-serif" font-size="15" fill="{INK}">{html.escape(text)}</text>')
    for x1,y1,x2,y2 in scene['arrows']:
        parts.append(f'<line x1="{x1}" y1="{y1}" x2="{x2}" y2="{y2}" stroke="{INK}" stroke-width="2" marker-end="url(#arrow)"/>')
    parts.append(f'<text x="{width/2}" y="{height-15}" text-anchor="middle" font-family="sans-serif" font-size="14" fill="{INK}">{html.escape(scene["note"])}</text></svg>')
    return '\n'.join(parts)+'\n'

def drawing(scene):
    width,height=scene['width'],scene['height']
    result=Drawing(width,height)
    ink=colors.HexColor(INK)
    result.add(Rect(0,0,width,height,fillColor=colors.HexColor('#faf9f5'),strokeColor=None))
    result.add(String(width/2,height-30,scene['title'],fontName='NotoBold',fontSize=20,textAnchor='middle',fillColor=ink))
    for box in scene['boxes']:
        x,y,w,h=box['x'],box['y'],box['w'],box['h']
        result.add(Rect(x,height-y-h,w,h,rx=8,ry=8,fillColor=colors.HexColor(PALE),strokeColor=ink,strokeWidth=2))
        start=y+h/2-(len(box['lines'])-1)*11
        for i,text in enumerate(box['lines']):
            result.add(String(x+w/2,height-(start+i*22)-5,text,fontName='Noto',fontSize=15,textAnchor='middle',fillColor=ink))
    for x1,y1,x2,y2 in scene['arrows']:
        y1,y2=height-y1,height-y2
        result.add(Line(x1,y1,x2,y2,strokeColor=ink,strokeWidth=2))
        angle=math.atan2(y2-y1,x2-x1)
        points=[x2,y2]
        for delta in (.45,-.45):
            points.extend([x2-10*math.cos(angle+delta),y2-10*math.sin(angle+delta)])
        result.add(Polygon(points,fillColor=ink,strokeColor=None))
    result.add(String(width/2,15,scene['note'],fontName='Noto',fontSize=14,textAnchor='middle',fillColor=ink))
    return result
