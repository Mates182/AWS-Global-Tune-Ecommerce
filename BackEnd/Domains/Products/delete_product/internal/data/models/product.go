package models

type Product struct {
	Title string `json:"Title" bson:"Title"`
	Stock int `json:"Stock" bson:"Stock"`
	Active bool `json:"Active" bson:"Active"`
	Tags []string `json:"Tags" bson:"Tags"`
	Weight float32 `json:"Weight" bson:"Weight"`
	Description string `json:"Description" bson:"Description"`
	Price float32 `json:"Price" bson:"Price"`
	Category string `json:"Category" bson:"Category"`
	Brand string `json:"Brand" bson:"Brand"`
	Sku string `json:"Sku" bson:"Sku"`
	Warranty string `json:"Warranty" bson:"Warranty"`
	ID string `json:"_id" bson:"_id"`
	Images []string `json:"Images" bson:"Images"`
	Thumbnail string `json:"Thumbnail" bson:"Thumbnail"`
}
