package models

type ShippingDetails struct {
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	Address Address `json:"Address" bson:"Address"`
	Email string `json:"Email" bson:"Email"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
	ID string `json:"_id" bson:"_id"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	Phone string `json:"Phone" bson:"Phone"`
}
