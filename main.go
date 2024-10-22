package main

import (
    "log"
    "net/http"
    "purecss-go/api" // Make sure the import path matches your actual project structure
)

func main() {
    // Serve the web frontend from the "web" directory
    http.Handle("/", http.FileServer(http.Dir("./web")))

    // Serve the API routes
    http.HandleFunc("/api/contact", api.ContactRoutes) // Correct function for /api/contact route

    // Start the server
    log.Println("Server starting on port 8888...")
    log.Fatal(http.ListenAndServe(":8888", nil))
}
