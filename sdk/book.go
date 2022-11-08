package sdk

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// GetBookList gets LOTR book list.
func (l LotrSDK) GetBookList(reqOpt RequestOptions) ([]Book, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s", BASE_URL, BOOK_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return []Book{}, response
	}

	var books []Book

	for _, doc := range response.Docs {
		book := Book{}
		mapstructure.Decode(doc, &book)
		books = append(books, book)
	}

	return books, response
}

// GetBookByID gets a single LOTR book by ID.
func (l LotrSDK) GetBookByID(id string) (Book, Response) {
	requestURL := fmt.Sprintf("%s/%s/%s", BASE_URL, BOOK_URL, id)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return Book{}, response
	}

	doc := response.Docs[0]
	book := Book{}
	mapstructure.Decode(doc, &book)

	return book, response
}
