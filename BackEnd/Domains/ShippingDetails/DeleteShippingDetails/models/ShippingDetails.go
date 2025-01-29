package models

type ShippingDetails struct {
	ID string `json:"ID" bson:"ID"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	Address Address `json:"Address" bson:"Address"`
	Phone string `json:"Phone" bson:"Phone"`
	Email string `json:"Email" bson:"Email"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
}
