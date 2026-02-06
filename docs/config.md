# 9. The Wiring (Configuration)

## 9.1 The "Twelve-Factor" Philosophy
Why do we do this?
Hardcoding values (like `port = 8080`) is the mark of a beginner.
Professional apps follow the **12-Factor App** methodology.
**Factor III**: "Store config in the environment".

### Comparison: Configuration
| Language | Beginner Approach | Professional Approach |
| :--- | :--- | :--- |
| **Python** | `PORT = 8000` top of file | `os.environ.get('PORT')` |
| **Java** | `application.properties` | System Properties / Env Vars |
| **Go** | `const Port = ":8080"` | `os.Getenv("PORT")` |

## 9.2 Reading the Environment
```go
func LoadConfig() Config {
    return Config{
        Port: os.Getenv("PORT"),
    }
}
```

### Anatomy of `os.Getenv`
1.  **`os`**: Standard library package for Operating System functions.
2.  **`Getenv`**: "Get Environment Variable".
3.  **`("PORT")`**: The key we are looking for.
    - If the computer has `PORT=9090` set in its memory, this function returns `"9090"`.
    - If not, it returns `""` (Empty String).

Professional apps often assume a default if the result is empty:
```go
port := os.Getenv("PORT")
if port == "" {
    port = "8081" // Default fallback
}
```

::: details 🎓 Knowledge Check: Why do we read configuration from "Environment Variables"?
**Answer**: To follow the **12-Factor App** methodology. This allows us to change settings (changing ports, database passwords) without changing the code itself or recompiling.
:::

