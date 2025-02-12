package models

type Address struct {
	CountryRegion string `json:"CountryRegion" bson:"CountryRegion"`
	Street string `json:"Street" bson:"Street"`
	City string `json:"City" bson:"City"`
	StateProvince string `json:"StateProvince" bson:"StateProvince"`
	ZIPPostalCode string `json:"ZIPPostalCode" bson:"ZIPPostalCode"`
}
