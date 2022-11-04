package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type lotrSDK struct {
	config Config
}

// InitFromConfig initializes lotrSDK from config file.
func InitFromConfig(configName string) lotrSDK {
	sdk := lotrSDK{}
	sdk.loadConfig(configName)

	return sdk
}

// PUBLIC METHODS

// END PUBLIC METHODS

// PRIVATE METHODS

func (l lotrSDK) processError(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())

	// os.Exit(1)
}

func (l *lotrSDK) loadConfig(configFileName string) {
	f, err := os.Open(configFileName)
	if err != nil {
		l.processError(err)
	}
	defer f.Close()

	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		l.processError(err)
	}

	l.config = cfg
}

func (l lotrSDK) makeHTTPRequest(url string, method string, body io.Reader) (Response, error) {
	var response Response

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, fmt.Errorf("client: could not create request: %s", err)
	}

	if l.config.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", l.config.Token))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("client: error making http request: %s", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("client: could not read response body: %s", err)
	}

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return response, fmt.Errorf("client: unmarshal error: %s", err)
	}

	return response, nil
}

// END PRIVATE METHODS
