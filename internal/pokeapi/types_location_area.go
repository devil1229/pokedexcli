package pokeapi

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

// LocationAreaResult represents the structure of each result item in the response.
type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
