"""Publication data rendered once for all reading formats."""
import html

def front_pages(meta):
    p = meta['publication']
    esc = html.escape
    last = max(int(c['id']) for c in meta['chapters'])
    title = esc(meta['title'])
    author = esc(p['author'])
    date = esc(p['date_label'])
    publisher = esc(p['publisher'])
    edition = esc(meta['edition'])
    title_body = (f'<p class="author">{author}</p><h1>{title}</h1>'
                  f'<p class="subtitle">{esc(p["subtitle"])}</p>'
                  f'<p class="levels">{esc(p["levels"])}</p>'
                  f'<p class="edition-note">Рабочая редакция {edition}</p>'
                  f'<p class="imprint">{publisher}<br />{date}</p>')
    details = (f'<h1>Выходные сведения</h1>'
               f'<p><strong>{author}</strong><br />{title}: практическая книга. '
               f'— {publisher}, {p["date"][:4]}. — Электронное издание. '
               f'Рабочая редакция {edition}.</p>'
               f'<p>{esc(p["description"])}</p>'
               f'<p>{esc(p["audience"])}</p>'
               f'<p><strong>Автор:</strong> {author}.<br />'
               f'{esc(p["author_bio"])}<br />'
               f'<strong>Издательский проект:</strong> {publisher}.<br />'
               f'<strong>Дата этой редакции:</strong> {date}.<br />'
               '<strong>Язык:</strong> русский.<br />'
               '<strong>Форматы:</strong> PDF, EPUB, HTML.</p>'
               f'<p>Этот файл содержит предисловие и главы 1–{last}. Полная книга ещё готовится. '
               'Дата относится к рабочей редакции, а не к завершённому коммерческому выпуску. '
               'ISBN и библиотечные классификационные индексы в этой редакции не указаны.</p>'
               f'<p>© {p["date"][:4]} {author}.<br />'
               'Обложка предоставлена автором. Подготовка рукописи и учебных материалов '
               'ведётся с использованием инструментов ИИ.</p>'
               '<p>Лицензия исходного репозитория — GNU AGPL-3.0; '
               '<a href="https://github.com/dake-edu/gopher-shop/blob/main/LICENSE">текст лицензии</a>. '
               'Шрифты Noto распространяются по '
               '<a href="https://openfontlicense.org/">SIL Open Font License 1.1</a>; '
               'копия лицензии включена в электронный комплект.</p>'
               f'<p><strong>Сайт:</strong> <a href="https://shanraq.org">shanraq.org</a>.<br />'
               f'<strong>Код книги:</strong> <a href="{esc(meta["repository"], quote=True)}">'
               'github.com/dake-edu/gopher-shop</a>.</p>')
    return [
        {'stem':'index', 'title':'Обложка', 'kind':'cover',
         'body':f'<img class="book-cover" src="cover/go-book-cover.jpg" alt="Обложка книги «{title}»: голубой персонаж с книгой в тележке" />'},
        {'stem':'title-page', 'title':'Титульный лист', 'kind':'titlepage', 'body':title_body},
        {'stem':'publication-details', 'title':'Выходные сведения', 'kind':'copyright-page', 'body':details},
    ]
