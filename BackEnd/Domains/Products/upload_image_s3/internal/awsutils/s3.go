package awsutils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"upload_image_s3/internal/data/requests"
	"upload_image_s3/internal/data/responses"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

// uploadImageToS3 uploads the image to S3 and returns the image URL
func uploadImageToS3(file multipart.File, filename string) (string, error) {
	// Initialize the AWS session (uses EC2 IAM role for credentials)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your region
	})
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}

	// Initialize the S3 service client
	svc := s3.New(sess)

	// Get the file type (you can use a library like mime to get the correct content type)
	contentType, err := getContentType(file)
	if err != nil {
		return "", fmt.Errorf("failed to determine content type: %w", err)
	}

	// Reset file pointer to start (after reading for content type)
	_, err = file.Seek(0, 0)
	if err != nil {
		return "", fmt.Errorf("failed to reset file pointer: %w", err)
	}

	// Prepare the upload parameters (no ACL parameter, since ACLs are not allowed)
	params := &s3.PutObjectInput{
		Bucket:      aws.String("globaltune-products-images-qa"), // Replace with your bucket name
		Key:         aws.String(filename),                        // Filename in S3
		Body:        file,                                        // The file to upload
		ContentType: aws.String(contentType),                     // Set the correct content type
	}

	// Upload the file to S3
	_, err = svc.PutObject(params)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// Construct the URL for the uploaded image
	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "globaltune-products-images-qa", filename)

	return imageURL, nil
}

// uploadImageHandler handles the image upload and returns the URL
func UploadImageHandler(c *gin.Context) {
	// Bind the form data (including file)
	var req requests.UploadImageRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Open the uploaded file
	file, err := req.Image.Open()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error opening file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer file.Close()

	// Extract the file extension from the uploaded file
	ext := filepath.Ext(req.Image.Filename) // Get the file extension

	// Construct the full filename using the custom filename provided in the request
	finalFilename := req.Filename + ext // Combine custom filename with the original extension

	// Upload to S3 and get the image URL
	imageURL, err := uploadImageToS3(file, finalFilename)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error upload file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload to S3", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.UploadImageResponse{
		Message:  "File uploaded successfully to S3",
		ImageURL: imageURL,
	})
}

// Function to determine Content-Type dynamically
func getContentType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Detect the content type from the file content
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
