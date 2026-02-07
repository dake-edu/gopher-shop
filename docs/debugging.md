# The Detective: Debugging

> **"Novices guess. Professionals inspect."**

You have written code. You ran it. It output `0`, but should have output `100`. What now?
A beginner starts random guessing and sprinkling `fmt.Println("I am here 1")` everywhere. This is the "Poke Method".

An Engineer takes a magnifying glass and **stops time**. This is called **Debugging**.

## The Magnifying Glass (VS Code & Delve)

Imagine you can freeze your program at any second, crawl inside its memory, and see exactly what is inside every variable.

### 1. Setting a Breakpoint
In VS Code, hover your mouse to the left of the line numbers. A **red dot** will appear. Click it.
This is a **Breakpoint**. It tells the computer: *"Run full speed until you hit this line, then FREEZE!"*

### 2. Inspecting Variables
Press `F5` (Start Debugging).
The program will start and then... hang. It hasn't crashed. It's waiting for you.
Look at the **Variables** panel on the left. You will see every variable currently alive.
- `i`: `0`
- `sum`: `0`

### 3. Step Over (F10)
Press `F10`. The program executes **one single line** and freezes again.
Watch the variables change in real-time.
- `sum` becomes `5`.
- `i` becomes `1`.

You are no longer guessing. You are watching the crime happen in slow motion.

::: details 🎓 Knowledge Check: Why not just use Println?
**Answer**: `Println` clutters your code and requires a recompile every time you want to check a different variable. Debugging lets you check *everything* without changing a single line of code.
:::
