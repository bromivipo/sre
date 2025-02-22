package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CircuitBreaker struct {
	errors_count    int
	state           string
	timeout_end	    time.Time
   }
   
var cb = CircuitBreaker{errors_count: 0, state: "closed", timeout_end: time.Now()}

func main() {
	http.Handle("/call-get-data-with-circuit-breaker", http.HandlerFunc(handler))
	http.ListenAndServe(":80", nil)
}

func GetDataWithCircuitBreaker(url string) (string, error) {
	if cb.state == "open" {
		if cb.timeout_end.Sub(time.Now()) < 0 {
			cb.state = "half"
			log.Print("Circuit Breaker state is now half-open. Allowing 1 test request\n")
		} else {
			log.Printf("Circuit Breaker state is open. No requests allowed until %s\n", cb.timeout_end.String())
			return "", fmt.Errorf("Circuit Breaker state is open. No requests allowed until %s\n", cb.timeout_end.String())
		}
	}
	var bodyString string
	resp, err := http.Get(url)
	if err != nil {
		return bodyString, err
	}
	if resp.StatusCode < 500 {
		log.Printf("Got response with status %d\n", resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		bodyString = string(body)
		cb.errors_count = 0
		if cb.state == "half" {
			cb.state = "closed"
			cb.errors_count = 0
			log.Print("Circuit Breaker state is now closed. Allowing all requests\n")
		}
		return bodyString, nil
	}
	err = fmt.Errorf("Got no response with status < 500")
	cb.errors_count++
	if cb.state == "half" || cb.errors_count == 3 {
		cb.state = "open"
		cb.timeout_end = time.Now().Add(10 * time.Second)
		log.Printf("Another error. Circuit Breaker state is now open. No requests allowed until %s\n", cb.timeout_end.String())
	} else {
		log.Print("Error encountered\n")
	}
	return bodyString, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	s, err := GetDataWithCircuitBreaker("http://localhost:8080/target-handler")
	if err == nil {
		w.Write([]byte(s))
	} else {
		w.Write([]byte("no response"))
	}
}