// auto-generated with ginshot
package service

import (
	requests "create-shipping-details/data/requests"
	responses "create-shipping-details/data/responses"
)

type CreateShippingDetailsService interface {
	CreateShippingDetailsHandler(request requests.CreateShippingDetailsRequest) (int, responses.CreateShippingDetailsResponse)
}
