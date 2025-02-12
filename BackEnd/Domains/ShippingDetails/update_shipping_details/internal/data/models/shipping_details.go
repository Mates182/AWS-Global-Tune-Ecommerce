package models

type ShippingDetails struct {
	Phone string `json:"Phone" bson:"Phone"`
	ShippingMethod string `json:"ShippingMethod" bson:"ShippingMethod"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	ID string `json:"_id" bson:"_id"`
	RecipientName string `json:"RecipientName" bson:"RecipientName"`
	Address Address `json:"Address" bson:"Address"`
	Email string `json:"Email" bson:"Email"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	SpecialInstructions string `json:"SpecialInstructions" bson:"SpecialInstructions"`
}
