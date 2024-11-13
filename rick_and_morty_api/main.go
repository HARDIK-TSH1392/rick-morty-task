package main

import (
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "github.com/rs/cors"             // Import the CORS package
    "rick_and_morty_api/handlers"
    "rick_and_morty_api/utils"
)

func main() {
    // mongoURI := "mongodb://localhost:27017"
	mongoURI := "mongodb://34.30.165.222:27017"
    utils.ConnectMongoDB(mongoURI)

    r := mux.NewRouter()
    r.HandleFunc("/search", handlers.SearchCharacter).Methods("GET")
    r.HandleFunc("/episodes/{name}", handlers.GetEpisodesByCharacterName).Methods("GET")

	r.HandleFunc("/database/character", handlers.FetchCharacterFromDatabase).Methods("GET")


    // Add CORS middleware configuration
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Frontend URL
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
        AllowCredentials: true,
    })

    // Wrap the router with the CORS middleware
    handler := c.Handler(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Println("Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))  // Use the CORS handler
}
