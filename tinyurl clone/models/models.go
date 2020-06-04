package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entry struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Hash string             `bson:"hash,omitempty"`
	URL  string             `bson:"url,omitempty"`
}
