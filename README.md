# Go: от первой строки до интернет-магазина

Бесплатный ознакомительный фрагмент книги для читателей [shanraq.org](https://shanraq.org/shop/go-book): **предисловие и главы 1–12 включительно**. От первой программы на Go до HTTP-каталога с HTML-шаблонами, `cmd`, `internal` и `web`.

- [Скачать ознакомительные PDF, EPUB, HTML и примеры](https://github.com/dake-edu/gopher-shop/releases).
- [Читать исходный текст](book/manuscript/00-preface.md).
- [Состав фрагмента и запуск примеров](book/README.md).
- [Полная книга на shanraq.org](https://shanraq.org/shop/go-book): 42 главы, продолжение проекта до заказов, тестовой оплаты, выдачи файлов, административной части и эксплуатации.

Редакция фрагмента: **0.20.0-sample**. Примеры проверяются на Go 1.27.1. Бесплатный фрагмент позволяет попробовать методику и код перед покупкой. Главы 13–42 и полные платные файлы в этот репозиторий и его выпуски не входят.

Ниже сохранено описание предыдущего курса. Его папки `lessons/`, `docs/` и корневое API — отдельные учебные материалы; контрольные точки новой книги находятся в `book/examples/`.

<div align="center">
  <img src="docs/public/gopher.png" alt="Gopher Shop Logo" width="200"/>
</div>

Welcome to **The Gopher Shop** — a practical earlier Go course with local teaching demos. It does not establish a professional qualification or provide a complete commercial store.

## 🚀 The Mission
The earlier course explores a teaching REST API for a digital book store from scratch. We don't just teach syntax; we teach **architecture, patterns, and logic**.

## 🧠 Learning Methodology
This course uses a unique **Visual Anchor System**. Instead of long walls of text, we provide consistent metaphors:

| Concept | Visual Anchor | Meaning |
| :--- | :--- | :--- |
| **Variable** | 📦 The Box | A labeled storage container. |
| **Slice** | 📚 The Shelf | A window into a growable list. |
| **Map** | 🎫 The Coat Check | Key-Value lookup (Ticket -> Coat). |
| **Validation** | 🛡️ The Quality Gate | Rejecting bad data immediately. |
| **Interface** | 🔌 The Universal Plug | Accepting any tool that fits the socket. |

## 🗺 Curriculum Roadmap

### Level 1: The Junior (Foundations)
*Syntax, Memory operations, Logic, Loops, and Basic HTTP.*
- **Landing Page**: [Junior Path](./docs/junior-path.md)
- **Goal**: Write a program that listens and responds.

### Level 2: The Apprentice (Building)
*Structs, JSON, Configuration, In-Memory Storage.*
- **Focus**: Building the "Components" of the shop.
- **Key Skill**: Separation of Concerns.

### Level 3: The Professional (Production)
*Interfaces, PostgreSQL, Middleware, Testing, CI/CD.*
- **Focus**: Making the shop reliable and production-ready.
- **Key Skill**: Dependency Injection and Automation.

## 🛠 How to Run

### 1. View the Documentation
Read the course material locally.
```bash
npm install
npm run docs:dev
# Visit http://localhost:5173
```

### 2. Run the Educational Demo
A standalone, single-file web app to visualize the final goal (UI + In-Memory Store).
```bash
go run ./cmd/web-demo
# Visit http://localhost:8082
```

### 3. Run the Main API (Capstone)
The legacy teaching API (requires PostgreSQL; the optional Compose file starts a local database). It has no complete order/payment flow or access control. Run it locally, not as a public store. For a first run, create `.env` from the example. If `.env` already exists, skip the copy command and compare its database settings instead.
```bash
cp .env.example .env
docker compose up -d --wait
docker compose exec -T db psql -U gopher -d gophership -v ON_ERROR_STOP=1 < scripts/init-legacy-db.sql
go run ./cmd/api
# Visit http://127.0.0.1:8080/health
```

Use Go 1.27.1, the tested toolchain, and run these commands from the repository root. They use the default local database user and database from `.env.example`; adjust the `psql` arguments if you change them. In PowerShell, use `Copy-Item .env.example .env`, and feed the schema with `Get-Content scripts/init-legacy-db.sql | docker compose exec -T db psql -U gopher -d gophership -v ON_ERROR_STOP=1`. If `.env` already exists, compare it with the example instead of overwriting it. The Compose database listens on host port 5442. The schema creates the initial table; it does not migrate an existing table.

The legacy API keeps its older positive-only `float64` price contract. This is not the money model of the new book, which uses integer minor units and permits explicitly free books.

---
*Built with ❤️ for the Go Community.*

## 📜 License & Legal

This project is protected by the **GNU Affero General Public License v3.0 (AGPL-3.0)**.

### 🎓 For Students & Individuals
**You are free to use this project** to learn, fork, modify, and build your own learning portfolio. We encourage open-source contributions!

### 🏢 For Commercial Educational Platforms
If you intend to use this methodology, code, or materials in a **commercial course, bootcamp, or proprietary platform**:
1.  **Open Source**: You must open-source your entire platform code under AGPL-3.0.
2.  **Commercial License**: If you cannot open-source your platform, you **MUST purchase a commercial license**.

> **"If you profit from our work, you must either contribute back (Open Source) or pay it forward (Commercial License)."**

Contact `baimurza.daulet@gmail.com` for commercial inquiries.


В редакции 0.20 расширены главы 3–7: восемь небольших запускаемых опытов перед кодом магазина объясняют значения, ветвления, циклы, функции, текст и контейнеры. Их ожидаемый вывод проверяется автоматически. Это первый этап педагогической переработки; наличие всех глав не означает завершённую проверку самостоятельного обучения.
