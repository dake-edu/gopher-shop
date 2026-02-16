package worker

import (
	"log/slog"
	"time"
)

// Order represents a job for the worker.
// In a real app, this would be a full struct with UserID, Items, etc.
type Order struct {
	ID    int
	Email string
}

// StartDispatcher is the "Factory Manager".
// It spins up a worker that listens to the "Conveyor Belt" (Channel).
func StartDispatcher(queue <-chan Order) {
	slog.Info("worker dispatcher started", "status", "ready")

	// ⚓ CONCURRENCY PATTERN: The Consumer Loop
	// This runs forever (until channel is closed).
	for order := range queue {
		process(order)
	}
}

// process simulates the heavy lifting (Shipping, Emailing, etc).
func process(order Order) {
	slog.Info("processing order", "order_id", order.ID, "step", "start")

	// Simulate work (e.g., talking to Stripe, printing label)
	time.Sleep(2 * time.Second)

	slog.Info("order shipped", "order_id", order.ID, "status", "completed")
}
