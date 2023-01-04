package dadata

import (
	"github.com/Unbel1evab7e/bank-interface/domain/properties"
	"github.com/Unbel1evab7e/guu"
)

type DaDataClient struct {
	Properties properties.DaDataProperties
}

func New(props *properties.DaDataProperties) *DaDataClient {
	return &DaDataClient{Properties: *props}
}

func (c DaDataClient) GetSuggestions(query string) (*SuggestionResponse, error) {
	headers := make(map[string]string)

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = "Token " + c.Properties.Key
	headers["X-Secret"] = c.Properties.Secret

	response, err := guu.ExecutePost[SuggestionResponse](c.Properties.SuggestionsUrl, SuggestionRequest{Query: query}, nil, headers)

	if err != nil {
		return nil, err
	}

	return response, nil
}
