// auto-generated with ginshot
package service

import (
	requests "create_shipping_details/internal/data/requests"
	responses "create_shipping_details/internal/data/responses"
)

type CreateShippingDetailsService interface {
	CreateShippingDetailsHandler(request requests.CreateShippingDetailsRequest) (int, responses.CreateShippingDetailsResponse)
}
