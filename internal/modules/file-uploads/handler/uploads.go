package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/common"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/file-uploads/dto"
)

func (h *FileHandler) Upload(c *gin.Context) {
	var fileUpload dto.UploadFileRequest
	if err := c.ShouldBind(&fileUpload); err != nil {
		errMsg := err.Error()
		c.JSON(400, &common.Response[dto.UploadFileResponse]{
			Status:  false,
			Message: "Bad Request",
			Data:    nil,
			Error:   &errMsg,
		})
		return
	}

	c.JSON(200, &common.Response[dto.UploadFileResponse]{
		Status:  true,
		Message: "Upload File Success",
		Data: &dto.UploadFileResponse{
			FileName: "test",
			FilePath: "test",
		},
		Error: nil,
	},
	)
	return
}
