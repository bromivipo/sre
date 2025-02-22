package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
var count_tries = 0
var every_nth = 0

func main() {
	http.Handle("/target-handler", http.HandlerFunc(targetHandler))
	http.Handle("/internal/set-every-nth", http.HandlerFunc(internalHandler))
	http.ListenAndServe(":8080", nil)
}

func targetHandler(w http.ResponseWriter, r *http.Request) {
	if every_nth != 0 {
		if count_tries != 0 {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			count_tries = (count_tries + 1) % every_nth
			return
		} else {
			count_tries = (count_tries + 1) % every_nth
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["response"] = "Response 200"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func internalHandler(w http.ResponseWriter, r *http.Request) {
	n, ok := r.URL.Query()["n"]
	n_int, _ := strconv.Atoi(n[0])
	if !ok || n_int < 0{
		http.Error(w, "Параметр 'n' отсутствует или < 0", http.StatusBadRequest)
	} else {
		every_nth = n_int
		count_tries = 1
		fmt.Fprintf(w, "Ок\n")
	}
}