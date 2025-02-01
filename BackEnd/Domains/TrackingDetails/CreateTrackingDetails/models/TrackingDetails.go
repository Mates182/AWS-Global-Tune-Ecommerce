package models

type TrackingDetails struct {
	OrderID string `json:"OrderID" bson:"OrderID"`
	Carrier string `json:"Carrier" bson:"Carrier"`
	CurrentLocation Location `json:"CurrentLocation" bson:"CurrentLocation"`
	TrackingHistory []TrackingHistory `json:"TrackingHistory" bson:"TrackingHistory"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	Status string `json:"Status" bson:"Status"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	Origin Location `json:"Origin" bson:"Origin"`
	Destination Location `json:"Destination" bson:"Destination"`
}
