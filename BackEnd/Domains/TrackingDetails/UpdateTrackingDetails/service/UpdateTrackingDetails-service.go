// auto-generated with ginshot
package service

import (
	requests "update-tracking-details-service/data/requests"
	responses "update-tracking-details-service/data/responses"
)

type UpdateTrackingDetailsService interface {
	UpdateTrackingDetailsHandler(request requests.UpdateTrackingDetailsRequest) (int, responses.UpdateTrackingDetailsResponse)
}
