package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type Task struct {
	ID       int    `json:"id"`
	TaskName string `json:"taskName"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var currentID = 1

func main() {
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	mux := http.NewServeMux()

	// 2. GET VA POST YO'LAKLARI (Tasklarni ko'rish va yangi qo'shish)
	mux.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(tasks)
			return
		}
		
		if r.Method == "POST" {
			var incoming struct {
				TaskName string `json:"taskName"`
			}
			err := json.NewDecoder(r.Body).Decode(&incoming)
			if err != nil || incoming.TaskName == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			
			newTask := Task{
				ID:        currentID,
				TaskName:  incoming.TaskName,
				Completed: false,
			}
			currentID++
			tasks = append(tasks, newTask)
			
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newTask)
			return
		}
	})

	mux.HandleFunc("/api/tasks/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		idStr := r.URL.Path[len("/api/tasks/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Method == "DELETE" {
			for i, t := range tasks {
				if t.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.Method == "PUT" {
			var incoming struct {
				TaskName  string `json:"taskName"`
				Completed *bool  `json:"completed"`
			}
			json.NewDecoder(r.Body).Decode(&incoming)

			for i, t := range tasks {
				if t.ID == id {
					if incoming.TaskName != "" {
						tasks[i].TaskName = incoming.TaskName
					}
					if incoming.Completed != nil {
						tasks[i].Completed = *incoming.Completed
					}
					json.NewEncoder(w).Encode(tasks[i])
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	// Serverni start qilish va CORS middleware'ni ulash
	http.ListenAndServe(":"+port, corsMiddleware(mux))
}
