# Chapter 09: Functions

> **"Don't repeat yourself (DRY)."**

If you copy-paste the same code 3 times, you are creating a maintenance nightmare.
We pack logic into reusable machines called **Functions**.

## 9.1 The Metaphor: The Meat Grinder
A function is like a meat grinder:
1.  **Top (Arguments)**: You throw in raw ingredients (Numbers, Strings).
2.  **Gears (Body)**: The machine processes them (Calculates, Formats).
3.  **Bottom (Return)**: Processed meat comes out (Result).

## 9.2 Anatomy of a Function
```go
// Keyword   Name      Inputs (Type)      Outputs (Type)
func         Add       (a int, b int)     int {
    return a + b
}
```

## 9.3 Multiple Returns
Unique to Go: A function can return **multiple things**.
Usually, we return `(Result, Error)`.

```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

**The "Guard Clause" Pattern**:
Notice how we check for errors *first* and return early? This keeps our code clean (no nested `else` hell).

## 9.3b Idiomatic Naming: Getters
In Java, you might write `GetName()`. In Go, we just say `Name()`.
-   **Bad**: `GetOwner()`, `GetPrice()`, `GetAll()`
-   **Good**: `Owner()`, `Price()`, `All()`

If a function *fetches* something remotely or expensively, we usually just verb it directly (`Fetch()`, `Find()`). If it's a simple property access, drop the `Get`.

## 9.4 Practice: Interaction & Validation
In this lesson, we build a function that accepts User Input, validates it (checks for errors), and returns a safe result.

<<< @/../lessons/04-interaction/main.go

::: details 🎓 Knowledge Check: What happens if we don't validate user input?
**Answer**: The "Quality Gate" stays open to garbage! Users could submit empty books, negative prices, or malicious code. Validation protects your database.
:::
