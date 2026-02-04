# Application Lifecycle & Graceful Shutdown

## The Cleanup Crew

> [!TIP]
> ⚓ **Visual Anchor:** The Clean Up Crew (Janitor)

When you close a shop, you don't just kill the lights while customers are at the register.
1.  **Stop New Customers**: Lock the door (Stop listening for requests).
2.  **Finish Existing Sales**: Let current customers finish (Wait for active requests).
3.  **Clean Up**: Sweep the floor (Close DB connections).
4.  **Go Home**: Exit the process.

## Why is this important?

If you kill the process (`SIGKILL` or Pulling the plug):
- **Data Corruption**: Half-written database transactions might be lost or corrupted.
- **Client Errors**: Users expecting a response get a "Connection Reset" error.

By using `server.Shutdown()`, we respect our data and our users.

## Health Check

The `/health` endpoint is not just a static "OK". It's a live probe.
- It pings the Database.
- If the DB is down, it returns `503 Service Unavailable`.
- This tells load balancers to stop sending traffic to this broken instance.
