# Chapter 05: Logic

Code that runs in a straight line is boring. Smart programs make decisions.

## 5.1 The Judge (Comparison Operators)
Before we make a decision, we need to judge the situation. The result is always a **Boolean** (`true` / `false`).

| Operator | Meaning | Example | Result |
| :--- | :--- | :--- | :--- |
| `==` | Equal to | `5 == 5` | `true` |
| `!=` | Not Equal | `5 != 3` | `true` |
| `<` | Less than | `10 < 5` | `false` |
| `>=` | Greater or Equal | `10 >= 10` | `true` |

## 5.2 The Circuit (Boolean Algebra)
Sometimes you need to check multiple things.

### AND (`&&`)
**Both** must be true.
> "I will buy the car if it is **Red** AND **Fast**."
- `true && true` $\rightarrow$ `true`
- `true && false` $\rightarrow$ `false`

### OR (`||`)
**at least One** must be true.
> "I will eat if I am **Hungry** OR **Bored**."
- `true || false` $\rightarrow$ `true`
- `false || false` $\rightarrow$ `false`

### NOT (`!`)
**Flip** the result.
- `!true` $\rightarrow$ `false`

## 5.3 The Fork (If / Else)
We use `if` to control the flow.

```go
isOpen := true
isHoliday := false

if !isOpen {
    // Stop here if closed
    fmt.Println("Come back tommorow.")
} else if isHoliday {
    // Open, but holiday
    fmt.Println("Holiday hours.")
} else {
    // Normal flow
    fmt.Println("Welcome!")
}
```

## 5.4 The Switch (Cleaner Logic)
If you have many conditions, `if/else` gets messy. Use `switch`.

```go
role := "admin"

switch role {
case "admin":
    fmt.Println("Full Access")
case "user":
    fmt.Println("Limited Access")
default:
    fmt.Println("Access Denied")
}
```
*Think of it as a railway switchyard, directing the train to the correct track.*

::: details 🎓 Knowledge Check
1.  **Does Go use parentheses `()` for `if` conditions?**
    No. We write `if x > 10 { }`.
2.  **What happens if you declare a variable inside an `if` statement?**
    It is only visible inside that `if` block (Scope).
3.  **Does Go have a `while` loop?**
    No. Go only has `for`. We use `for condition { }` to mimic `while`.
:::
