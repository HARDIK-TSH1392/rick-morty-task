// handlers/character.go
package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "rick_and_morty_api/models"
    "rick_and_morty_api/utils"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/go-resty/resty/v2"
    "github.com/gorilla/mux"
    "log"
)

var client = resty.New()

func SearchCharacter(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    collection := utils.Client.Database("rick_and_morty_db").Collection("characters")

    // Check if the character exists in MongoDB
    var character models.Character
    err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&character)
    if err == mongo.ErrNoDocuments {
        // If not found in MongoDB, fetch from the Rick and Morty API
        response, err := client.R().Get("https://rickandmortyapi.com/api/character/?name=" + name)
        if err != nil || response.StatusCode() != 200 {
            http.Error(w, "Character not found", http.StatusNotFound)
            return
        }

        // Parse API response
        var characterData map[string]interface{}
        json.Unmarshal(response.Body(), &characterData)

        // Extract the first character from the "results" array
        results := characterData["results"].([]interface{})
        if len(results) > 0 {
            characterMap := results[0].(map[string]interface{})
            character = models.Character{
                Name:     characterMap["name"].(string),
                Gender:   characterMap["gender"].(string),
                Species:  characterMap["species"].(string),
                Status:   characterMap["status"].(string),
                Type:     characterMap["type"].(string),
                Episodes: convertEpisodes(characterMap["episode"].([]interface{})),
            }

            // Save the character to MongoDB
            _, err := collection.InsertOne(context.TODO(), character)
            if err != nil {
                log.Println("Failed to save character:", err)
                http.Error(w, "Failed to save character", http.StatusInternalServerError)
                return
            }
            
            // Log success message
            log.Printf("Character '%s' successfully saved to MongoDB.\n", character.Name)
        }
    } else if err != nil {
        log.Println("Database error:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Return character data as response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(character)
}

// Utility function to convert episode URLs from API response to a slice of strings
func convertEpisodes(episodes []interface{}) []string {
    var episodeList []string
    for _, episode := range episodes {
        episodeList = append(episodeList, episode.(string))
    }
    return episodeList
}

func GetEpisodesByCharacterName(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]

    var character models.Character
    collection := utils.Client.Database("rick_and_morty_db").Collection("characters")

    // Check if the character exists in MongoDB
    err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&character)
    if err == mongo.ErrNoDocuments {
        http.Error(w, "Character not found in database", http.StatusNotFound)
        return
    } else if err != nil {
        log.Println("Database error:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Return only the episode numbers as response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(character.Episodes)
}

// FetchCharacterFromDatabase retrieves a character from MongoDB by name
func FetchCharacterFromDatabase(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    collection := utils.Client.Database("rick_and_morty_db").Collection("characters")

    var character models.Character
    err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&character)
    if err == mongo.ErrNoDocuments {
        http.Error(w, "Character not found in database", http.StatusNotFound)
        return
    } else if err != nil {
        log.Println("Database error:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(character)
}
