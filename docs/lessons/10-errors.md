# Chapter 10: Errors (The Guard Rails)

**Goal**: Learn how to say "No" safely.
**Concept**: In Go, errors are not "Exceptions" that crash your program. They are just **values**, like a string or an integer. Generally, they are the polite refusal of a function to do what you asked.

## 1. The Logic of "Maybe"
In many languages, if something goes wrong, the program *explodes* (Exception) unless you catch it.
In Go, functions simply return two things:
1.  The Result (if successful)
2.  The Error (if failed)

```go
func Divide(a, b int) (int, error) {
    if b == 0 {
        // We cannot divide by zero. Return 0 and an error.
        return 0, fmt.Errorf("cannot divide by zero")
    }
    // Success! Return result and nil (no error).
    return a / b, nil
}
```

## 2. strict "if err != nil"
The most famous pattern in Go. You will type this thousands of times. Memorize the rhythm.

```go
result, err := Divide(10, 0)
if err != nil {
    // 🚨 HANDLE IT!
    fmt.Println("Error:", err)
    return
}
// ✅ PROCEED
fmt.Println("Result:", result)
```

> [!IMPORTANT]
> **The Golden Rule**: Never ignore an error. Even if you think "it will never happen." It will.

## 3. Visual Signal: The Fork in the Road 🛤️
Imagine a railroad switch.
-   **Main Track**: The Happy Path (Success).
-   **Side Track**: The Error Path.

Every time you call a function that returns an error, you are at a switch. You MUST decide where the train goes.

## 4. Why is this better?
It forces you to think about failure **immediately**.
-   **Java/Python**: "I'll try this block and maybe catch exceptions later." (Implicit failure)
-   **Go**: "This function *might* fail. I must handle that right now." (Explicit failure)

## Checkpoint
-   Errors are values.
-   `nil` means "no error".
-   Always check `if err != nil`.
