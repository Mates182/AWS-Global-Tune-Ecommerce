package models

type ShippingDetails struct {
	ID string `json:"ID" bson:"ID"`
	Phone string `json:"Phone" bson:"Phone"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	Address Address `json:"Address" bson:"Address"`
	Email string `json:"Email" bson:"Email"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
}
