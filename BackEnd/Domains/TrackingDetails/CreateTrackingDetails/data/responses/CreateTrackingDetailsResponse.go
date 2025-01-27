package response

import "create-tracking-details-service/models"

type CreateTrackingDetailsResponse struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
	Message string `json:"Message"`
}
