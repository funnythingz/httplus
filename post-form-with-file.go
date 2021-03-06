package httplus

import (
	"net/http"
)

func PostFormWithFile(targetUrl string, valueParams map[string]string, keyName string, filename string) (*http.Response, error) {
	contentType, bodyBuf, err := CreateContentTypeAndBody(valueParams, keyName, filename)
	if err != nil {
		return nil, err
	}

	return http.Post(targetUrl, contentType, bodyBuf)
}
