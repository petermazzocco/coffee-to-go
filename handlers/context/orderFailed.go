package context

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

// Let's say we need to quickly query the database of our
// supplier to see if we have any available products.
// If the query takes longer than the timeout, we'll cancel the request.

func OrderNewProducts() {
	// Increase the value of the duration to see what happens
	duration := time.Duration(100) * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	url := "https://api.sampleapis.com/coffee/hot"

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
