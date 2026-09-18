package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"gojobtracker/internal/handlers"
	"gojobtracker/internal/store"
)

func main() {
	st := store.New()
	h := handlers.New(st)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", h.Health)
	mux.HandleFunc("/api/jobs", h.Jobs)
	mux.HandleFunc("/api/jobs/", h.JobByID)
	mux.HandleFunc("/api/stats", h.Stats)
	mux.Handle("/", http.FileServer(http.Dir("web")))

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }

	server := &http.Server{
		Addr: ":" + port,
		Handler: logging(cors(mux)),
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Job Tracker running at http://localhost:%s", port)
	log.Fatal(server.ListenAndServe())
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
