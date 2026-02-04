# Data Validation & Guard Clauses

## The "Bouncer" Pattern

> [!TIP]
> ⚓ **Visual Anchor:** The Club Bouncer

Validation is like a bouncer at a club.
- **Bad ID?** (Invalid JSON) --> REJECTED immediately.
- **Underage?** (Price < 0) --> REJECTED immediately.
- **No Name?** (Empty Title) --> REJECTED immediately.

Only after passing all checks can you enter the club (Database).

## Guard Clauses

In Go, we prefer "Guard Clauses" over deep nesting.

**Bad (Deep Nesting):**
```go
if title != "" {
    if price > 0 {
        // Do Logic
    }
}
```

**Good (Guard Clauses):**
```go
if title == "" {
    return Error("title required")
}
if price <= 0 {
    return Error("invalid price")
}
// Do Logic
```

This keeps the "Happy Path" on the left edge of the screen.
