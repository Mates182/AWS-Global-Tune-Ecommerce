package dbcontext

import (
	"context"
	"update_product/internal/data/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	DB *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{DB: db.Collection("products")}
}

func (r *ProductRepository) Create(Product models.Product) error {
	if Product.ID != "" {
		idTag := "_id"
		existing := r.DB.FindOne(context.Background(), bson.M{idTag: Product.ID})
		if existing.Err() == nil {
			return fmt.Errorf("Product with the ID already exists")
		}
	}

	_, err := r.DB.InsertOne(context.Background(), Product)
	return err
}

func (r *ProductRepository) Read(id string) (*models.Product, error) {
	var Product models.Product
	err := r.DB.FindOne(context.Background(), bson.M{"ID": id}).Decode(&Product)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("Product not found")
	}
	return &Product, err
}

func (r *ProductRepository) List() ([]models.Product, error) {
	cursor, err := r.DB.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var arr []models.Product
	if err := cursor.All(context.Background(), &arr); err != nil {
		return nil, err
	}
	return arr, nil
}

func (r *ProductRepository) Update(Product models.Product) error {
	idTag:= "_id"
	filter := bson.M{idTag: Product.ID}
	update := bson.M{"$set": Product}

	result, err := r.DB.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("Product not found")
	}
	return nil
}

func (r *ProductRepository) Delete(id string) error {
	idTag := "_id"
	filter := bson.M{idTag: id}
	result, err := r.DB.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("Product not found")
	}
	return nil
}
