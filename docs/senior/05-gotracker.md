# Chapter 05: Project "GoTracker"

> **"Theory is nice. Practice is better."**

We are going to build **GoTracker**.
It is a microservice that listens for `orders.created` and simulates shipping.

## 1. The Specification

### Tech Stack
-   **Language**: Go
-   **Storage**: Redis (For package status)
-   **Broker**: Kafka (or Redpanda for dev)
-   **Protocol**: gRPC (for admin queries) + HTTP (for public tracking)

### Features
1.  **Ingest**: Listen to Kafka topic `orders.created`.
2.  **Process**:
    -   Parse JSON.
    -   Generate a "Tracking Number" (UUID).
    -   Save to Redis: `SET tracking:123 "Preparing"`
3.  **Simulate**:
    -   Wait 10 seconds -> Update to "Shipped".
    -   Wait 10 seconds -> Update to "Delivered".
4.  **Query**:
    -   `GET /track/{id}` returns the status.

## 2. Directory Structure
We will create a **separate repository** (or folder) for this.
```text
go-tracker/
├── cmd/
│   └── tracker/
│       └── main.go
├── internal/
│   ├── consumer/ (Kafka Logic)
│   ├── store/    (Redis Logic)
│   └── service/  (Business Logic)
└── docker-compose.yml
```

## 3. Visual Signal (The Delivery Fleet) 🚚
**Concept**: Background Workers.
**Signal**: A fleet of delivery trucks waiting for orders.
-   **Shop**: The Storefront.
-   **Kafka**: The Loading Dock.
-   **GoTracker**: The Trucks that take boxes from the Dock and drive them to the customer.

```mermaid
graph LR
    Shop["🏠 Gopher Shop"] --> |Publish| Kafka{"📮 Kafka"}
    Kafka --> |Consume| Truck1["🚚 Truck 1"]
    Kafka --> |Consume| Truck2["🚚 Truck 2"]
    
    Truck1 --> |Update| Redis[("🎒 Redis")]
    Truck2 --> |Update| Redis
    
    User(("👤 Customer")) -.-> |"GET /track"| Redis
```

## 4. The Challenge
This project is not a tutorial step-by-step.
**You contain the knowledge now.**
You have the blueprints (Middle Level).
You have the tools (Senior Level).

**Build the Fleet.**
