# Unit Testing & Mocking

## The Testing Pyramid

> [!TIP]
> ⚓ **Visual Anchor:** The Pyramid

Stable software is built like a pyramid:
1.  **Unit Tests (Base)**: Fast, cheap, covers individual functions (e.g., `Book.Validate`).
2.  **Integration Tests (Middle)**: Covers interactions (e.g., Database queries).
3.  **e2e Tests (Top)**: Slow, covers full user flows.

We focus heavily on the **Base** (Unit Tests) because they give instant feedback.

## Tests as "Living Documentation"

Tests are better than comments because they **never lie**.
If you want to know how `Book.Validate` works, look at `internal/models/book_test.go`. The test cases show exactly what is allowed and what is rejected.

## Mocking

To test the API handlers without a real database, we use **Mocking**.
The `internal/store/mock.go` file allows us to simulate the database behavior, making our tests deterministic and fast.
