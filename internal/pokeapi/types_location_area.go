package pokeapi


//response struct of the location area api
//JSON lint(https://jsonlint.com/) is a useful tool for debugging JSON, it makes it easier to read.
//JSON to Go(https://mholt.github.io/json-to-go/) a useful tool for converting JSON to Go structs. 
//You can use it to generate the structs you'll need to parse the PokeAPI response. 
//Keep in mind it sometimes can't know the exact type of a field that you want, because there are multiple valid options.
//For nullable strings, use *string
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
