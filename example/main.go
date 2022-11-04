package main

import (
	"fmt"

	"github.com/abaibekenov-SDK/sdk"
)

func main() {
	lotrSDK := sdk.InitFromConfig(sdk.DefaultConfigFileName)

	books, resp := lotrSDK.GetBookList()
	fmt.Println(resp)

	bookByID, resp := lotrSDK.GetBookByID(books[0].ID)
	fmt.Println(bookByID)
}
