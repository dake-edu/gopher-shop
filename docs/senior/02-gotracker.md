# Chapter 02: GoTracker (The Delivery System)

> **"Amateurs talk strategy. Professionals talk logistics."** - Omar N. Bradley

**GoTracker** is our second independent microservice. It is responsible for one thing: getting the physical book to the customer.

## The Mission
When a user clicks "Buy" in `Gopher Shop`, an event `OrderCreated` is fired.
`GoTracker` listens for this event.

## Tech Stack
*   **Language**: Go 1.24
*   **Transport**: Kafka (Consumer Groups)
*   **Storage**: Redis (Hot State) + Postgres (Archive)
*   **Observability**: Prometheus Metrics ("Parcels Shipped per Second")

## The Domain Model
```go
type Parcel struct {
    ID        string    // UUID
    OrderID   int       // Reference to Gopher Shop Order
    Status    string    // "Pending", "Picked", "Shipped", "Delivered"
    Location  GeoPoint  // Real-time coordinates
}
```

We will build this service from scratch, mimicking a real "Greenfield" project in a big company.
