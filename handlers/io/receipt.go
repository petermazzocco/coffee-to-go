package ioPkg

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CreateReceipt() {
	// Create a receipt file
	f, err := os.Create("receipt.txt")
	if err != nil {
		log.Println("Error creating receipt:", err)
		return
	}
	// Defer the close of the file
	defer f.Close()

	// Write the receipt content to the file
	r, err := createReceiptString("Total Paid: $6.99", f)
	if err != nil {
		log.Println("Error creating receipt:", err)
		return
	}
	fmt.Printf("Receipt created successfully! %d\n", r)
}

func ReadReceipt() {
	f, err := os.Open("receipt.txt")
	if err != nil {
		log.Println("Error finding receipt, did you make one yet?:", err)
		return
	}

	defer f.Close()

	r, err := readReceipt(f)
	if err != nil {
		log.Println("Error reading receipt:", err)
		return
	}
	fmt.Printf("Receipt read successfully! %d\n", r)
}

// Create the receipt content
func createReceiptString(s string, w io.Writer) (int, error) {
	// Write the receipt content to the writer
	i, err := w.Write([]byte(s))
	if err != nil {
		return 0, err
	}
	return i, nil
}

func readReceipt(r io.Reader) (int, error) {
	i, err := r.Read([]byte{})
	if err != nil {
		return 0, err
	}
	return i, nil
}
