# Chapter 11: The Basic Server

**Goal**: Make the computer listen and talk back.
**Concept**: `net/http` is the standard library for web servers.

<<< @/../lessons/01-server/main.go

::: details 🎓 Knowledge Check: What does `http.ListenAndServe` do?
**Answer**: It starts the server and keeps it running (listening) in a loop. Without it, your program would exit immediately!
:::

