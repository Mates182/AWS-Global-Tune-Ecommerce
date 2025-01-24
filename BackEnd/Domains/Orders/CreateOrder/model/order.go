package model

type Order struct {
	ID    string `json:"id"`
	Items []struct {
		Number            int    `json:"number"`
		Date              string `json:"date"`
		BillingDetailsID  string `json:"billing_details_id"`
		ShippingDetailsID string `json:"shipping_details_id"`
		InvoiceID         string `json:"invoice_id"`
		PaymentID         string `json:"payment_id"`
	} `json:"items"`
}
