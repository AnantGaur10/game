package utils 

import (
	"log"
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v any,status int) {
    

    // Set the content type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Write the status code
    w.WriteHeader(status)

    // Try encoding the value to JSON, and handle potential errors
    if err := json.NewEncoder(w).Encode(v); err != nil {
        // Log the error if necessary
        log.Printf("Failed to encode JSON: %v", err)

        // Respond with a 500 Internal Server Error and a generic message
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}
