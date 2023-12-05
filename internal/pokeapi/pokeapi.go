package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = ""

type Client struct {
	httpClient http.Client
}

func newClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func pokeapi(url string) (LocationAreaResponse, error) {
	resp, err := http.Get(url)
	var response LocationAreaResponse

	if err != nil {
		//fmt.Println("Error:", err)
		return response, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Error reading response body:", err)
		return response, err
	}
	// Unmarshal the JSON data into the LocationAreaResponse struct
	errMarshal := json.Unmarshal(body, &response)

	if errMarshal != nil {
		fmt.Println("Error:", err)
		return response, errMarshal
	}

	// Accessing the fields of the struct
	// fmt.Println("Count:", response.Count)
	// fmt.Println("Next:", response.Next)
	// fmt.Println("Previous:", response.Previous)

	// Accessing the results
	// for _, result := range response.Results {
	// 	fmt.Printf("Name: %s, URL: %s\n", result.Name, result.URL)
	// }

	return response, nil

}