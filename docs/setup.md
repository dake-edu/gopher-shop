# The Workbench (Setup)

A master craftsman needs sharp tools. Let's set up your environment.

## 1. Install Go (The Compiler)
Go is a compiled language. You need the "Compiler" to turn your text code into a binary program that the computer understands.

- **Download**: [go.dev/dl](https://go.dev/dl/)
- **Verify**: Open your terminal (Command Prompt) and type:
  ```bash
  go version
  ```
  You should see something like `go version go1.21...`.

## 2. The Code Editor (VS Code)
We recommend **Visual Studio Code (VS Code)**. It is free and powerful.
1. Download [VS Code](https://code.visualstudio.com/).
2. **Crucial Step**: Install the **Go Extension** by the Go Team at Google.
   - Click the "Extensions" box on the left sidebar.
   - Search for "Go".
   - Install the one with the blue Gopher icon.

## 3. The Commands (Your Toolbar)

You will use three main commands in your terminal:

### A. `go mod init [name]`
**"Initialize Project"**
This creates a `go.mod` file. Think of it as the "Passport" for your project. It tells Go clearly who you are.
```bash
go mod init github.com/myname/myproject
```

### B. `go run main.go`
**"Test Run"**
This compiles your code in a temporary folder and runs it immediately. We use this 99% of the time while developing.

### C. `go build`
**"Manufacture"**
This creates a permanent binary file (executable). You can send this file to a friend, and they can run it **without** having Go installed!

---
**Next**: Now that the shop is open, let's learn how to read the blueprints.
