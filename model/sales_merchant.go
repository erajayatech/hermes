package model

import "time"

type SalesMerchant struct {
	ID                             int64     `json:"id"`
	OrderID                        int64     `json:"order_id"`
	OrderNumberOms                 string    `json:"order_number_oms"`
	MerchantCode                   string    `json:"merchant_code"`
	MerchantName                   string    `json:"merchant_name"`
	MerchantShippingTitle          string    `json:"merchant_shipping_title"`
	MerchantShippingCode           string    `json:"merchant_shipping_code"`
	MerchantShippingAmount         int64     `json:"merchant_shipping_amount"`
	MerchantShippingAmountFormated string    `json:"merchant_shipping_amount_formated"`
	CreatedAt                      time.Time `json:"created_at"`
	UpdatedAt                      time.Time `json:"updated_at"`
}
