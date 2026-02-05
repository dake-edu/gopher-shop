# Junior Path: The Foundations

Welcome to **Step 0**. Before building complex web servers, we need to understand the building blocks of Go.

## The Visual Anchor System

We use simple metaphors to make abstract concepts concrete.

### 1. Checks (Variables)
Think of a variable as a labeled box. You put a value inside, and you can take it out later.
In a strongly typed language like Go, the box has a specific shape. You can't put a `Square` (String) into a `Round` hole (Integer).

### 2. The Container (Structs)
A **Struct** is a composite data type. Imagine a shipping container or a labeled cardboard box.
It groups related variables together.

```go
// This is the blueprint for our container
type Book struct {
    Title  string
    Author string
    Price  float64
}
```

### 3. The Shelf (Slices)
A **Slice** is a dynamic list. Imagine a long wooden shelf.
You can place strictly one type of item on it (e.g., only Books).
You can add more items (`append`), and the shelf grows magically to accommodate them.

## Foundations: Variables, Constants, and Types

To build the "Container" and "Shelf", we need basic materials.

### The "Box" (Variable)
A **Variable** is a box that you can open, empty, and refill.
- **Go Syntax**: `var count int = 10`
- **Metaphor**: You have a box labeled "Count". You put the number 10 in it. Later, you can take 10 out and put 11 in.

### The "Stone Carving" (Constant)
A **Constant** is a value carved into stone. It can never change once the program runs.
- **Go Syntax**: `const Pi = 3.14`
- **Metaphor**: A stone tablet. You can read it a million times, but you cannot erase the writing.

### Choosing Materials (Types)
- **Strings**: For text (Title, Author). Like paper notes.
- **Float64**: For money (Price). Allows decimals ($35.99).
- **Int**: For whole numbers (ID, Count). You can't have 1.5 Books.

## Your First Code
We created a simple program in `lessons/01-basics/main.go`.

Run it to see your shelf in action:
```bash
go run lessons/01-basics/main.go
```
