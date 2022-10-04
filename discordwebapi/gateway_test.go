package discordwebapi

import (
	"os"
	"strings"
	"testing"
)

func TestGetGatewayBot(t *testing.T) {
	t.Parallel()
	//todo: make a short test that doesn't hit server

	if testing.Short() {
		t.Skipf("test short flag set, skipping integration tests")
	}

	api := API{}
	api.ApiKey = os.Getenv("discordapikey")
	if api.ApiKey == "" {
		apikeyBytes, err := os.ReadFile("../apikeyfile")
		if err != nil {
			t.Fatalf("error on reading apikeyfile (%v)\n", err)
		}
		api.ApiKey = strings.ReplaceAll(string(apikeyBytes), "\n", "")
	}

	_, err := api.GetGatewayBot()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetGatewayURI(t *testing.T) {
	t.Parallel()
	//todo: make a short test that doesn't hit server
	if testing.Short() {
		t.Skipf("test short flag set, skipping integration tests")
	}

	api := API{}
	uri, err := api.GetGateway()
	if err != nil {
		t.Fatalf("could not get gateway URI (%v)\n", err)
	}
	if uri.String() == "" {
		t.Fatal("uri returned is empty")
	}
}
