# 9. The Wiring (Configuration)

A professional machine has a Control Panel. We don't hardwire settings deep inside the engine.

## 9.1 The Control Room
In our app, the `internal/config` package is the **Control Room**.
Instead of writing `port := "8081"` inside `main.go`, we ask the centralized configuration.

## 9.2 The "Twelve-Factor App" Rule
Professional apps follow the **12-Factor Methodology**.
**Rule #3**: Store config in the **Environment**, not in the code.

### Why?
1.  **Security**: You don't want Database Passwords committed to GitHub!
2.  **Flexibility**: You can change the Port or Database URL without recompiling the code.

## 9.3 Environment Variables
We use `os.Getenv` to read from the system.

```go
func LoadConfig() Config {
    return Config{
        Port: os.Getenv("PORT"), // Reads from the machine's settings
    }
}
```

**Visual Anchor**:
- **Hardcoding**: Welding a switch in the "On" position.
- **Config**: Installing a switch that can be flipped from outside.
