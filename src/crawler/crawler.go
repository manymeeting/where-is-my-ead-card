package crawler


import (
    "net/url"
    "net/http"
    "io/ioutil"
    "fmt"
    )

func Crawler() {

    params := url.Values{}

    Url, err := url.Parse("http://google.com")
    if err != nil {
        panic(err.Error())

    }
    params.Set("a", "fdfds")
    params.Set("id", string("1"))
    //如果参数中有中文参数,这个方法会进行URLEncode
    Url.RawQuery = params.Encode()
    urlPath := Url.String()
    resp, err := http.Get(urlPath)
    defer resp.Body.Close()
    s, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(s))

}