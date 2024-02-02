package model

import "time"

type SalesAttribute struct {
	ID             int64     `json:"id"`
	OrderID        int64     `json:"order_id"`
	AttributeName  string    `json:"attribute_name"`
	AttributeValue string    `json:"attribute_value"`
	CreatedAt      time.Time `json:"created_at"`
}
