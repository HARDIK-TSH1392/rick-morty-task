// models/character.go
package models

type Character struct {
    ID       int      `json:"id" bson:"id"`
    Name     string   `json:"name" bson:"name"`
    Gender   string   `json:"gender" bson:"gender"`
    Species  string   `json:"species" bson:"species"`
    Status   string   `json:"status" bson:"status"`
    Type     string   `json:"type" bson:"type"`
    Episodes []string `json:"episodes" bson:"episodes"` // Store episode URLs as strings
}
