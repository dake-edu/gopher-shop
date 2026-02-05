# 15. The Safety Net (Testing)

You wouldn't drive a car that hasn't been crash-tested. Don't ship code that hasn't been unit-tested.

## 15.1 The Crash Test Dummy (Unit Tests)
A **Unit Test** checks one small part of the machine in isolation.
In Go, test files end in `_test.go`.

```go
func TestAddRef(t *testing.T) {
    result := Add(2, 2)
    if result != 4 {
        t.Errorf("Expected 4, got %d", result)
    }
}
```

## 15.2 Mocking (The Stunt Double)
When testing the Web Handler, we don't want to connect to a real Database (it's slow and risky).
We use a **Mock** (a fake database) that pretends to work.

- **Real DB**: "I will connect to port 5432..."
- **Mock DB**: "I will just return 'Success' immediately."

This allows us to test the *Logic* without the *Infrastructure*.
