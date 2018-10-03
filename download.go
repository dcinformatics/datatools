package datatools

import (
	"net/http"
)

// GetPage loads the provided url and returns the response.
func GetPage(url string) *http.Response {
	resp, err := http.Get(url)
	Check(err)

	return resp
}
