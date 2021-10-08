package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s", dir)

func TestProvider(t *testing.T) {

	pact := &dsl.Pact{
		Consumer: "todo-list-front",
		Provider: "todo-list-backend",
	}
	go startServer()

	f := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:9000",
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/todo-list-front-todo-list-backend.json", pactDir))},
		RequestFilter:   f,
		StateHandlers: types.StateHandlers{
			"there is task": func() error {
				success = "true"
				return nil
			},
		},
	})
}

var success = "true"

func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/taks/add", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, fmt.Sprintf(`{"success":"%v"}}`, success))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
