# 11. The Inspector (Validation)

Garbage In, Garbage Out. To keep our shop clean, we need strict quality control.

## 11.1 The Club Bouncer
Imagine your database as an exclusive VIP Club.
The **Validator** is the Bouncer at the door.

- **Request**: "I want to add a book with Price: -50."
- **Bouncer**: "Rejected. Price cannot be negative."

## 11.2 Guard Clauses
In Go, we prefer **Guard Clauses**. We check for errors *first* and return immediately. This keeps the "Happy Path" (success logic) aligned to the left, easy to read.

**Bad (Deep Nesting):**
```go
if title != "" {
    if price > 0 {
       // Save to DB
    }
}
```

**Good (Guard Clause):**
```go
if title == "" {
    return Error("Title is required") // Stop!
}
if price <= 0 {
    return Error("Invalid Price") // Stop!
}

// Save to DB (Happy Path)
```

**Visual Anchor**:
The Bouncer stops you at the door. He doesn't let you walk all the way to the bar before checking your ID.
