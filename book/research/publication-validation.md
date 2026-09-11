# Проверка редакции 0.2 — 2026-09-11

Платформа: macOS arm64. `check.py`: Go 1.27.1, 42 команды, 9 текстов (вступление и главы 1–8). `validate_preview.py`: 10 HTML-страниц, 65 локальных ссылок, структура EPUB и контрольные суммы архивов проверены.

PDF: 44 страницы A4, Unicode-шрифты Noto, векторные опорные схемы. Извлечение текста через pdftotext выполнено. PDF не тегирован.

Официальный [EPUBCheck 5.3.0](https://github.com/w3c/epubcheck/releases/tag/v5.3.0), команда `java -jar epubcheck.jar go-book-preview.epub`:

```text
Validating using EPUB version 3.3 rules.
No errors or warnings detected.
Messages: 0 fatals / 0 errors / 0 warnings / 0 infos
```

EPUBCheck не подтверждает удобство чтения и полную доступность. Ручная проверка в читалках, браузерах с file:// и проверка доступности остаются задачами редакции. Удалённый CI ещё не подтверждён.

Корневое приложение вместе с новым инструментом `book/tools/syntaxaudit.go` прошло `go test ./...` на Go 1.27.1. Проверен вид страницы 42: устранён отсутствующий в PDF-шрифте знак стрелки в подписи схемы.
