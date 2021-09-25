package tools

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Return Body from response
func GetBodyFromUrl(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("%v", err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("%v", err)
	}

	return body
}
