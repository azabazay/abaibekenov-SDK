package main

import (
	"fmt"

	"github.com/abaibekenov-SDK/sdk"
)

func main() {
	lotrSDK := sdk.InitFromConfig(sdk.DefaultConfigFileName)

	// exampleBookRequest(lotrSDK)
	// exampleChapterRequest(lotrSDK)
	exampleMovieRequest(lotrSDK)
	// exampleCharacterRequest(lotrSDK)
	// exampleQuoteRequest(lotrSDK)
}

func exampleBookRequest(lotrSDK sdk.LotrSDK) {
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
	}

	books, _ := lotrSDK.GetBookList(requestOptions)
	fmt.Printf("Books: %v\n", books)

	bookByID, _ := lotrSDK.GetBookByID(books[0].ID)
	fmt.Printf("Book by ID: %v\n", bookByID)
}

func exampleChapterRequest(lotrSDK sdk.LotrSDK) {
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
	}

	chapters, _ := lotrSDK.GetChapterList(requestOptions)
	fmt.Printf("Chapters: %v\n", chapters)

	chapterByID, _ := lotrSDK.GetChapterByID(chapters[0].ID)
	fmt.Printf("Chapter by ID: %v\n", chapterByID)

	books, _ := lotrSDK.GetBookList(requestOptions)
	chaptersByBookID, _ := lotrSDK.GetChaptersByBookID(books[0].ID, requestOptions)
	fmt.Printf("Chapters by Book ID: %v\n", chaptersByBookID)
}

func exampleMovieRequest(lotrSDK sdk.LotrSDK) {
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
		Filters: []sdk.Filter{
			{
				Option: sdk.GreaterThanOrEqualTo,
				Field:  "academyAwardWins",
				Value:  "5",
			},
		},
	}

	movies, _ := lotrSDK.GetMovieList(requestOptions)
	fmt.Printf("Movies: %v\n", movies)

	movieByID, _ := lotrSDK.GetMovieByID(movies[0].ID)
	fmt.Printf("Movie by ID: %v\n", movieByID)
}

func exampleCharacterRequest(lotrSDK sdk.LotrSDK) {
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
		Filters: []sdk.Filter{
			{
				Option: sdk.Includes,
				Field:  "race",
				Value:  "Hobbit,Human",
			},
		},
	}

	characters, _ := lotrSDK.GetCharacterList(requestOptions)
	fmt.Printf("Characters: %v\n", characters)

	characterByID, _ := lotrSDK.GetCharacterByID(characters[0].ID)
	fmt.Printf("Character by ID: %v\n", characterByID)
}

func exampleQuoteRequest(lotrSDK sdk.LotrSDK) {
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
	}

	quotes, _ := lotrSDK.GetQuoteList(requestOptions)
	fmt.Printf("Quotes: %v\n", quotes)

	quoteByID, _ := lotrSDK.GetQuoteByID(quotes[0].ID)
	fmt.Printf("Quote by ID: %v\n", quoteByID)

	movies, _ := lotrSDK.GetMovieList(requestOptions)
	quotesByMovieID, _ := lotrSDK.GetQuoteByMovieID(movies[0].ID, requestOptions)
	fmt.Printf("Quotes by Movie ID: %v\n", quotesByMovieID)

	characters, _ := lotrSDK.GetCharacterList(requestOptions)
	quotesByCharacterID, _ := lotrSDK.GetQuoteByCharacterID(characters[0].ID, requestOptions)
	fmt.Printf("Quotes by Character ID: %v\n", quotesByCharacterID)
}
