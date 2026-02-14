# Go Code Review Checklist ✅

This checklist is designed to help you write **idiomatic, production-ready Go**. Use it before submitting a PR or when reviewing your own code.

## 1. Style & Idioms (The "Go Way")

- [ ] **Formatted?** Did you run `gofmt` (or `goimports`)?
- [ ] **Naming:**
    -   Variables are short (`i`, `ctx`, `req`) where scope is small.
    -   Variables are descriptive (`cleanupInterval`, `validationTimeout`) where scope is large.
    -   Acronyms are consistent (`ServeHTTP`, not `ServeHttp`).
- [ ] **Error Handling:**
    -   Errors are handled, not ignored (`_`).
    -   Errors are wrapped with context when passed up: `fmt.Errorf("parsing config: %w", err)`.
    -   **Anti-Pattern**: `if err != nil { return err }` (without context in high-level logic).
- [ ] **Early Returns:**
    -   Keep the "happy path" aligned to the left.
    -   **Anti-Pattern**: Huge `else` blocks. Return early on failure.

## 2. Architecture & Design

- [ ] **Interfaces:**
    -   **Accept Interfaces, Return Structs**.
    -   Interfaces are defined where they are *used*, not where they are implemented.
    -   **Anti-Pattern**: Returning an interface from a function (pre-optimizing for abstraction).
- [ ] **Package Structure:**
    -   Avoid `common` or `utils` packages if possible (they become dumping grounds).
    -   No cyclic dependencies.
- [ ] **Config:**
    -   No hardcoded magic numbers or strings (use constants or config structs).

## 3. Concurrency (Danger Zone ⚠️)

- [ ] **Lifecycle:**
    -   If you start a goroutine, do you know how it stops?
    -   **Anti-Pattern**: "Fire and forget" goroutines without context cancellation or WaitGroups.
- [ ] **Data Races:**
    -   Did you run tests with `-race`?
    -   Are maps protected by a Mutex if accessed concurrently?
- [ ] **Channels:**
    -   Are channels closed by the sender?
    -   **Anti-Pattern**: Using channels for simple signaling where a Mutex would be clearer.

## 4. Performance

- [ ] **Allocations:**
    -   Pre-allocate slices/maps if size is known: `make([]string, 0, 100)`.
- [ ] **IO:**
    -   Are response bodies closed? `defer resp.Body.Close()`.
    -   Are you reading the whole file into memory? (Use `io.Copy` or scanners for large files).

---

> **💡 "The Legend" Tip:**
> When defending your code in an interview, don't just say "it works".
> Say: *"I chose to return a struct here to allow the caller to decide the abstraction, following the 'Accept Interfaces, Return Structs' idiom. It reduced necessary allocations by X."*
