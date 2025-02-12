package requests

import "mime/multipart"

type UploadImageRequest struct {
	Title       string                `json:"title" bson:"title"`
	Description string                `json:"description" bson:"description"`
	Image       *multipart.FileHeader `form:"image" json:"image" bson:"image"`
	Filename    string                `form:"filename" json:"filename"`
}
