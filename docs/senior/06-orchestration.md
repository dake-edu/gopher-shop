# Chapter 06: Orchestration (The Captain)

> **"It works on my machine" is not a valid excuse.**

You have 3 services now:
1.  Gopher Shop (Web + API)
2.  Postgres (DB)
3.  GoTracker (Microservice)
4.  Redis (Cache)
5.  Kafka (Broker)

Starting them manually (5 terminal windows) is madness.

## 1. Docker Compose (The Local Conductor)
We define the whole orchestra in one file: `docker-compose.yml`.

```yaml
services:
  shop:
    build: ./gopher-shop
    ports: ["8080:8080"]
    depends_on: [postgres, kafka]
  
  tracker:
    build: ./go-tracker
    depends_on: [redis, kafka]

  postgres:
    image: postgres:15
  
  redis:
    image: redis:7
    
  kafka:
    image: bitnami/kafka:latest
```

Command: `docker-compose up`.
Result: The whole city starts up.

## 2. Kubernetes (The Container Ship) 🚢
When you go to production (AWS, Google Cloud), you use **Kubernetes (K8s)**.
K8s is like Docker Compose, but across 100 servers.
-   **Pod**: A small group of containers (e.g., 1 Shop instance).
-   **Deployment**: "I want 3 replicas of the Shop."
-   **Service**: "Give me a Load Balancer IP to talk to the Shop."

## 3. Visual Signal (The Container Ship) ⚓
**Concept**: Container Orchestration.
**Signal**: A Captain managing a massive ship.
-   **Containers**: The standard metal boxes.
-   **Docker**: The crane that lifts them.
-   **Kubernetes**: The Captain who decides where each box goes to balance the ship.

```mermaid
graph TD
    Captain[👮 Captain (K8s Master)]
    
    subgraph Ship ["Node 1"]
        C1["📦 Container: Shop"]
        C2["📦 Container: Shop"]
    end
    
    subgraph Ship2 ["Node 2"]
        C3["📦 Container: Tracker"]
        C4["📦 Container: Redis"]
    end
    
    Captain --> |Schedule| Ship
    Captain --> |Schedule| Ship2
```

## 4. Final Words on Ops
As a Backend Engineer, you don't need to be a DevOps expert.
But you MUST understand how your code runs.
-   How do I view logs? (`kubectl logs`)
-   How do I restart it? (`kubectl rollout restart`)
-   How do I scale it? (`kubectl scale --replicas=5`)
