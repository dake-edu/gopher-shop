# 8. The Blueprint (Data Modeling)

Before we build a house, we need a blueprint. In Go, we define our data using **Structs**.

## 8.1 The Struct (The Container)
As we learned in Chapter 1, a `struct` is a collection of fields. But for a professional app, we need more detail.

```go
type Book struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Author  string  `json:"author"`
    Price   float64 `json:"price"`
}
```

## 8.2 Tags (The Sticky Notes)
Notice the text in backticks: `` `json:"title"` ``.
These are **Struct Tags**.
- **Metaphor**: Sticky notes attached to the blueprint fields.
- **Purpose**: They tell Go how to translate this struct into other formats (like JSON for the web).
- **Example**: "When you send this to the browser, name this field `title` (lowercase), not `Title`."

## 8.3 Domain Rules
Our blueprint also implies rules:
- **ID**: Must be unique.
- **Price**: Cannot be negative.
- **Title**: Cannot be empty.

We will enforce these rules later using "The Inspector" (Chapter 11).
