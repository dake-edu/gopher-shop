# 13. The Warehouse (PostgreSQL)

Memory is fast, but it forgets. For a real business, we need a **Steel Vault**.

## 13.1 SQL (The Language of Data)
Structured Query Language (SQL) is the standard way to talk to databases.
- `CREATE TABLE`: Build a shelf.
- `INSERT`: Put a box on the shelf.
- `SELECT`: Find a box.

## 13.2 The Driver (`lib/pq`)
Go doesn't speak SQL natively. It needs a translator, called a **Driver**.
We use `lib/pq` to talk to Postgres.

## 13.3 Docker (The Container)
Installing Postgres on every developer's laptop is a pain.
We use **Docker** to run Postgres in a lightweight container.

```yaml
services:
  db:
    image: postgres:15
    environment:
      POSTGRES_PASSWORD: secret
```

**Visual Anchor**:
- **Memory**: A Whiteboard (Fast, Erased easily).
- **Postgres**: A Steel Safe (Secure, Persists forever).
