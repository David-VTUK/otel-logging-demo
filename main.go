package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Transaction defines the structure for a single transaction log
type Transaction struct {
	ID         string  `json:"transaction_id"`
	CardNumber string  `json:"card_number"`
	Value      float64 `json:"value"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
}

// A list of currencies to choose from
var currencies = []string{"USD", "EUR", "GBP", "JPY", "AUD", "CAD"}

func main() {

	log.Println("--- Transaction Logger Started ---")

	for {
		// Generate a new random transaction
		tx := generateTransaction()

		// Marshal the transaction struct into JSON format
		jsonData, err := json.Marshal(tx)
		if err != nil {
			log.Printf("Error marshaling transaction to JSON: %v", err)
			continue
		}

		// Log the JSON string
		log.Println(string(jsonData))

		// Wait for a random duration (500ms - 3s) before the next transaction
		sleepDuration := time.Duration(rand.Intn(2500)+500) * time.Millisecond
		time.Sleep(sleepDuration)
	}
}

// generateTransaction creates a new Transaction with random data
func generateTransaction() Transaction {
	return Transaction{
		ID:         uuid.New().String(),
		CardNumber: generateCardNumber(),
		Value:      generateValue(),
		Currency:   generateCurrency(),
		Status:     generateStatus(),
	}
}

// generateCardNumber creates a 16-digit string matching a simple regex like ^(4|5)\d{15}$
func generateCardNumber() string {
	var b strings.Builder

	// Start with 4 (Visa-like) or 5 (Mastercard-like)
	firstDigit := []string{"4", "5"}[rand.Intn(2)]
	b.WriteString(firstDigit)

	// Add 15 random digits
	for i := 0; i < 15; i++ {
		b.WriteString(fmt.Sprintf("%d", rand.Intn(10)))
	}
	return b.String()
}

// generateValue creates a random monetary value between 1.00 and 5000.00
func generateValue() float64 {
	// Generate a value between 1.0 and 5000.0
	val := 1.0 + rand.Float64()*(5000.0-1.0)
	// Round to 2 decimal places
	return math.Round(val*100) / 100
}

// generateCurrency selects a random currency from the global list
func generateCurrency() string {
	return currencies[rand.Intn(len(currencies))]
}

// generateStatus determines if a transaction is "Normal" or "Suspicious"
func generateStatus() string {
	// 90% chance of being "Normal", 10% chance of being "Suspicious"
	if rand.Intn(100) < 10 {
		return "Suspicious"
	}
	return "Normal"
}
