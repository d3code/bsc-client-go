package test

import (
    "fmt"
    cmc "github.com/d3code/bsc-client-go/bsc"
    "net/url"
    "os"
    "testing"
)

var testClient = cmc.Client(os.Getenv("BSC_KEY")).PrintResponse(true)

func TestGetAccountTokenTx(t *testing.T) {
    values := url.Values{}
    values.Add("address", "0x5d28d1ae01d00cb04cae37797458989c4b3b9f23")

    response, err := testClient.GetAccountTokenTx(values)
    if err != nil {
        t.Error(err)
    }

    fmt.Println(response)
}

func TestGetAccountTxList(t *testing.T) {
    values := url.Values{}
    values.Add("address", "0x5d28d1ae01d00cb04cae37797458989c4b3b9f23")

    response, err := testClient.GetAccountTxList(values)
    if err != nil {
        t.Error(err)
    }

    fmt.Println(response)
}
