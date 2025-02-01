package models

type TrackingDetails struct {
	CurrentLocation Location `json:"CurrentLocation" bson:"CurrentLocation"`
	Destination Location `json:"Destination" bson:"Destination"`
	OrderID string `json:"OrderID" bson:"OrderID"`
	Carrier string `json:"Carrier" bson:"Carrier"`
	Status string `json:"Status" bson:"Status"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	Origin Location `json:"Origin" bson:"Origin"`
	TrackingHistory []TrackingHistory `json:"TrackingHistory" bson:"TrackingHistory"`
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
}
