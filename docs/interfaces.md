# 12. The Architect (Interfaces)

## 12.1 The Implicit Contract (Duck Typing)
In strict languages (Java, C++), you must **explicitly** sign the contract.
`class MyStore implements BookRepository`.

In Go, the contract is **implicit**. This is called "Duck Typing".
> *"If it walks like a duck and quacks like a duck, it is a duck."*

If your struct has the methods `GetAll()` and `GetByID()`, Go automatically considers it a `BookRepository`.

### Anatomy of an Interface
```go
type BookRepository interface {
    GetAll() ([]Book, error)
    GetByID(id string) (Book, error)
}
```
1.  **`type`**: New Type Definition.
2.  **`BookRepository`**: The name (Contract Name).
3.  **`interface`**: The Kind. It contains only Method Signatures, no code.
4.  **`GetByID`**: Method Name.
5.  **`(id string)`**: Usage requirements (Input).
6.  **`(Book, error)`**: Expected result (Output).

## 12.2 Dependency Injection (DI)
Big phrase, simple concept.
**Don't build your tools inside your house. Buy them and bring them in.**

**Bad (Tight Coupling)**:
```go
func NewHandler() *Handler {
    repo := PostgresStore{} // Hardcoded! Can't switch to Memory.
    return &Handler{repo: repo}
}
```

**Good (Dependency Injection)**:
```go
func NewHandler(repo BookRepository) *Handler {
    return &Handler{repo: repo} // Flexible! Accepts any repo.
}
```
Now `main.go` decides which tool to use.

::: details 🎓 Knowledge Check: Why do we use Interfaces for Dependency Injection?
**Answer**: To decouple our code. If `Handler` depends on an `Interface`, we can swap the real database for a "Fake Database" (Mock) during testing, or switch from Postgres to MySQL without rewriting the Handler.
:::

