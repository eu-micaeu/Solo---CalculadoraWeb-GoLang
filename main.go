package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		display := r.FormValue("display")
		result, err := strconv.ParseFloat(display, 64)
		if err != nil {
			http.Error(w, "Erro ao converter o valor", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%.2f", result)
	} else {
		http.ServeFile(w, r, "index.html")
	}
}
