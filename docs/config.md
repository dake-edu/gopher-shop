# Centralized Configuration

## Single Source of Truth

> [!TIP]
> ⚓ **Visual Anchor:** The Control Room

The `internal/config` package acts as the **Control Room** for the entire application.
Instead of having settings scattered across multiple files (hardcoded strings), everything is managed in one place.

## Security & The "Twelve-Factor App"

We strictly use **Environment Variables** for configuration.
- **Code** goes to GitHub (Public/Private).
- **Secrets** (Passwords, Keys) stay on the Server.

**Why?**
If you hardcode a password in Go code and push it to GitHub, it's leaked forever.
By using `os.Getenv`, the code remains clean and safe.
