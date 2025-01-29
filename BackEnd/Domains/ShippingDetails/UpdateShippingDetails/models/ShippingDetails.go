package models

type ShippingDetails struct {
	Phone string `json:"Phone" bson:"Phone"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
	ID string `json:"ID" bson:"ID"`
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	Address Address `json:"Address" bson:"Address"`
	Email string `json:"Email" bson:"Email"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
}
