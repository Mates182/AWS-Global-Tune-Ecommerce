package models

type ShippingDetails struct {
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	Phone string `json:"Phone" bson:"Phone"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
	ID string `json:"_id" bson:"_id"`
	Address Address `json:"Address" bson:"Address"`
	Email string `json:"Email" bson:"Email"`
}
