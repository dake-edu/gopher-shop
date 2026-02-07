# Chapter 08: Maps

> **"Give me the ticket, I'll give you the coat."**

Arrays and Slices are great if you want to find things by number (Index 0, Index 1).
But what if you want to find a User by their **Name**? You don't want to loop through a list of 1,000,000 users.
You need the **Map** (Hash Table).

## 8.1 The Coat Check Metaphor
Imagine a coat check at a theater.
1.  You give your coat.
2.  You get a specific ticket #123 (The **Key**).
3.  Later, you give back ticket #123.
4.  The attendant instantly finds your specific coat (The **Value**).

They don't search all coats. They go straight to hook #123. That is $O(1)$ complexity (Instant speed).

## 8.2 Syntax
```go
// Key: String (Name) -> Value: Integer (Year)
phones := make(map[string]int)

// Add
phones["Alice"] = 12345
phones["Bob"] = 67890

// Get
fmt.Println(phones["Alice"]) // Prints 12345
```

## 8.3 The Zero Value Trap
If you ask for a key that doesn't exist, Go **does not crash**. It returns the "Zero Value" (0 for int, "" for string).
```go
fmt.Println(phones["Ghost"]) // Prints 0. Wait, does a ghost have phone number 0?
```

### The "Comma Ok" Idiom
To check if a key *actually* exists, we use the second return value.
```go
price, ok := phones["Ghost"]
if !ok {
    fmt.Println("Ghost not found!")
} else {
    fmt.Println("Ghost's number is", price)
}
```

**Visual Anchor**:
> **The Library Card Catalog**:
> - You don't walk through every shelf looking for "Harry Potter" (Looping).
> - You go to the Catalog (Map), look up "Harry Potter" (Key), and it tells you exactly "Aisle 4, Shelf B" (Value).
> - You walk straight there.

::: details 🎓 Knowledge Check
: What two values does a map lookup return?
**Answer**:
1.  **Value**: The data you asked for (e.g., the price).
2.  **Exists**: A boolean (`true`/`false`) telling you if the key was actually found. Always check this!
:::
