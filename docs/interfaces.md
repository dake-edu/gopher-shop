# 12. The Architect (Interfaces)

As your shop grows, you need flexibility. You might want to switch from Memory storage to PostgreSQL, or File storage.
You don't want to rewrite your entire app every time.

## 12.1 The Universal Plug
An **Interface** is like a standard wall socket.
- The Wall Socket doesn't care if you plug in a Lamp, a TV, or a Vacuum.
- It just guarantees: "I provide electricity."

## 12.2 Defining the Contract
In Go, an interface defines **Behavior** (what it *does*), not Data.

```go
type BookRepository interface {
    GetAll() ([]Book, error)
    GetByID(id string) (Book, error)
}
```

Any tool that has these two methods "satisfies" the interface and can be plugged in.

## 12.3 Dependency Injection
Instead of the app creating the storage, we **give** the storage to the app.

```go
// "Hole in the wall"
type Handler struct {
    repo BookRepository
}
```
Now we can plug in *anything* that fits the `BookRepository` shape.
