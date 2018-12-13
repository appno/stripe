package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/appno/stripe/schema"
)

func handler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var doc interface{}
	err := decoder.Decode(&doc)
	if err != nil {
		panic(err)
	}

	result, err := schema.DocumentValidator.IsCompliant(doc)
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = io.WriteString(w, string(bytes))

	if err != nil {
		panic(err)
	}
}

// Serve : Run document validation server
func Serve(port string) error {
	http.HandleFunc("/", handler)
	fmt.Printf("Running server on port %s...\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
