package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"


//creating a client for all http request that we will make
type Client struct {
	httpClient http.Client
}

//func to intialize the client with timeout of a minute so that 
//our program should not weight for request to complete 
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
