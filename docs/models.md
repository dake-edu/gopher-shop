# 8. The Blueprint (Data Modeling)

## 8.1 The Struct (The Container)
In Java or Python, you might use a `Class`. In Go, we use a `struct`.
A struct is purely data. It has no hidden machinery.

```go
type Book struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Price   float64 `json:"price"`
}
```

### Anatomy of the Code
1.  **`type`**: New Type Definition. We are inventing a custom data type named `Book`.
2.  **`struct`**: The kind of type. It means "Structure", a composite of fields.
3.  **`ID string`**: Field Name (`ID`) and Data Type (`string`).
    - *Note*: Capitalized `ID` means **Public** (Exported). If we wrote `id` (lowercase), other packages could not see it.
4.  **`` `json:"id"` ``**: The Tag. (See 8.2).

### Comparison: Struct vs Class
| Feature | Java Class | Python DataClass | Go Struct |
| :--- | :--- | :--- | :--- |
| **Inheritance** | Yes (`extends`) | Yes | **No (Composition)** |
| **Methods** | Inside Class | Inside Class | **Attached Separately** |
| **Visibility** | `public`/`private` | `_` convention | **Capitalized (Public)** |

## 8.2 Tags (Reflecting Instructions)
What is that stuff in backticks? `` `json:"id"` ``?

This is **Metadata**.
Imagine you are handing this struct to a specialized robot called the "JSON Encoder".
The robot reads the tag:
> *"Oh, the human called this field `ID` in Go, but when I turn it into text for the browser, I should label it `id` (lowercase)."*

If you omit tags, the JSON output will look like `{"ID": "123"}`, which is not standard for web APIs (which prefer `snake_case` or `camelCase`).
