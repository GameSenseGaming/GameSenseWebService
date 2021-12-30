package api

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) []byte {
	request, error := http.NewRequest("GET", url, nil)
	if error != nil {
		panic(error)
	}
	request.Header.Add("content-type", "application/json")

	client := &http.Client{}
	reponse, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer reponse.Body.Close()

	data, error := ioutil.ReadAll(reponse.Body)
	if error != nil {
		panic(error)
	}

	return data
}
