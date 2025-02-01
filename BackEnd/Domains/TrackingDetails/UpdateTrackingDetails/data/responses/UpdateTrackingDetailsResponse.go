package response

import "update-tracking-details-service/models"

type UpdateTrackingDetailsResponse struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
	Message string `json:"Message"`
}
