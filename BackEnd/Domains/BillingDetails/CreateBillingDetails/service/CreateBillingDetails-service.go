// auto-generated with ginshot
package service

import (
	requests "create-billing-details-service/data/requests"
	responses "create-billing-details-service/data/responses"
)

type CreateBillingDetailsService interface {
	CreateBillingDetailsHandler(request requests.Request) (int, responses.Response)
}
