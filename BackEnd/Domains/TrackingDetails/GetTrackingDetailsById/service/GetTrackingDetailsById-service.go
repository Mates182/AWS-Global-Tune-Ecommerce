// auto-generated with ginshot
package service

import (
	requests "get-tracking-details-by-id-service/data/requests"
	responses "get-tracking-details-by-id-service/data/responses"
)

type GetTrackingDetailsByIdService interface {
	GetTrackingDetailsByIdHandler(request requests.GetTrackingDetailsByIdRequest) (int, responses.GetTrackingDetailsByIdResponse)
}
