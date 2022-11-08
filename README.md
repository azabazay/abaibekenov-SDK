# abaibekenov-SDK
LOTR SDK

Requirements:
- `lotr.config.json` file with following structure:
```
{
  token: "..." # your token goes here
}
```

Example book requests:
```
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
	}

	books, _ := lotrSDK.GetBookList(requestOptions)

	bookByID, _ := lotrSDK.GetBookByID(books[0].ID)
```

Example movie requests:
```
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

	movieByID, _ := lotrSDK.GetMovieByID(movies[0].ID)
```

Example character requests:
```
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

	characterByID, _ := lotrSDK.GetCharacterByID(characters[0].ID)
```

Example chapter requests:
```
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
	}

	chapters, _ := lotrSDK.GetChapterList(requestOptions)

	chapterByID, _ := lotrSDK.GetChapterByID(chapters[0].ID)

	books, _ := lotrSDK.GetBookList(requestOptions)
	chaptersByBookID, _ := lotrSDK.GetChaptersByBookID(books[0].ID, requestOptions)
```

Example quote requests:
```
	requestOptions := sdk.RequestOptions{
		Limit:     1000,
		SortField: "name",
		SortType:  "asc",
	}

	quotes, _ := lotrSDK.GetQuoteList(requestOptions)

	quoteByID, _ := lotrSDK.GetQuoteByID(quotes[0].ID)

	movies, _ := lotrSDK.GetMovieList(requestOptions)
	quotesByMovieID, _ := lotrSDK.GetQuoteByMovieID(movies[0].ID, requestOptions)

	characters, _ := lotrSDK.GetCharacterList(requestOptions)
	quotesByCharacterID, _ := lotrSDK.GetQuoteByCharacterID(characters[0].ID, requestOptions)
```