package crawler


import (
    "net/url"
    "net/http"
    "io/ioutil"
    )

func FetchHTML(urlStr string) (rawHTML string) {

    params := url.Values{}

    Url, err := url.Parse(urlStr)
    if err != nil {
        panic(err.Error())

    }
    // Set params if necesary
    // Sample: 
    //params.Set("foo", "123")
    
    Url.RawQuery = params.Encode()
    urlPath := Url.String()
    resp, err := http.Get(urlPath)
    defer resp.Body.Close()
    s, err := ioutil.ReadAll(resp.Body)
    
    rawHTML = string(s)
    return rawHTML
}