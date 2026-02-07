# Chapter 06: Loops

Managing one variable is easy. Managing a thousand is hard. We need containers.

## 6.1 The Egg Carton (Arrays)
An **Array** is a fixed-size container.
- **Rule**: When you buy an egg carton for 12 eggs, it always has 12 slots. You cannot magically stretch it to hold 13.
- **Syntax**: `[3]int` (An array of exactly 3 integers).

```go
var carton [3]int
carton[0] = 10
carton[1] = 20
// carton[3] = 30 // ERROR! Index out of bounds.
```
*Arrays are rarely used directly because they are too rigid.*

## 6.2 The Magic Shelf (Slices)
A **Slice** is a window into an underlying array. Ideally, it feels like a dynamic list that can grow.
- **Syntax**: `[]int` (No number).

```go
shelf := []string{"Book A", "Book B"}
shelf = append(shelf, "Book C") // Magic! It grows.
```

## 6.3 The Conveyor Belt (For Loops)
How do we process 1,000 items? We don't write 1,000 lines of code.
We use a **Loop**.

### The Standard Loop
"Init; Condition; Post"
```go
// Start at 0; Keep going while i < 5; Add 1 to i after each step
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### The Range Loop (Best for Libraries)
Ideal for Slices. It gives you the **Index** (Position) and the **Value** (Copy of the item).

```go
library := []string{"Go", "Rust", "Python"}

for index, title := range library {
    fmt.Printf("%d: %s\n", index, title)
}
```


## 6.4 Maps (The Keychains)
Arrays are ordered lists (1st, 2nd, 3rd).
**Maps** are labelled storage (Key -> Value).

```go
// [Key] -> Value
menu := make(map[string]float64)
menu["Coffee"] = 2.50
```

*Metaphor*: A Coat Check. You give a ticket (Key), you get the coat (Value).
You don't say "Give me the 5th coat". You say "Give me the coat for ticket 'Coffee'".

**Visual Anchor**:
- **Array**: Solid concrete block with slots.
- **Slice**: A flexible viewing window.
- **Loop**: A factory conveyor belt processing items one by one.

::: details 🎓 Knowledge Check
1.  **How do you make an infinite loop in Go?**
    `for { }`.
2.  **What does `range` do?**
    It iterates over elements of an array, slice, string, map, or channel.
3.  **How do you stop a loop?**
    Use the `break` keyword.
:::
