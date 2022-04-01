package webTool

import (
	"io"
	"net/http"
	"os"
)

// Get only get information from url
func Get(url string) ([]byte, error) {

	client := &http.Client{}

	token := os.Getenv("GCP_TOKEN")
	bearer := "Bearer " + token

	// Formatting request with some information, you can set header
	// body permit add some information to post request but not necessary to method get
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("authorization", bearer)
	if err != nil {
		return nil, err
	}

	// Executing the request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Parsing the information from memory to byte
	// to ready body you need parse to string
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}