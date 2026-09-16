# Chapter 06: Orchestration (The Captain)

> **"It works on my machine" is not a valid excuse.**

This optional deployment sketch discusses five components; the repository does not include a working GoTracker or Kafka deployment:
1.  Gopher Shop (Web + API)
2.  Postgres (DB)
3.  GoTracker (Microservice)
4.  Redis (Cache)
5.  Kafka (Broker)

Starting them manually (5 terminal windows) is madness.

## 1. Docker Compose (The Local Conductor)
We define the whole orchestra in one file: `docker-compose.yml`.

The following is a structural sketch, not a runnable Compose configuration. It omits database credentials and volumes, Kafka configuration, health checks, and the application Dockerfiles. Do not treat it as an installation recipe.

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

For an actual validated configuration, the command is `docker compose up`. A short `depends_on` orders startup but does not establish service readiness. Add health checks and application retry behavior; pin and test image versions before deployment.

## 2. Kubernetes (The Container Ship) 🚢
Kubernetes is an optional orchestration platform. A small service can run on a single host or a managed application platform; production does not require a cluster. Choose it only when scheduling and operational needs justify its cost.
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
