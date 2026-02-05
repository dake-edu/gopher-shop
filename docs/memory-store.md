# 10. The Memory (In-Memory Storage)

## 10.1 The Map (Hash Table)
We need a way to find a Book by its ID instantly.
In Computer Science, this is a **Hash Map** (or Dictionary).

```go
var storage = make(map[string]Book)
```

### Anatomy of the Declaration
1.  **`var`**: Create a variable.
2.  **`make`**: **Crucial Keyword**. Maps must be initialized before use.
    - If you just did `var m map[string]int`, it is `nil`. Writing to it causes a **Panic** (Crash).
    - `make` allocates the memory bucket for the map.
3.  **`map`**: The type.
4.  **`[string]`**: The Key Type (The ID is a string).
5.  **`Book`**: The Value Type (The full Book struct).

## 10.2 Pointers vs Values (`*Book` vs `Book`)
This is the hardest concept for beginners coming from Python/JS.

```go
func GetBook() *Book  // Returns a Pointer (Address)
func GetBook() Book   // Returns a Value (Photocopy)
```

### The Metaphor
- **Value (`Book`)**: You ask for the Mona Lisa. I take a **Photocopy** and hand it to you. If you draw a mustache on the copy, the real Mona Lisa is unchanged.
- **Pointer (`*Book`)**: You ask for the Mona Lisa. I write **"Room 303, Wall 2"** on a card and hand it to you. If you go there and draw a mustache, the real painting is changed forever.

### Why use Pointers?
1.  **Performance**: Copying a huge book is slow. Passing a small address card is fast.
2.  **Mutability**: We *want* to modify the original (e.g., updating the price).

### Comparison: Memory
| Language | Default Behavior |
| :--- | :--- |
| **Python** | Everything is a Reference (Pointer-like). |
| **Java** | Objects are References. Primitives (`int`) are Values. |
| **Go** | Everything is a **Value** (Copy) by default. You must explicitly ask for a Pointer (`*`). |
