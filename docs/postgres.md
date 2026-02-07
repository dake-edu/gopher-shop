# 13. The Warehouse (PostgreSQL)

## 13.1 The Low-Level Driver
To talk to a SQL database, Go uses the standard package `database/sql`.
However, `database/sql` is just a **Manager**. It needs a worker (Driver) to speak the specific language of PostgreSQL.

We import the driver **Blankly**:
```go
import _ "github.com/lib/pq"
```
### Anatomy of the Blank Import (`_`)
1.  **`import`**: Bring in code.
2.  **`_` (Underscore)**: "I am not going to call this package directly in my code."
    - *Why?* We only want the package's `init()` function to run.
    - Inside `lib/pq`, the `init()` function registers itself with Go's SQL Manager: "Hey, I know how to talk to Postgres!"

### The Enemy: SQL Injection
If you concatenate strings (`"SELECT * FROM users WHERE name = " + input`), you create a **Hole in the Wall**. A hacker can pass `'; DROP TABLE users; --` and destroy your database.

**The Solution: Armored Windows ($1)**
When you use placeholders (`$1`, `$2`), you send the data in a separate, sealed envelope. The database treats it strictly as *text*, never as *commands*.

```go
// BAD (Hole in the Wall)
db.Query("SELECT * FROM users WHERE id = " + id)

// GOOD (Armored Window)
db.Query("SELECT * FROM users WHERE id = $1", id)
```

### 3. Connect & Ping
on Pooling (`sql.Open`)
```go
db, err := sql.Open("postgres", "user=dake dbname=shop...")
```
**Crucial Concept**: `sql.Open` does NOT establish a connection.
It prepares a **Connection Pool**.
- It opens connections only when needed.
- It keeps them open for reuse (Performance).
- It handles network drops automatically.

Comparison:
- **Python/Django**: Often opens a new connection per request (Slow).
- **Go**: Maintains a pool of long-lived connections (Fast).

::: details 🎓 Knowledge Check: Does `sql.Open` connect to the database immediately?
**Answer**: **No!** It only initializes the **Connection Pool** and config. The actual connection happens lazily when you first try to query the DB (e.g., `db.Ping()`).
:::

