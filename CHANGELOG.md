# Changelog

All notable changes to this project will be documented in this file.

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
