package goroutines

import (
	"context"
	"log"
	"math/rand"
	"time"
)

// We need to constantly monitor the temperature of our
// espresso machine so that it doesn't over heat

var (
	lowestTemp  = 90.0
	highestTemp = 110.0
	arrayLength = 100
)

func generateRandomFloat64Array() []float64 {
	min := 85.0
	max := 115.0
	array := make([]float64, arrayLength)
	for i := range arrayLength {
		randomNumber := min + rand.Float64()*(max-min)
		array[i] = float64(randomNumber)
	}
	return array
}

// Goroutine to continously monitor temp of our espresso machine
func monitor(ctx context.Context, c chan float64, a chan string, temps []float64) error {
	for {
		select {
		// After one second
		case <-time.After(time.Second):
			for _, temp := range temps {
				duration := time.Duration(int64(rand.Intn(100)))
				time.Sleep(duration * time.Millisecond)
				// If the temp is out of range, cancel the context
				if temp < lowestTemp || temp > highestTemp {
					a <- "Espresso machine temperature out of range!"
					ctx.Done()
					// If the temp is in range, send it to the channel
				} else if temp >= lowestTemp && temp <= highestTemp {
					c <- temp
				}
			}
			close(c)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func TempCheck() {
	temps := make(chan float64, arrayLength)
	alerts := make(chan string, arrayLength)
	tempArray := generateRandomFloat64Array()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go monitor(ctx, temps, alerts, tempArray)

	for t := range temps {
		log.Printf("Machine temperature: %.2f", t)
	}

	a := <-alerts
	close(alerts)
	log.Printf("🚨: %s", a)

}
