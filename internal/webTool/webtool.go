package webTool

import (
	"github.com/tonnytg/webreq"
	_ "github.com/tonnytg/webreq"
	"os"
)

// Get only get information from url
func Get(url string) ([]byte, error) {

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + os.Getenv("GCP_API_KEY"),
	}
	body, err := webreq.Get(url, headers, 5)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Post get json format in slice of byte from data and use in bufferData to send url
func PostJson(url string, jsonDataByte []byte) ([]byte, error) {

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + os.Getenv("GCP_API_KEY"),
	}
	body, err := webreq.Post(url, headers, 5, jsonDataByte)
	if err != nil {
		return nil, err
	}

	return body, nil
}
