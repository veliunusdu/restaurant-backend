package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            *primitive.ObjectID `bson:"_id"`
	Quantity      *int                `json:"quantity" validate:"required"`
	Unit_Price    *float64            `json:"unit_price" validate:"required"`
	Food_id       *string             `json:"food_id" validate:"required"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	Order_item_id *string             `json:"order_id" validate:"required"`
}
