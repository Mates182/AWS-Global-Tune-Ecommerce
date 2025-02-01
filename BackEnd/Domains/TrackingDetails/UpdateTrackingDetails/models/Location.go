package models

type Location struct {
	State string `json:"State" bson:"State"`
	ZIPPostalCode string `json:"ZIPPostalCode" bson:"ZIPPostalCode"`
	Country string `json:"Country" bson:"Country"`
	City string `json:"City" bson:"City"`
}
