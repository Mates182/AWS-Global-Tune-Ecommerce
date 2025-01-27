package models

type Location struct {
	Country string `json:"Country" bson:"Country"`
	City string `json:"City" bson:"City"`
	State string `json:"State" bson:"State"`
	ZIPPostalCode string `json:"ZIPPostalCode" bson:"ZIPPostalCode"`
}
