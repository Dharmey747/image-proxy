package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ImageProxy(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Query().Get("url")

	// check if the query parameter is a url
	if url == "" || !(strings.Contains(url, "http")) {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"message": "url parameter is empty or invalid"})
		return
	}

	// make a request to the url in the query parameter, using a custom user agent to prevent a failed/denied responses
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	// check if the request has been successful
	if resp.StatusCode != 200 || err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"message": fmt.Sprintf("error proxying image (%s)", resp.Status)})
		return
	}

	// check if the response is actually an image
	if !strings.Contains(resp.Header.Get("content-type"), "image") {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"message": "proxy response is not an image"})
		return
	}

	// read the response bytes
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"message": "error parsing proxy response"})
		return
	}

	// send the response content
	writer.Write(content)
}
