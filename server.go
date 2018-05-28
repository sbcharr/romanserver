package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	roman "github.com/squeakysimple/romanserver/romannumerals"
)

const (
	statusBadRequest = http.StatusBadRequest
	statusNotFound   = http.StatusNotFound
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", romanNumeralHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func romanNumeralHandler(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	if urlPathElements[1] == "roman_number" {
		number, err := strconv.Atoi(urlPathElements[2])
		if err != nil {
			inputErrorHandler(w, 400)
			return
		}
		if number < 1 || number > 15 {
			inputErrorHandler(w, 404)
		} else {
			fmt.Fprintf(w, "%q", html.EscapeString(roman.Numerals[number]))
		}
	} else {
		inputErrorHandler(w, 400)
	}
}

func inputErrorHandler(w http.ResponseWriter, status int) {
	switch status {
	case 400:
		w.WriteHeader(statusBadRequest)
		w.Write([]byte("400 - Bad Request"))
	case 404:
		w.WriteHeader(statusNotFound)
		w.Write([]byte("404 - Not Found"))
	}
}
