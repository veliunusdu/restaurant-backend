package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            *primitive.ObjectID `bson:"_id"`
	First_Name    *string             `json:"first_name" validate:"required"`
	Last_Name     *string             `json:"last_name" validate:"required"`
	Password      *string             `json:"phone" validate:"required"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	User_id       *string             `json:"user_id" validate:"required"`
	Email         *string             `json:"email" validate:"required"`
	Avatar        *string             `json:"avatar" validate:"required"`
	Phone         *string             `json:"phone" validate:"required"`
	Token         *string             `json:"token"`
	Refresh_Token *string             `json:"refresh_token"`
}
