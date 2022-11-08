package sdk

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// GetCharacterList gets LOTR character list.
func (l LotrSDK) GetCharacterList(reqOpt RequestOptions) ([]Character, Response) {
	requestURL := MakeRequestURL(fmt.Sprintf("%s/%s", BASE_URL, CHARACTER_URL), reqOpt)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	fmt.Println("RESPO")
	fmt.Println(response)
	if err != nil {
		l.processError(err)

		return []Character{}, response
	}

	var characters []Character

	for _, doc := range response.Docs {
		character := Character{}
		mapstructure.Decode(doc, &character)
		characters = append(characters, character)
	}

	return characters, response
}

// GetCharacterByID gets a single LOTR character by ID.
func (l LotrSDK) GetCharacterByID(id string) (Character, Response) {
	requestURL := fmt.Sprintf("%s/%s/%s", BASE_URL, CHARACTER_URL, id)

	response, err := l.makeHTTPRequest(requestURL, http.MethodGet, nil)
	if err != nil {
		l.processError(err)

		return Character{}, response
	}

	doc := response.Docs[0]
	character := Character{}
	mapstructure.Decode(doc, &character)

	return character, response
}
