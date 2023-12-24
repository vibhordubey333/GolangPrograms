package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	Amount float64
}

/*
	Stage 1 (Validation)

Each order is received and validated for correctness and completeness. Any invalid orders are filtered out.
*/
func validateOrders(done <-chan struct{}, input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	for {
		select {
		case <-done:
			return
		case order, ok := <-input:
			if !ok {
				return
			}

			// Simulating validation
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Validate order
			if order.Amount > 0 {
				output <- order
			}
		}
	}
}

/*
	Stage 2 (Enrichment)

The valid orders are then enriched with additional information, such as customer details or product data, to enhance their content.
*/
func enrichOrders(done <-chan struct{}, input <-chan Order, output chan<- Order) {
	defer close(output) // Close the output channel when the function finishes

	for {
		select {
		case <-done:
			return
		case order, ok := <-input:
			if !ok {
				return
			}

			// Simulating enrichment
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Enrich order
			order.Amount *= 1.1
			output <- order
		}
	}
}

/*
 Stage 3 (Calculation)

The enriched orders are processed further to perform calculations, such as total order value or shipping costs.
*/

func calculateOrderValues(done <-chan struct{}, input <-chan Order, output chan<- Order) {
	defer close(output) // Signal the wait group that the processing is done for this goroutine

	for {
		select {
		case <-done:
			return
		case order, ok := <-input:
			if !ok {
				return
			}

			// Simulating calculation
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			// Perform calculation
			order.Amount += 5.0
			output <- order
		}
	}
}

/*
Each stage is implemented as a separate goroutine, and channels are used to connect the stages. The orders flow through the pipeline, with each stage
concurrently processing the orders it receives. This allows for parallelism, as multiple orders can be processed simultaneously by different stages.
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	// Create done channel to signal termination
	done := make(chan struct{})
	defer close(done)

	// Create channels to connect the stages
	orders := make(chan Order)
	validOrders := make(chan Order)
	enrichedOrders := make(chan Order)
	calculatedOrders := make(chan Order)

	// Start the stages concurrently
	go validateOrders(done, orders, validOrders)
	go enrichOrders(done, validOrders, enrichedOrders)
	go calculateOrderValues(done, enrichedOrders, calculatedOrders)

	// Generate sample orders
	go func() {
		for i := 1; i <= 10; i++ {
			order := Order{
				ID:     i,
				Amount: float64(i * 10),
			}
			orders <- order
		}
		close(orders)
	}()

	// Receive the final output from the pipeline
	for order := range calculatedOrders {
		fmt.Printf("Processed order ID: %d, Final amount: %.2f\n", order.ID, order.Amount)
	}
}

/*
Output:
Processed order ID: 1, Final amount: 16.00
Processed order ID: 2, Final amount: 27.00
Processed order ID: 3, Final amount: 38.00
Processed order ID: 4, Final amount: 49.00
Processed order ID: 5, Final amount: 60.00
Processed order ID: 6, Final amount: 71.00
Processed order ID: 7, Final amount: 82.00
Processed order ID: 8, Final amount: 93.00
Processed order ID: 9, Final amount: 104.00
Processed order ID: 10, Final amount: 115.00
*/
