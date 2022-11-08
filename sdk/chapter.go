package sdk

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// GetChapterList gets LOTR chapter list.
func (l LotrSDK) GetChapterList(reqOpt RequestOptions) ([]Chapter, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s", BASE_URL, CHAPTER_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return []Chapter{}, response
	}

	var chapters []Chapter

	for _, doc := range response.Docs {
		chapter := Chapter{}
		mapstructure.Decode(doc, &chapter)
		chapters = append(chapters, chapter)
	}

	return chapters, response
}

// GetChapterByID gets a single LOTR chapter by ID.
func (l LotrSDK) GetChapterByID(id string) (Chapter, Response) {
	requestURL := fmt.Sprintf("%s/%s/%s", BASE_URL, CHAPTER_URL, id)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return Chapter{}, response
	}

	doc := response.Docs[0]
	chapter := Chapter{}
	mapstructure.Decode(doc, &chapter)

	return chapter, response
}

// GetChaptersByBookID gets LOTR chapters by book ID.
func (l LotrSDK) GetChaptersByBookID(id string, reqOpt RequestOptions) ([]Chapter, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s/%s/%s", BASE_URL, BOOK_URL, id, CHAPTER_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return []Chapter{}, response
	}

	var chapters []Chapter

	for _, doc := range response.Docs {
		chapter := Chapter{}
		mapstructure.Decode(doc, &chapter)
		chapters = append(chapters, chapter)
	}

	return chapters, response
}
