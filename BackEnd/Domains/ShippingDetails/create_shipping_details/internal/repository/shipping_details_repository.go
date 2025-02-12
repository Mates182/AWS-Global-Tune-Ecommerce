package dbcontext

import (
	"context"
	"create_shipping_details/internal/data/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShippingDetailsRepository struct {
	DB *mongo.Collection
}

func NewShippingDetailsRepository(db *mongo.Database) *ShippingDetailsRepository {
	return &ShippingDetailsRepository{DB: db.Collection("shipping_details")}
}

func (r *ShippingDetailsRepository) Create(ShippingDetails models.ShippingDetails) error {
	if ShippingDetails.ID != "" {
		idTag := "_id"
		existing := r.DB.FindOne(context.Background(), bson.M{idTag: ShippingDetails.ID})
		if existing.Err() == nil {
			return fmt.Errorf("ShippingDetails with the ID already exists")
		}
	}

	_, err := r.DB.InsertOne(context.Background(), ShippingDetails)
	return err
}

func (r *ShippingDetailsRepository) Read(id string) (*models.ShippingDetails, error) {
	var ShippingDetails models.ShippingDetails
	err := r.DB.FindOne(context.Background(), bson.M{"ID": id}).Decode(&ShippingDetails)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("ShippingDetails not found")
	}
	return &ShippingDetails, err
}

func (r *ShippingDetailsRepository) List() ([]models.ShippingDetails, error) {
	cursor, err := r.DB.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var arr []models.ShippingDetails
	if err := cursor.All(context.Background(), &arr); err != nil {
		return nil, err
	}
	return arr, nil
}

func (r *ShippingDetailsRepository) Update(ShippingDetails models.ShippingDetails) error {
	idTag:= "_id"
	filter := bson.M{idTag: ShippingDetails.ID}
	update := bson.M{"$set": ShippingDetails}

	result, err := r.DB.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("ShippingDetails not found")
	}
	return nil
}

func (r *ShippingDetailsRepository) Delete(id string) error {
	idTag := "_id"
	filter := bson.M{idTag: id}
	result, err := r.DB.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("ShippingDetails not found")
	}
	return nil
}
