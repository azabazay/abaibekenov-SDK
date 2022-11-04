package sdk

import (
	"fmt"
	"net/http"
)

// GetList gets LOTR book list.
func (l lotrSDK) GetBookList() ([]Book, Response) {
	requestURL := fmt.Sprintf("%s/%s", BASE_URL, BOOK_URL)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return []Book{}, response
	}

	var books []Book

	for _, doc := range response.Docs {
		books = append(books, Book{
			ID:   doc["_id"].(string),
			Name: doc["name"].(string),
		})
	}

	return books, response
}

// GetBookByID gets a single LOTR book by ID.
func (l lotrSDK) GetBookByID(id string) (Book, Response) {
	requestURL := fmt.Sprintf("%s/%s/%s", BASE_URL, BOOK_URL, id)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return Book{}, response
	}

	doc := response.Docs[0]
	book := Book{
		ID:   doc["_id"].(string),
		Name: doc["name"].(string),
	}

	return book, response
}
