package request

import "create_product/internal/data/models"

type CreateProductRequest struct {
	Product models.Product `json:"Product"`
}
