package request

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func init() {
}

func mapToJSON(inputMap map[string]interface{}) string {
	jsonData, _ := json.Marshal(inputMap)
	return string(jsonData)
}

// Sent formdata type requesr, use basic auth
// url: request url
// method: request method
// formData: formdata map
// return: response body
func POST(url string, BodyData map[string]interface{}) (responseData map[string]interface{}) {

	response, err := http.Post(url, "application/json", bytes.NewBufferString(mapToJSON(BodyData)))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Sent get type request, use basic auth
// url: request url
// queryData: query data map
func GET(url string) (responseData map[string]interface{}) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func PUT(url string) (responseData map[string]interface{}) {
	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		log.Fatal(err)
	}
	return
}
