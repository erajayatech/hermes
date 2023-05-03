package model

type GiftcardDetail struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Source    string `json:"source"`
	Card      []Card `json:"cards"`
}

type Card struct {
	CardNumber  string `json:"card_number"`
	CardPin     string `json:"card_pin"`
	Denom       string `json:"denom"`
	ExpiryDate  string `json:"expiry_date"`
	OrderNumber string `json:"order_number"`
}
