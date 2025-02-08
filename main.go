package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"unicode"

	"github.com/gorilla/mux"
)

// NumberInfo struct for response
type NumberInfo struct {
	Integer     int      `json:"number"`
	IsPrime     bool     `json:"is_prime"`
	IsPerfect   bool     `json:"is_perfect"`
	Properties  []string `json:"properties"`
	SumOfDigits int      `json:"digit_sum"`
	FunFact     string   `json:"fun_fact"`
}

// ErrorResponse struct for errors
type ErrorResponse struct {
	Integer string `json:"number"`
	Anomaly bool   `json:"error"`
}

// checkPrime verifies if a number is prime
func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// checkPerfect verifies if a number is a perfect number
func checkPerfect(n int) bool {
	if n < 2 {
		return false
	}
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n
}

// checkArmstrong verifies if a number is an Armstrong number
func checkArmstrong(n int) bool {
	temp, sum, digits := n, 0, 0

	// Count number of digits
	for temp != 0 {
		digits++
		temp /= 10
	}

	temp = n // Reset temp

	// Compute sum of each digit raised to the power of digits
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == n
}

// calculateDigitSum finds the sum of digits of a number
func calculateDigitSum(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// determineProperties finds the properties of a number
func determineProperties(n int) []string {
	properties := []string{}
	if checkArmstrong(n) {
		properties = append(properties, "armstrong")
	}
	if n%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}
	return properties
}

// retrieveFunFact fetches a fun fact from Numbers API
func retrieveFunFact(n int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", n)
	resp, err := http.Get(url)
	if err != nil {
		return "Could not fetch fun fact"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading fun fact"
	}

	return string(body)
}

// handleRequest processes API requests
func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := r.URL.Query().Get("number")

	// Check for empty query
	if query == "" {
		errorResponse := ErrorResponse{
			Anomaly: true,
			Integer: "null",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Check if input is alphabetic
	isAlphabetic := true
	for _, ch := range query {
		if !unicode.IsLetter(ch) {
			isAlphabetic = false
			break
		}
	}
	if isAlphabetic {
		errorResponse := ErrorResponse{
			Anomaly: true,
			Integer: "alphabet",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Try converting to an integer
	num, err := strconv.Atoi(query)
	if err != nil {
		errorResponse := ErrorResponse{
			Anomaly: true,
			Integer: "invalid",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Generate response data
	response := NumberInfo{
		Integer:     num,
		IsPrime:     checkPrime(num),
		IsPerfect:   checkPerfect(num),
		Properties:  determineProperties(num),
		SumOfDigits: calculateDigitSum(num),
		FunFact:     retrieveFunFact(num),
	}

	// Send valid JSON response
	json.NewEncoder(w).Encode(response)
}

// startServer initializes the router
func startServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/classify-number", handleRequest).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	startServer()
}

