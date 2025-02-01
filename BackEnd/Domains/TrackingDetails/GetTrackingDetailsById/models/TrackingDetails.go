package models

type TrackingDetails struct {
	CurrentLocation Location `json:"CurrentLocation" bson:"CurrentLocation"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	OrderID string `json:"OrderID" bson:"OrderID"`
	Carrier string `json:"Carrier" bson:"Carrier"`
	Origin Location `json:"Origin" bson:"Origin"`
	Status string `json:"Status" bson:"Status"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	Destination Location `json:"Destination" bson:"Destination"`
	TrackingHistory []TrackingHistory `json:"TrackingHistory" bson:"TrackingHistory"`
}
