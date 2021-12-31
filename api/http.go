package api

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) []byte {
	request, error := http.NewRequest("GET", url, nil)
	ErrorHandler(error)
	request.Header.Add("content-type", "application/json")

	client := &http.Client{}
	reponse, error := client.Do(request)
	ErrorHandler(error)
	defer reponse.Body.Close()

	data, error := ioutil.ReadAll(reponse.Body)
	ErrorHandler(error)

	return data
}
