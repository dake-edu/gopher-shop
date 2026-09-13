# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased] — Book development

- Added chapter 9 on error handling, bounded file input and resource cleanup, with standard-library-only examples and failure-path tests.
- Enforced standard-library-only checkpoint dependencies and documented the native Go template/UI approach.

- Added the author-provided cover, title page and publication details to all three formats in preview 0.4; recorded the confirmed author name in PDF/EPUB metadata.

- Expanded preview 0.3.1 syntax colours: distinct package/import/func keywords and bracket families, with an updated reader legend.

- Updated the book preview to 0.3: white pages, shared syntax colours for listings and inline code, and a reader-facing colour legend.
- Added everyday mental images to chapters 1–8 and an introduction about building products with AI.
- Added source-fidelity and colour-contrast tests for publication tooling.

- Added a Russian book manuscript (preface and chapters 1–8), aligned with this repository.
- Added independent Go 1.27.1 checkpoints, expected-output checks, exercises and a negative mutation check.
- Added a chapter registry and PDF/EPUB/offline HTML preview builds.
- Added eight reproducible support diagrams, symbol-by-symbol explanations and a language coverage ledger.
- Added lexical first-use checks and UTF-8-safe publication tooling; EPUBCheck 5.3.0 reports no errors or warnings.
- Added a dedicated CI matrix for nested book modules.
- Retained earlier course material and API while the source audit continues.

## [v1.3.1] - 2026-02-18
### Fixed
- **Documentation**: Fixed broken links in `junior-path.md` pointing to old `lessons/` directory.
- **Build**: Resolved `npm run docs:build` failures caused by dead links.
## [v1.3.0] - 2026-02-18
### Added
- **Level 3 (Senior/Architect)**: Complete new section covering Microservices, Advanced Concurrency, Caching (Redis), Kafka, K8s, and Career.
- **Visual Signals**: Implemented "Shatalov Method" visual metaphors across ALL chapters (Junior, Middle, Senior).
- **GoTracker**: Specification for a delivery microservice (Capstone Project).
- **License**: Switched to **AGPL-3.0** to protect educational IP.

### Changed
- **Restructuring**: Split documentation into 3 distinct levels (Junior, Middle, Senior).
- **Mermaid Syntax**: Fixed rendering issues by enforcing strict quoting on all diagrams.
- **Navigation**: Updated sidebar to reflect the new 3-tier structure.


## [v1.2.0] - 2026-02-16
### Added
- **Architecture**: Introduced "3-Layer Cake" (Handler -> Service -> Store) architecture.
- **Observability**: Replaced standard `log` with structured logging (`slog`) and added `RequestLogger` middleware.
- **Concurrency**: Implemented Async Worker pattern (`internal/worker`) for non-blocking checkout.
- **Documentation**: Updated chapters 15 (Structure), 13 (Debugging), and 22 (Concurrency) with new diagrams.

### Changed
- Refactored `cmd/web-demo` to use Dependency Injection.
- `POST /checkout` is now asynchronous and returns immediately.

## [v1.1.0] - 2026-02-16
### Added
- **CI/CD**: GitHub Actions workflow for automated testing.
- **UI**: Added book covers, sticky sidebar, and update notification banner.
- **Pedagogy**: Added "Interview Defense" sections to lessons.

### Fixed
- Database connection port mismatch (`5432` -> `5442`).
- Dynamic footer year in documentation.
