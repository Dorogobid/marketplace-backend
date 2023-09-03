package svc

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *Service) UploadFile(file *multipart.FileHeader, baseURL string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	defer src.Close()

	filePath := filepath.Join("static", fmt.Sprintf("%s%s", uuid.NewString(), filepath.Ext(file.Filename)))
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("os create: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("copy file: %w", err)
	}

	fURL, err := url.JoinPath(baseURL, filePath)
	if err != nil {
		return "", fmt.Errorf("join url: %w", err)
	}

	return fURL, nil
}
