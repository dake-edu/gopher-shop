# 16. The Robot (CI/CD)

We configure the robot using **YAML** (Yet Another Markup Language).
It's just key-value pairs, like JSON but without brackets.

## 16.1 Anatomy of a Workflow File
File: `.github/workflows/go-check.yml`

```yaml
name: Go Check                # 1. The Name of the Workflow
on: [push, pull_request]      # 2. Even Triggers

jobs:
  test:                       # 3. The Job Name
    runs-on: ubuntu-latest    # 4. The Computer (Runner)
    steps:
      - uses: actions/checkout@v3  # 5. Step: Download Code
      - uses: actions/setup-go@v4  # 6. Step: Install Go
        with:
          go-version: '1.21'
      - name: Test
        run: go test ./...    # 7. Step: Run Command
```

### Breaking it Down
1.  **`on`**: When should the robot wake up? (On every Push).
2.  **`runs-on`**: GitHub gives us a free computer running Ubuntu Linux.
3.  **`uses`**: These are pre-made scripts ("Actions") written by others.
    - `checkout`: Downloads your repository.
    - `setup-go`: Installs the Go compiler.
4.  **`run`**: This executes a command in the terminal of that Linux computer.

## 16.2 "Green Build"
If `go test` exits with code `0` (Success), the build is Green.
If it exits with `1` (Error), the build is Red, and GitHub turns the Merge Button grey (if configured).
This saves you from breaking the Main branch.
