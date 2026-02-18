# Changelog

All notable changes to this project will be documented in this file.
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
