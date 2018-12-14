package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/appno/stripe/schema"
)

func handler(w http.ResponseWriter, req *http.Request) {
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	document, err := schema.DocumentFromBytes(bytes)
	if err != nil {
		panic(err)
	}

	compliance := document.GetPastDueCompliance()
	fmt.Println(compliance.DebugString())

	data, err := json.Marshal(compliance)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = io.WriteString(w, string(data))

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
