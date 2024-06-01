package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive" // Import the primitive package for ObjectID
)

type User struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`        // User ID
	Name   string             `json:"name" bson:"name"`     // User name
	Gender string             `json:"gender" bson:"gender"` // User gender
	Age    int                `json:"age" bson:"age"`       // User age
}
