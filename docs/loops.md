# The Conveyor Belt (Loops)

Imagine you work in a warehouse. You have a huge shelf full of boxes (a **Slice**).
Your boss tells you: *"Put a price tag on every single box."*

You don't run to the shelf and jump on all the boxes at once.
You use a **Conveyor Belt**.

1. The belt brings **one box** to you.
2. You stick the tag on it.
3. The belt moves, creating the **next box**.
4. You repeat this until the shelf is empty.

## Loops in Go (`for range`)

In Go, the `for` loop is your conveyor belt.

```go
// The Shelf
library := []Book{book1, book2, book3}

// The Loop (Conveyor Belt)
for _, book := range library {
    // Processing ONE item at a time
    fmt.Println(book.Title)
}
```

- `library`: The big collection (The Shelf).
- `book`: The single item currently in front of you (The Box).
- `range`: The motor that moves the belt.

## Your Mission
1. Open `lessons/04-loops/main.go`.
2. Notice how we use the loop to build an HTML Table row (`<tr>`) for every book.
3. Run it: `go run lessons/04-loops/main.go`.
4. Visit `http://localhost:8081`.

You successfully converted data into a structured table!
