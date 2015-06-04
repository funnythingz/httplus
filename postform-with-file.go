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
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(keyName, filename)
	if err != nil {
		log.Println("error writing to buffer")
		return nil, err
	}

	fh, err := os.Open(filename)
	if err != nil {
		log.Println("error opening file")
		return nil, err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}

	contentType := bodyWriter.FormDataContentType()

	for key, val := range valueParams {
		_ = bodyWriter.WriteField(key, val)
	}

	bodyWriter.Close()

	return http.Post(targetUrl, contentType, bodyBuf)
}
