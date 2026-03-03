package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        *primitive.ObjectID `bson:"_id"`
	Note_id   *string             `json:"note_id" validate:"required"`
	Title     *time.Time          `json:"order_date" validate:"required"`
	Text      *string             `json:"text" validate:"required"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}
