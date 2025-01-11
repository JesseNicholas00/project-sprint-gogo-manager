package image

import (
	"net/http"
	"strings"

	"github.com/JesseNicholas00/GogoManager/services/image"
	"github.com/JesseNicholas00/GogoManager/utils/errorutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const maxSize int64 = 20 << 20
const minSize int64 = 10 << 10

func (ctrl *imageController) uploadFile(c echo.Context) error {
	mpf, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is wrong",
		})
	}
	files := mpf.File["file"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is wrong",
		})
	}
	file := files[0]

	fileParts := strings.Split(file.Filename, ".")
	fileType := fileParts[len(fileParts)-1]
	if fileType != "jpg" && fileType != "jpeg" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is wrong",
		})
	}
	if file.Size < minSize || file.Size > maxSize {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is wrong",
		})
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	params := &s3.PutObjectInput{
		Bucket: aws.String(ctrl.bucket),
		Key:    aws.String(uuid.NewString() + "." + fileType),
		Body:   src,
	}

	result, err := ctrl.service.Upload(c.Request().Context(), params)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res := image.UploadImageRes{
		Uri: result.Location,
	}
	return c.JSON(http.StatusOK, res)
}
