package chain

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"log"

	"appengine"
	"appengine/urlfetch"
)

const BASE_URL = "https://api.chain.com/v2/bitcoin/"

func ChainUrl(path string, key string) string {
	params := url.Values{}
	return ChainUrlParams(path, params, key)
}

func ChainUrlParams(path string, params url.Values, key string) string {
	params.Add("api-key-id", key)
	return fmt.Sprintf("%s%s?%s", BASE_URL, path, params.Encode())
}

func ForwardRequest(url string, w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	resp, err := client.Get(url)
    if err != nil {
    	log.Fatal(err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    defer resp.Body.Close()

    bodyBytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(err)
        http.Error(w, readErr.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bodyBytes))
}

