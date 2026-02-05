# Your First Web Server

In Step 0, we printed books to the console. Now, let's share them with the world using a Web Server.

## The Visual Anchor: The Waiter

Web servers can seem magical, but they work just like a restaurant.

### 1. The Customer (Client)
Your Web Browser (Chrome, Safari) is the customer. It wants information (a web page).

### 2. The Order (Request)
When you type `http://localhost:8081` and hit Enter, the browser sends a **Request** to the server. It says: *"I would like the homepage, please."*

### 3. The Waiter (Handler)
This is the Go code we write. We tell Go: *"When you get an order for '/', run this function."*
The function takes the order, looks at our bookshelf (Slice), and prepares the "food".

### 4. The Food (Response)
The server sends back an **HTML Response**. The browser receives this and displays it as a formatted web page.

## The Code

We use the standard `net/http` package. No external tools needed!

```go
// "Hiring" the waiter for the root path "/"
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // Writing the response (The Food)
    fmt.Fprintf(w, "<h1>Hello World!</h1>")
})

// Opening the restaurant on port 8081
http.ListenAndServe(":8081", nil)
```

### Try it out!
1. Run the server:
   ```bash
   go run lessons/02-web-basics/main.go
   ```
2. Open your browser and go to `http://localhost:8081`.
