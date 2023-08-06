package main

import (
  "os"
	"fmt"
  "log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Print("> Request received")
		fmt.Fprintf(w, "Hello, World!")
	})

  port, ok := os.LookupEnv("PORT")
  if !ok {
    log.Println("PORT environment variable not set")
    os.Exit(1)
  }

  addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening on %s...\n", addr)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
