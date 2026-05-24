package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Status struct {
	Env     string `json:"env"`
	Status  string `json:"status"`
	Version string `json:"version"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Status{
			Env:     "production",
			Status:  "running",
			Version: "1.0.0",
		})
	})

	http.HandleFunc("/api/v1/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		tasks := []Task{
			{ID: "1", Title: "Deploy High-Availability Infrastructure", Completed: true},
			{ID: "2", Title: "Configure CI/CD GitOps Pipelines", Completed: true},
		}
		json.NewEncoder(w).Encode(tasks)
	})

	fmt.Printf("MuchToDo Backend Server running smoothly on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}