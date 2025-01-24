package request

type Request struct {
	Data struct {
		Date              string `json:"date"`
		BillingDetailsID  string `json:"billing_details_id"`
		ShippingDetailsID string `json:"shipping_details_id"`
		InvoiceID         string `json:"invoice_id"`
		PaymentID         string `json:"payment_id"`
	} `json:"data"`
}
