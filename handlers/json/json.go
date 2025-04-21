package jsonPkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Title struct {
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	TotalSales  int    `json:"totalSales"`
}

func UnmarshalJson() {
	resp, err := http.Get("https://api.sampleapis.com/coffee/hot")
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Here is the full json body (uncomment to see it)
	// It needs to be wrapped in a string because it's in bytes
	// fmt.Printf("here is their full list %s\n", string(body))

	var t []Title
	if err := json.Unmarshal(body, &t); err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, title := range t {
		fmt.Println("Title:", title.Title) // Change this to see other properties
	}
}
