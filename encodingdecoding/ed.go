/*
Working with and using JSON is more
common nowadays than working with XML
This is primarily becaused USING json
requires less markup than working with xml
This means less data per message needs to be
sent over the network which helps with
the overall performance of the system
JSON can be transformed into BSON (Binary Javascript Object Notation)
which reduces the size of each message even further

*/

// decoding JSON
// first aspect is to explore NewDecoder func
// and Decode method from the json package
// If you are consuming JSON from a webresponse or a file
// NewDecoder and Decode are to be remebbered

// Let's look at an example that works with the http package to perform a Get request against the Google Search API that returns JSON in reponse

// This is how the JSON response looks like

/*
{    "responseData":
{        "results": [
            {                "GsearchResultClass": "GwebSearch",
                            "unescapedUrl": "https://www.reddit.com/r/golang",                "url": "https://www.reddit.com/r/golang",
							"visibleUrl": "www.reddit.com",
							"cacheUrl": "http://www.google.com/search?q=cache:W...",
							"title": "r/\u003cb\u003eGolang\u003c/b\u003e - Reddit",
							"titleNoFormatting": "r/Golang - Reddit",
							"content": "First Open Source \u003cb\u003eGolang\u..."
							            },
							            {
							               "GsearchResultClass": "GwebSearch",
							           ]                "unescapedUrl": "http://tour.golang.org/",                "url": "http://tour.golang.org/",                "visibleUrl": "tour.golang.org",                "cacheUrl": "http://www.google.com/search?q=cache:O...",                "title": "A Tour of Go",                "titleNoFormatting": "A Tour of Go",                "content": "Welcome to a tour of the Go programming ..."            }        ]    }




*/

// so we have to retrieve and decode the response into a struct type so that we can use it

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (

	// gResult maps to the result document received from the search

	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}
	// gResponse contains the top level document
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln("error")
	}
	defer resp.Body.Close()

	// Decode the JSON response into our struct type
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Fatalln("error")
	}

	fmt.Println(gr)

}
