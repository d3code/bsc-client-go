package cmc

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
)

const baseUrl = "https://api.bscscan.com/api"

type client struct {
    apiKey        string
    printResponse bool
}

func Client(apiKey string) *client {
    c := &client{
        apiKey:        apiKey,
        printResponse: false,
    }

    return c
}

func (c *client) PrintResponse(printResponse bool) *client {
    c.printResponse = printResponse
    return c
}

func doGetRequest(q url.Values, c *client) ([]byte, error) {
    req, err := http.NewRequest(http.MethodGet, baseUrl, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Accepts", "application/json")

    q.Add("apikey", c.apiKey)
    req.URL.RawQuery = q.Encode()

    httpClient := http.Client{
        Timeout: 30 * time.Second,
    }

    res, requestError := httpClient.Do(req)

    if requestError != nil {
        return nil, requestError
    }

    resBody, responseError := ioutil.ReadAll(res.Body)
    if responseError != nil {
        return nil, responseError
    }

    if c.printResponse && resBody != nil {
        responseString := string(resBody)
        println(responseString)
    }

    return resBody, nil
}
