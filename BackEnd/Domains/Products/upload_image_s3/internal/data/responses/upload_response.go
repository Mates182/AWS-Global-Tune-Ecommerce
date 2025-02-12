package responses

type UploadImageResponse struct {
	Message  string `json:"message"`
	ImageURL string `json:"image_url"`
}
