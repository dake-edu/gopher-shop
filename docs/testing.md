# Chapter 25: Testing

## 25.1 Table-Driven Tests
In other languages, you might write 10 separate test functions for 10 cases.
In Go, we use a **Table**. It's cleaner.

```go
func TestAdd(t *testing.T) {
    // The Table
    cases := []struct {
        A, B, Expected int
    }{
        {1, 1, 2},
        {5, 0, 5},
        {-1, 1, 0},
    }

    // The Loop
    for _, tc := range cases {
        result := Add(tc.A, tc.B)
        if result != tc.Expected {
            t.Errorf("Add(%d, %d) = %d; want %d", tc.A, tc.B, result, tc.Expected)
        }
    }
}
```

### Anatomy of `t *testing.T`
- **`t`**: The Controller.
- **`t.Errorf`**: "Mark this test as Failed, log the message, but **Continue** running other tests."
- **`t.Fatal`**: "Stop now. This is critical."

## 25.2 Mocks (The Stunt Double)
Why do we mock?
If your test actually connects to GitHub, and your internet is down, your test fails.
Code logic didn't break; the internet did.
**Tests must confirm LOGIC, not INFRASTRUCTURE.**

We create a fake struct that "looks like" the real dependency (satisfies the Interface) but just returns dummy data.

::: details 🎓 Knowledge Check: Why do we use "Table-Driven Tests"?
**Answer**: To avoid code duplication. Instead of writing 10 separate test functions for 10 scenarios, we write **one** logic loop and feed it a slice (Table) of inputs and expected outputs.
:::

