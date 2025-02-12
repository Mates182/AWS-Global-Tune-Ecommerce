// auto-generated with ginshot
package service

import (
	requests "create_product/internal/data/requests"
	responses "create_product/internal/data/responses"
)

type CreateProductService interface {
	CreateProductHandler(request requests.CreateProductRequest) (int, responses.CreateProductResponse)
}
