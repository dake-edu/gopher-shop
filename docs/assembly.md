# 17. The Grand Assembly

You have bricks, wood, and glass. Now let's build the house.
In this final chapter, we look at `cmd/api/main.go`. This is where all the isolated pieces we learned about (Router, Config, Database) are wired together.

## 17.1 Standard Project Layout
Why isn't everything in the root folder?
Professional Go projects often use the **Standard Go Project Layout**:

1.  **`cmd/`**: The Main Applications.
    - `cmd/api/main.go`: The entry point for our API.
    - `cmd/cli/main.go`: (Optional) CLI tools.
    - *Rule*: No logic here. Just wiring.
2.  **`internal/`**: The Library Code.
    - `internal/models`: Data Structures.
    - `internal/store`: Database Logic.
    - *Rule*: Code here allows our app to work, but other people can't import it (Go enforces this privacy).

## 17.2 The Main Wiring
Open `cmd/api/main.go`. Let's read it like a schematic.

### Step 1: Power On (Configuration)
```go
cfg := config.Load()
```
- **Why?**: Before doing anything, we need to know *how* to run (Port, DB Password).
- **Chapter**: 9 (Configuration).

### Step 2: Unlock the Warehouse (Database)
```go
db, err := sql.Open("postgres", cfg.DB.DSN())
```
- **Why?**: We establish the connection pool. We don't query yet; we just prepare the line.
- **Chapter**: 13 (Postgres).

### Step 3: Hire the Staff (Dependency Injection)
```go
var bookStore store.BookRepository = store.NewPostgresBookStore(db)
```
- **Why?**: We create the `bookStore` worker and **give** it the database connection (`db`).
- **Concept**: Dependency Injection (Chapter 12).

### Step 4: Security Checkpoints (Middleware)
```go
server := &http.Server{
    Handler: middleware.Recoverer(middleware.Logger(mux)),
}
```
- **Why?**: We wrap the router (`mux`) in our "Onion Layers".
    1.  Request hits `Recoverer` (Safety Net).
    2.  Hits `Logger` (Record keeping).
    3.  Hits `mux` (The Router).
- **Chapter**: 14 (Middleware).

### Step 5: Open for Business
```go
server.ListenAndServe()
```
- **Why?**: This starts the infinite loop that listens for traffic on the port.
- **Chapter**: 7 (Web Server).

### 4. Graceful Shutdown (Dying with Dignity)
In production, servers restart often. You don't want to kill active users mid-request.
We catch OS signals (`SIGINT`, `SIGTERM`) and give the server a standardized timeout (e.g., 5 seconds) to finish current jobs before quitting.

```go
// Wait for Ctrl+C
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

// Give 5 seconds to finish
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

## 17.3 You Did It!
You have built a modular, professional-grade REST API.

## The Checklist
- [ ] **Data**: `struct` with JSON tags? (Ch 8)
- [ ] **Logic**: Handlers inject a Repository? (Ch 12)
- [ ] **Safety**: Unit Tests passed? (Ch 15)
- [ ] **Deployment**: CI/CD pipeline green? (Ch 16)

If you can say **YES** to all these, you are no longer a Junior.
**Welcome to the Professional Tier.**

::: details 🎓 Knowledge Check: Why shouldn't we put business logic in `cmd/api/main.go`?
**Answer**: Separation of Concerns. `cmd/` is only for **wiring** (starting the engine). Logic belongs in `internal/` so it can be tested in isolation and reused.
:::

