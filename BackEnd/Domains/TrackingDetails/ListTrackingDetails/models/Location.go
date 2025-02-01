package models

type Location struct {
	ZIPPostalCode string `json:"ZIPPostalCode" bson:"ZIPPostalCode"`
	Country string `json:"Country" bson:"Country"`
	City string `json:"City" bson:"City"`
	State string `json:"State" bson:"State"`
}
