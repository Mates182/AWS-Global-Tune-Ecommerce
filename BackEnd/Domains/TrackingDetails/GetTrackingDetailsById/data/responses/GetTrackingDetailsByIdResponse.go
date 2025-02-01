package response

import "get-tracking-details-by-id-service/models"

type GetTrackingDetailsByIdResponse struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
	Message string `json:"Message"`
}
