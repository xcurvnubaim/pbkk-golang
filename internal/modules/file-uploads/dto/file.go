package dto

import "mime/multipart"

type UploadFileRequest struct {
	FileName string                `form:"file_name" binding:"required"`
	File     *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadFileResponse struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}
