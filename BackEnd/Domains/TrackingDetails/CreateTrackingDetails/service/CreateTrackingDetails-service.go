// auto-generated with ginshot
package service

import (
	requests "create-tracking-details-service/data/requests"
	responses "create-tracking-details-service/data/responses"
)

type CreateTrackingDetailsService interface {
	CreateTrackingDetailsHandler(request requests.CreateTrackingDetailsRequest) (int, responses.CreateTrackingDetailsResponse)
}
