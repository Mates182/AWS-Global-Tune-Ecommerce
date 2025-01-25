package model

type Order struct {
	ID    string     `json:"id"`
	Items OrderItems `json:"items"`
}

type OrderItems struct {
	Number            int    `json:"number"`
	Date              string `json:"date"`
	BillingDetailsID  string `json:"billing_details_id"`
	ShippingDetailsID string `json:"shipping_details_id"`
	InvoiceID         string `json:"invoice_id"`
	PaymentID         string `json:"payment_id"`
}
