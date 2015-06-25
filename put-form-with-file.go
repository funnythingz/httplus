package httplus

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func (h *Http) PutFormWithFile(targetUrl string, valueParams map[string]string, keyName string, filename string) (*http.Response, error) {
	contentType, bodyBuf, err := h.CreateContentTypeAndBody(valueParams, keyName, filename)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", targetUrl, bodyBuf)
	req.Header.Add("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}
