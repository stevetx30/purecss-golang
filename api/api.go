package api

import (
    "encoding/json"
    "net/http"
    "fmt"
)

// ContactForm represents the structure of the form data
type ContactForm struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:"message"`
}

// Message represents the structure of the API response
type Message struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// ContactRoutes handles the /api/contact requests and responses
func ContactRoutes(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        handlePostRequest(w, r)
    } else if r.Method == http.MethodGet {
        handleGetRequest(w, r)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// handleGetRequest handles GET requests to /api/contact
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
    response := Message{Status: "success", Message: "Hello from the Contact API"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// handlePostRequest handles POST requests to /api/contact for form submissions
func handlePostRequest(w http.ResponseWriter, r *http.Request) {
    // Decode the JSON body into the ContactForm struct
    var form ContactForm
    err := json.NewDecoder(r.Body).Decode(&form)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Validate the form fields
    if form.Name == "" || form.Email == "" || form.Message == "" {
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        return
    }

    // Log the form submission for demonstration
    fmt.Printf("Received contact form submission: Name=%s, Email=%s, Message=%s\n", form.Name, form.Email, form.Message)

    // Send a success response back to the client
    response := Message{Status: "success", Message: "Form submitted successfully"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
