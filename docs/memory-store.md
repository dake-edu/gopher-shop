# 10. The Memory (In-Memory Storage)

For our first version of the shop, we need a place to store books.
We will use the computer's **RAM** (Random Access Memory).

## 10.1 The Whiteboard
Think of RAM as a **Whiteboard**.
- **Pros**: It is incredibly fast. Writing takes nanoseconds.
- **Cons**: It is volatile. If you restart the server (erase the board), all data is lost.
- **Use Case**: Perfect for learning, caching, or temporary data.

## 10.2 The Map (The Lookup Table)
In Go, the best tool for this is a **Map**.
It connects a **Key** (ID) to a **Value** (Book).

```go
// "A map where Keys are Strings and Values are Books"
var storage = make(map[string]Book)
```

## 10.3 Pointers (The Address Card)
When we retrieve a book, do we want a **Photocopy** or the **Original**?
- **Value (`Book`)**: A photocopy. If you write on it, the original doesn't change.
- **Pointer (`*Book`)**: A card with the shelf address. If you go to that shelf and change the book, it changes for everyone.

Professional apps often use **Pointers** to avoid copying large objects and to share data.
