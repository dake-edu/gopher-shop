# 7. Talking to the World (HTTP)

Your Go program is fast and smart. But it is lonely.
To let users verify it, we need **HTTP** (HyperText Transfer Protocol).

## 7.1 The Conversation
The web works on a simple conversation model:
1.  **Client (Browser)**: "Hey, give me the homepage." (**Request**)
2.  **Server (Go)**: "Here is the HTML." (**Response**)

## 7.2 The Handler (The Worker)
In Go, we define a function to handle this conversation.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Logic goes here
}
```

### `w` (The Output Stream)
`http.ResponseWriter`
- Think of this as a **Pipe** connected to the user's browser.
- Whatever you write into this pipe (`fmt.Fprint(w, ...)`) appears on their screen.

### `r` (The Envelope)
`*http.Request`
- This is the letter the user sent you.
- It contains:
    - **Method**: GET (Read), POST (Write).
    - **URL**: What page they want.
    - **Body**: Data they sent.

## 7.3 The Switchboard (Multiplexer)
You need to tell Go which function handles which page.

```go
// "If they ask for '/', call the Home function"
http.HandleFunc("/", HomeHandler)

// "If they ask for '/about', call the About function"
http.HandleFunc("/about", AboutHandler)
```

## 7.4 The Loop (ListenAndServe)
Finally, the server needs to start listening. It enters an infinite loop, waiting for phone calls.

```go
// Listen on port 8081
http.ListenAndServe(":8081", nil)
```
*This program never ends until you kill it.*
