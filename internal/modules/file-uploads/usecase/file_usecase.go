package usecase

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/file-uploads/interfaces"
)

type fileUseCase struct {
	fileRepository interfaces.FileRepository
}

func NewFileUseCase(fileRepository interfaces.FileRepository) *fileUseCase {
	return &fileUseCase{fileRepository}
}
