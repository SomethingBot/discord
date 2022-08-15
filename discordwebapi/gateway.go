package discordwebapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/SomethingBot/dizzy/libinfo"
)

type GatewayWebsocketInformationSessionStartLimit struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

type GatewayWebsocketInformation struct {
	Url               string                                       `json:"url"`
	Shards            int                                          `json:"shards"`
	SessionStartLimit GatewayWebsocketInformationSessionStartLimit `json:"session_start_limit"`
}

func GetGatewayWebsocketInformation(discordApiGatewayURL string, apiKey string) (GatewayWebsocketInformation, error) {
	if discordApiGatewayURL == "" {
		discordApiGatewayURL = "https://discord.com/api/gateway/bot"
	}

	req, err := http.NewRequest("GET", discordApiGatewayURL, nil)
	if err != nil {
		return GatewayWebsocketInformation{}, err
	}
	req.Header.Set("User-Agent", libinfo.BotUserAgent)
	req.Header.Add("Authorization", "Bot "+apiKey)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GatewayWebsocketInformation{}, err
	}

	defer func() {
		if err != nil {
			err2 := resp.Body.Close()
			if err2 != nil {
				err = fmt.Errorf("could not close body (%v) after error (%w)", err2, err)
			}
		}
	}()

	if resp.StatusCode != http.StatusOK {
		var errorData []byte
		errorData, err = io.ReadAll(resp.Body)
		if err != nil {
			return GatewayWebsocketInformation{}, fmt.Errorf("could not readall data from resp.Body (%w)", err)
		}
		return GatewayWebsocketInformation{}, WebAPIError{JsonData: errorData} //todo: convert to real error type
	}

	var gwi GatewayWebsocketInformation

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&gwi)
	if err != nil {
		return GatewayWebsocketInformation{}, fmt.Errorf("could not decode json (%w)", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return GatewayWebsocketInformation{}, fmt.Errorf("could not close body (%w)", err)
	}

	return gwi, nil
}

//GetGatewayWebsocketURI returns the current Discord Gateway WSS URL, pass discordApiGatewayURL as "" to use default  //todo: make it so test doesn't have to hit server
func GetGatewayWebsocketURI(discordApiGatewayURL string) (url.URL, error) {
	if discordApiGatewayURL == "" {
		discordApiGatewayURL = "https://discord.com/api/gateway"
	}

	req, err := http.NewRequest("GET", discordApiGatewayURL, nil)
	if err != nil {
		return url.URL{}, err
	}
	req.Header.Set("User-Agent", libinfo.BotUserAgent)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return url.URL{}, err
	}

	urlJson := struct {
		Url string `json:"url"`
	}{}

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&urlJson)
	if err != nil {
		err2 := resp.Body.Close()
		if err2 != nil {
			return url.URL{}, fmt.Errorf("could not close body (%v), after error (%w)", err2, err)
		}
		return url.URL{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return url.URL{}, err
	}

	uri, err := url.ParseRequestURI(urlJson.Url)
	if err != nil {
		return url.URL{}, err
	}

	return *uri, err
}
