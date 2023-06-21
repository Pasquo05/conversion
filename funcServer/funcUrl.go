package funcserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type urlPath struct {
	base string
	path string
}

type Prova struct {
	Disclaimer string
	License    string
	Timestamp  int64
	Base       string
	Rates      rates
}

func NewUrl() urlPath {

	/*
		url := ""
		path := ""
		fmt.Println("inserire l'url e il path da voler utilizzare")
		fmt.Scanf(&url)
		fmt.Scanf(&path)
	*/

	url := urlPath{base: "https://openexchangerates.org/api/", path: "latest.json"}

	return url
}

func (urlInput urlPath) ConvertUrl(key string) string {

	resource := urlInput.path
	params := url.Values{}
	params.Add("app_id", key)
	params.Add("symbols", "EUR")

	u, _ := url.ParseRequestURI(urlInput.base)
	u.Path = resource
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u) // "http://example.com/path?param1=value1&param2=value2"

	return urlStr
}

func GetRespondBody(url string) Prova {

	resp, errGet := http.Get(url)
	if errGet != nil {
		log.Fatal("panic")
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("errore")
	}
	usd := Prova{}

	err = json.Unmarshal(resBody, &usd)
	if err != nil {
		fmt.Println(err)
	}

	return usd

}
