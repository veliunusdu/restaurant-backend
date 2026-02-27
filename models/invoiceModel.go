package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"_id"`
	Invoice_id       *string            `json:"invoice_id" validate:"required"`
	Order_id         *string            `json:"order_id" validate:"required"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CARD|eq=CASH|eq=ONLINE"`
	Payment_status   *string            `json:"payment_status" validate:"required, eq=Pending|eq=Paid|eq=Failed"`
	Payment_due_date *string            `json:"payment_due_date" validate:"required"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}
