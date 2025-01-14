package image

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/JesseNicholas00/GogoManager/services/image"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const maxSize int64 = 101 << 10 // 100 KiB with padding

func (ctrl *imageController) uploadFile(c echo.Context) error {
	mpf, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image failed to upload",
		})
	}
	files := mpf.File["file"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is an empty file",
		})
	}
	file := files[0]

	fileParts := strings.Split(file.Filename, ".")
	fileType := strings.ToLower(fileParts[len(fileParts)-1])
	if fileType != "jpg" && fileType != "jpeg" && fileType != "png" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image file type is wrong",
		})
	}
	if file.Size > maxSize {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image file size is wrong",
		})
	}

	filename := uuid.NewString() + "." + fileType
	uri := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", ctrl.bucket, ctrl.region, filename)

	go func(uploader *manager.Uploader, file *multipart.FileHeader, bucket, name string) {
		src, err := file.Open()
		if err != nil {
			return
		}
		defer src.Close()
		params := &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(name),
			Body:   src,
			ACL:    types.ObjectCannedACLPublicRead, // Allowed public read
		}

		_, err = uploader.Upload(context.Background(), params)
		if err != nil {
			return
		}
	}(ctrl.service, file, ctrl.bucket, filename)

	res := image.UploadImageRes{
		Uri: uri,
	}
	return c.JSON(http.StatusOK, res)
}
