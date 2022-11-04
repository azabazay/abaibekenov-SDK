package main

import (
	"fmt"

	"github.com/abaibekenov-SDK/sdk"
)

func main() {
	lotrSDK := sdk.InitFromConfig(sdk.DefaultConfigFileName)

	resp := lotrSDK.GetBook()
	fmt.Println(resp)
}
