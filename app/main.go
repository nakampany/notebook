package main

import (
	"fmt"
	"log"
	"net/http"
	"note/app/api"
	"note/app/db"
	"os"
	"time"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! The time is %s", time.Now())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := api.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", r))
}
