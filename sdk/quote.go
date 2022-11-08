package sdk

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// GetQuoteList gets LOTR quote list.
func (l LotrSDK) GetQuoteList(reqOpt RequestOptions) ([]Quote, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s", BASE_URL, QUOTE_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	fmt.Println("RESPO")
	fmt.Println(response)
	if err != nil {
		l.processError(err)

		return []Quote{}, response
	}

	var quotes []Quote

	for _, doc := range response.Docs {
		quote := Quote{}
		mapstructure.Decode(doc, &quote)
		quotes = append(quotes, quote)
	}

	return quotes, response
}

// GetQuoteByID gets a single LOTR quote by ID.
func (l LotrSDK) GetQuoteByID(id string) (Quote, Response) {
	requestURL := fmt.Sprintf("%s/%s/%s", BASE_URL, QUOTE_URL, id)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return Quote{}, response
	}

	doc := response.Docs[0]
	quote := Quote{}
	mapstructure.Decode(doc, &quote)

	return quote, response
}

// GetQuoteByMovieID gets LOTR chapters by movie ID.
func (l LotrSDK) GetQuoteByMovieID(id string, reqOpt RequestOptions) ([]Quote, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s/%s/%s", BASE_URL, MOVIE_URL, id, CHAPTER_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return []Quote{}, response
	}

	var quotes []Quote

	for _, doc := range response.Docs {
		quote := Quote{}
		mapstructure.Decode(doc, &quote)
		quotes = append(quotes, quote)
	}

	return quotes, response
}

// GetQuoteByCharacterID gets LOTR chapters by character ID.
func (l LotrSDK) GetQuoteByCharacterID(id string, reqOpt RequestOptions) ([]Quote, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s/%s/%s", BASE_URL, CHARACTER_URL, id, CHAPTER_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return []Quote{}, response
	}

	var quotes []Quote

	for _, doc := range response.Docs {
		quote := Quote{}
		mapstructure.Decode(doc, &quote)
		quotes = append(quotes, quote)
	}

	return quotes, response
}
