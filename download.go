package utils

import (
	"bytes"
	"io"
	"net/http"
)

func DownloadImageAsByte(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
