package response

import "delete-tracking-details-service/models"

type DeleteTrackingDetailsResponse struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
	Message string `json:"Message"`
}
