# Logic: The Fork in the Road

Websites aren't just static posters. They change. They think.
To make your code "think", you use **Logic**.

## The Boolean (The Switch)
A **Boolean** (`bool`) is the simplest type. It can only be `true` (ON) or `false` (OFF).

```go
var isOpen bool = true
```

## The If/Else Statement (The Fork)
Imagine you are walking down a road. You see a sign: **"Is the Shop Open?"**.

- **YES (True)**: You turn **Left** and enter the shop.
- **NO (False)**: You turn **Right** and go home.

In Go, we write this "Fork in the Road" like this:

```go
if isOpen {
    // Left Path: Show the books
    fmt.Println("Welcome!")
} else {
    // Right Path: Go home
    fmt.Println("Closed.")
}
```

## Your Mission
1. Open `lessons/03-conditionals/main.go`.
2. Run it: `go run lessons/03-conditionals/main.go`.
3. Visit `http://localhost:8081` — You should see the books.
4. **Change the code**: Set `isOpen := false`.
5. Restart the server.
6. Refresh the page — You should see the "Closed" message.

You just made your website dynamic!
