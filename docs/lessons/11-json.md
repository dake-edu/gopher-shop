# Chapter 11: JSON (The Universal Language)

**Goal**: Learn how to trade data with the outside world.
**Concept**: Your Go program sends structs. The Browser wants JSON (JavaScript Object Notation). We need a translator.

## 1. The Sticker (Struct Tags)
We learned about Structs (The Blueprint). But how does Go know that `BookTitle` in Go code should be `title` in JSON?
**Tags**. Think of them as sticky notes attached to the fields for the translator to read.

```go
type Book struct {
    Title  string  `json:"title"`       // lowercase "title" in JSON
    Price  float64 `json:"price"`
    Hidden string  `json:"-"`           // Ignore this field!
}
```

## 2. Marshaling (Packing)
**Marshal** = Converting Go Struct -> JSON String.
(Ideally, a byte slice, but think of it as text).

```go
myBook := Book{Title: "Go Guide", Price: 19.99}
jsonData, err := json.Marshal(myBook)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(jsonData))
// Output: {"title":"Go Guide","price":19.99}
```

## 3. Unmarshaling (Unpacking)
**Unmarshal** = JSON String -> Go Struct.

```go
jsonStr := `{"title":"Go Guide","price":19.99}`
var newBook Book
err := json.Unmarshal([]byte(jsonStr), &newBook)
```

## 4. Visual Signal: The Translator 🗣️
-   **Go**: Speaks "Binary/Structs".
-   **Web**: Speaks "Text/JSON".
-   **`encoding/json`**: The interpreter standing in the middle, reading the Sticky Notes (Tags) to know what to say.

## Why do we care?
When you build your API in the next level, **every single request** will involve:
1.  Unmarshaling user Input (JSON -> Struct).
2.  Marshaling your Response (Struct -> JSON).
