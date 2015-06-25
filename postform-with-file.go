package httplus

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func (h *Http) PostFormWithFile(targetUrl string, valueParams map[string]string, keyName string, filename string) (*http.Response, error) {
	contentType, bodyBuf, err := h.CreateContentTypeAndBody(valueParams, keyName, filename)
	if err != nil {
		return nil, err
	}

	return http.Post(targetUrl, contentType, bodyBuf)
}
