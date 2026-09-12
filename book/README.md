# Go: от первой строки до книжного магазина

Практическая книга shanraq.org, связанная с репозиторием [dake-edu/gopher-shop](https://github.com/dake-edu/gopher-shop).

Статус: **рабочая редакция 0.3.1**, вступление и главы 1–8. Полная программа содержит 42 главы на уровнях Junior, Middle и Senior. Книга и итоговый магазин ещё не завершены.

## Читать рабочий фрагмент

После сборки:

- [PDF](build/preview/go-book-preview.pdf)
- [EPUB](build/preview/go-book-preview.epub)
- [HTML в браузере](build/preview/html/index.html) или [ZIP для скачивания](build/preview/go-book-preview-html.zip)

Сборки локальные и исключены из Git. В клоне GitHub выполните команды ниже. HTML-архив работает после распаковки без сервера и сети; поиск доступен при включённом JavaScript.

## Источники и код

- `book.json` — реестр глав, модулей, версии Go и форматов.
- `manuscript/` — рукопись с включениями из исходников.
- `examples/` — самостоятельная контрольная точка каждой исполняемой главы.
- `build/reading/` — текст глав с подставленным проверяемым кодом.
- [Программа](editorial/blueprint.md), [стиль](editorial/style.md), [аудит](editorial/audit.md), [покрытие тем](editorial/coverage.md).
- [Полная карта языка](editorial/language-map.md) и [правила объяснения и опорных сигналов](editorial/teaching-contract.md).
- [Продолжение работы](CONTINUE.md).

Из корня репозитория, с Go 1.27.1 и Python 3.12 или новее:

```sh
python3 book/tools/check.py
```

В Windows используйте `python` вместо `python3`, если так называется установленная команда. Скрипт запускает вложенные модули, сверяет вывод, выполняет vet и тесты, проверяет учебные изменения и создаёт читательские Markdown-файлы. Примеры не требуют внешних Go-зависимостей. Дополнительная проверка лексических токенов, включая тесты, обнаруживает новые знаки до назначенной главы объяснения; качество объяснения проверяется редактором.

Для сборки трёх форматов создайте отдельное Python-окружение. macOS/Linux:

```sh
python3 -m venv book/.venv
book/.venv/bin/python -m pip install -r book/requirements-build.txt
book/.venv/bin/python book/tools/build_preview.py
book/.venv/bin/python book/tools/validate_preview.py
```

PowerShell:

```powershell
python -m venv book/.venv
book/.venv/Scripts/python.exe -m pip install -r book/requirements-build.txt
book/.venv/Scripts/python.exe book/tools/build_preview.py
book/.venv/Scripts/python.exe book/tools/validate_preview.py
```

Локальная проверка выполнена на macOS arm64. [CI для коммита fdf3656](https://github.com/dake-edu/gopher-shop/actions/runs/34596143433) прошёл на Windows, Linux и macOS; на Linux также собраны и проверены все три формата. EPUB прошёл структурную проверку и EPUBCheck 5.3.0 без ошибок и предупреждений (2026-09-11). Ручная проверка в читалках ещё предстоит. PDF пока без тегированной структуры доступности; это рабочая вёрстка, не финальный доступный выпуск.

Шрифты Noto включены с [лицензией SIL OFL 1.1](assets/fonts/OFL.txt). Корневая лицензия репозитория сохранена без изменений. Финальные метаданные издания оформляются перед выпуском.

В редакции 0.3 страницы белые, листинги и код внутри предложений имеют общую статическую подсветку. В главе 2 есть цветовая легенда, в главах 1–8 — бытовые опорные образы. Во введении обсуждается самостоятельное создание продуктов с ИИ без обещания гарантированной карьеры. Правила: [цвет, образы и голос книги](editorial/color-and-images.md).

В редакции 0.3.1 отдельно окрашены `package`, `import`, `func`, остальные ключевые слова и скобки; легенда находится в главе 2.
