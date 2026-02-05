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

## 13.2 Connection Pooling (`sql.Open`)
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
