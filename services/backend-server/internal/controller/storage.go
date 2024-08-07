package controller

import (
	"io"
	"net/url"

	"github.com/google/uuid"
)

func (c *Controller) UploadFile(patientId uuid.UUID, body io.Reader) (string, error) {
	return c.storage.UploadFile(patientId, body)
}

func (c *Controller) GenerateURL(key string) (*url.URL, error) {
	return c.storage.GenerateURL(key)
}

func (c *Controller) DeleteFile(key string) error {
	return c.storage.DeleteFile(key)
}
