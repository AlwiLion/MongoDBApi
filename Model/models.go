package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

type JsonResponseWithArray struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Data    []primitive.M `json:Data`
}
type JsonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
