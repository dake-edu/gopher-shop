# The Bridge: Crossing to Professional Engineering

Congratulations! You have built a functional web application in **one single file** (the Demo).
It works. It's fast. It's simple.

**But...**
Imagine if you had 50 engineers working on that one file.
Imagine if `main.go` was 50,000 lines long.
Imagine if you wanted to swap "In-Memory" storage for a database without rewriting everything.

## The Professional Leap
We typically refactor for three reasons:

### 1. Cognitive Load (The Brain RAM)
*   **Junior (All-in-One)**: You have to hold the entire program in your head to change one line.
*   **Professional (Modules)**: You only need to understand the module you are working on.

### 2. Dependency Injection (The Universal Plug)
*   **Junior**: Often hardcodes dependencies (e.g., global variables).
*   **Professional**: Passes dependencies (Interfaces) to functions. This makes testing easier because you can pass a "Fake Database" during tests.

### 3. Protection (The Internal Folder)
Go has a special directory name: `internal/`.
The compiler allows imports only from within the tree rooted at the parent of `internal`. This limits package dependencies; it does not restrict access to application data.

## The Map
Here is how your Single File concept maps to the Professional Project Structure:

| Concept | Junior (Demo) | Professional (Main API) |
| :--- | :--- | :--- |
| **Startup** | `main.go` | `cmd/api/main.go` |
| **Data Models** | `type Book struct` (in main) | `internal/models/book.go` |
| **Storage** | Global `var books []Book` | `internal/store/postgres` |
| **Handlers** | `http.HandleFunc("/add")` | `cmd/api/handlers.go` |
| **HTML** | String Constant | `web/templates/` |

## Next Steps
Now that you understand **WHAT** we are building (The Shop), let's learn **HOW** to build it professionally.
Proceed to **Phase 3: The Professional**.

::: details 🎓 Knowledge Check: What is special about the `internal/` directory?
**Answer**: The importing code must be inside the tree rooted at the parent of `internal`. The rule is about import paths, not authentication or secret data.
:::


The precise `internal` rule: an import is allowed only from code inside the tree rooted at the parent of that `internal` directory. This is an import boundary, not data privacy or an authorization mechanism.
