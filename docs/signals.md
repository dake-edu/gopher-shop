# Visual Anchors (The Pedagogical Signals)

To ensure "No Dark Zones", we use consistent visual metaphors throughout the course.
This catalog tracks every signal used to explain an abstract concept.

## 1. The Container (Structs)
*A Blueprint for a composite object.*

```mermaid
classDiagram
    class Box {
      +ID: int
      +Value: string
    }
    note for Box "A Struct is a custom Shape."
```

## 2. The Shelf (Slices)
*A window into a growing array.*

```mermaid
graph LR
    A[Index 0] --- B[Index 1] --- C[Index 2]
    style A fill:#f9f,stroke:#333
    style B fill:#f9f,stroke:#333
    style C fill:#fff,stroke:#333,stroke-dasharray: 5 5
    note[Slice View]
```

## 3. The Conveyor Belt (Loops)
*Processing items one by one.*

```mermaid
graph LR
    Input[Items] --> Process((Range Loop))
    Process --> Output[Item 1]
    Process --> Output2[Item 2]
```

## 4. The Quality Gate (Validation)
*Previously "The Club Bouncer"*
*Rejects bad data immediately.*

```mermaid
graph TD
    Data[Incoming Data] --> Gate{Quality Gate}
    Gate -- Valid --> Warehouse[(Database)]
    Gate -- Invalid --> Reject[Return Error]
    style Gate fill:#007d9c,color:white
```

## 5. The Universal Plug (Interfaces)
*Dependency Injection.*

```mermaid
graph LR
    Socket((Interface))
    Plug1[Postgres] --> Socket
    Plug2[Memory] --> Socket
```
