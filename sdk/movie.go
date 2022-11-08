package sdk

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// GetMovieList gets LOTR movie list.
func (l LotrSDK) GetMovieList(reqOpt RequestOptions) ([]Movie, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s", BASE_URL, MOVIE_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return []Movie{}, response
	}

	var movies []Movie

	for _, doc := range response.Docs {
		movie := Movie{}
		mapstructure.Decode(doc, &movie)
		movies = append(movies, movie)
	}

	return movies, response
}

// GetMovieByID gets a single LOTR movie by ID.
func (l LotrSDK) GetMovieByID(id string) (Movie, Response) {
	requestURL := fmt.Sprintf("%s/%s/%s", BASE_URL, MOVIE_URL, id)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return Movie{}, response
	}

	doc := response.Docs[0]
	movie := Movie{}
	mapstructure.Decode(doc, &movie)

	return movie, response
}
