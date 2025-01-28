package models

type TrackingHistory struct {
	Status string `json:"Status" bson:"Status"`
	Location Location `json:"Location" bson:"Location"`
	Timestamp string `json:"Timestamp" bson:"Timestamp"`
}
