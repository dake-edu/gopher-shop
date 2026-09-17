"""Publication data rendered once for all reading formats."""
import html
from pathlib import Path

def front_pages(meta):
    p = meta['publication']
    esc = html.escape
    last = max(int(c['id']) for c in meta['chapters'])
    title = esc(meta['title'])
    author = esc(p['author'])
    date = esc(p['date_label'])
    publisher = esc(p['publisher'])
    edition = esc(meta['edition'])
    title_body = (f'<p class="author">{author}</p><div class="title-block"><h1>{title}</h1>'
                  f'<p class="subtitle">{esc(p["subtitle"])}</p>'
                  f'<p class="levels">{esc(p["levels"])}</p>'
                  f'<p class="edition-note">Ознакомительный фрагмент {edition}</p></div>'
                  f'<p class="imprint">{publisher}<br />{esc(p["date"][:4])}</p>')
    details = (f'<h1>Выходные сведения</h1>'
               f'<p><strong>{author}</strong><br />{title}: практическая книга. '
               f'— {publisher}, {p["date"][:4]}. — Электронное издание. '
               f'Ознакомительный фрагмент {edition}.</p>'
               f'<p>{esc(p["description"])}</p>'
               f'<p>{esc(p["audience"])}</p>'
               f'<p><strong>Автор:</strong> {author}.<br />'
               f'{esc(p["author_bio"])}<br />'
               f'<strong>Издательский проект:</strong> {publisher}.<br />'
               f'<strong>Дата этой редакции:</strong> {date}.<br />'
               '<strong>Язык:</strong> русский.<br />'
               '<strong>Форматы:</strong> PDF, EPUB, HTML.</p>'
               f'<p>Этот файл содержит предисловие и главы 1–{last}. Все 42 главы написаны; полный черновик проходит редакционную проверку. '
               'Дата относится к рабочей редакции, а не к завершённому коммерческому выпуску. '
               'ISBN и библиотечные классификационные индексы в этой редакции не указаны.</p>'
               f'<p>© {p["date"][:4]} {author}.<br />'
               'Исходная обложка предоставлена автором; новая версия адаптирована с помощью ИИ. Подготовка рукописи и учебных материалов '
               'ведётся с использованием инструментов ИИ.</p>'
               '<p>Персонаж Go gopher создан Renee French: '
               '<a href="https://go.dev/blog/gopher">источник</a>, '
               '<a href="https://creativecommons.org/licenses/by/4.0/">CC BY 4.0</a>. '
               'В обложке, логотипе и иконках использована адаптация с помощью ИИ: '
               'добавлены тележка и книга, изменены композиция и оформление. '
               'Издание не является официальным продуктом проекта Go.</p>'
               '<p>Код, перенесённый из исходного репозитория, сохраняет GNU AGPL-3.0; '
               '<a href="https://github.com/dake-edu/gopher-shop/blob/main/LICENSE">текст лицензии</a>. '
               'Шрифты Noto распространяются по '
               '<a href="https://openfontlicense.org/">SIL Open Font License 1.1</a>; '
               'копия лицензии включена в электронный комплект. Этот фрагмент опубликован бесплатно для ознакомления. Главы 13–42 и полные PDF, EPUB, HTML предназначены для продажи через shanraq.org. Бесплатное чтение фрагмента не изменяет условия лицензирования материалов.</p>'
               f'<p><strong>Сайт:</strong> <a href="https://shanraq.org">shanraq.org</a>.<br />'
               '<strong>Код книги:</strong> <a href="https://github.com/dake-edu/gopher-shop/releases">бесплатный архив примеров глав 2–12</a>. Покупка для запуска этих примеров не требуется.</p>')
    return [
        {'stem':'index', 'title':'Обложка', 'kind':'cover',
         'body':f'<img class="book-cover" src="cover/{esc(Path(p['cover']).name, quote=True)}" alt="Обложка книги «{title}»: голубой суслик с тележкой и книгой с той же обложкой" />'},
        {'stem':'title-page', 'title':'Титульный лист', 'kind':'titlepage', 'body':title_body},
        {'stem':'publication-details', 'title':'Выходные сведения', 'kind':'copyright-page', 'body':details},
    ]
