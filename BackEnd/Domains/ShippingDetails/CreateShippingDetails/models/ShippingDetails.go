package models

type ShippingDetails struct {
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
	ID string `json:"ID" bson:"ID"`
	Address Address `json:"Address" bson:"Address"`
	Phone string `json:"Phone" bson:"Phone"`
	Email string `json:"Email" bson:"Email"`
}
