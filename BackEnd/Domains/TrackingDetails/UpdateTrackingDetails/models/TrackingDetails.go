package models

type TrackingDetails struct {
	TrackingNumber string `json:"TrackingNumber" bson:"TrackingNumber"`
	OrderID string `json:"OrderID" bson:"OrderID"`
	Carrier string `json:"Carrier" bson:"Carrier"`
	EstimatedDeliveryDate string `json:"EstimatedDeliveryDate" bson:"EstimatedDeliveryDate"`
	Status string `json:"Status" bson:"Status"`
	Origin Location `json:"Origin" bson:"Origin"`
	Destination Location `json:"Destination" bson:"Destination"`
	CurrentLocation Location `json:"CurrentLocation" bson:"CurrentLocation"`
	TrackingHistory []TrackingHistory `json:"TrackingHistory" bson:"TrackingHistory"`
}
