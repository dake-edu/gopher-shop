# The Anatomy of Code

To a beginner, code looks like matrix hieroglyphs.
Let's demystify it. We will explain **every single symbol** in a basic program.

```go
package main              // 1. The Package Declaration

import "fmt"              // 2. The Import Statement

func main() {             // 3. The Function Declaration
    fmt.Println("Hello")  // 4. The Statement
}
```

## 1. `package main`
**"This is an executable program."**
- In Go, code is organized into boxes called **Packages**.
- The `main` package is special. It tells the compiler: *"Produce a program I can run, not just a library helper."*
- If you named it `package calculator`, it would be a library, not a runnable app.

## 2. `import "fmt"`
**"Bring in the tools."**
- **Keyword**: `import` tells Go we need extra functionality.
- **String**: `"fmt"` stands for **F**or**m**a**t**. It is a standard library package for input/output (printing text).
- **Metaphor**: You are a carpenter. You are reaching into your toolbox to grab the "Hammer" (fmt).

## 3. `func main() { ... }`
**"The Entry Point."**
- **Keyword**: `func` defines a **Function** (a block of logic).
- **Name**: `main`. When you run the program, Go looks strictly for a function named `main` to start.
- **`()` Parentheses**: These hold **Arguments** (Inputs). Here, it is empty, meaning the function takes no inputs.
- **`{ }` Curly Braces**: These mark the **Scope** (The Body). Everything *inside* these braces is part of the function.
    - `{` = Start of logic.
    - `}` = End of logic.

## 4. `fmt.Println("Hello")`
**"Do something."**
- **`fmt`**: The package we imported.
- **`.` (The Dot)**: The **Selector**. It means *"Open the `fmt` box AND look inside for..."*
- **`Println`**: The specific tool we want. (Print Line).
- **`()`**: **Invocation**. *"Activate this tool now!"*
- **`"Hello"`**: The **Data**. We pass this text *into* the tool.
    - **`""` Quotes**: Tell Go this is text (String), not code code.

---
**Summary**:
"I am a runnable program (`package main`). I need the format tool (`import`). When I start (`func main`), open the format tool (`.`) and use the Print Line feature to show 'Hello' (`Println(...)`)."
