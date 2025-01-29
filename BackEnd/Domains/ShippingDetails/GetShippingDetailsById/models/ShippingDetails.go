package models

type ShippingDetails struct {
	ID string `json:"ID" bson:"ID"`
	Address Address `json:"Address" bson:"Address"`
	Phone string `json:"Phone" bson:"Phone"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	Email string `json:"Email" bson:"Email"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
}
