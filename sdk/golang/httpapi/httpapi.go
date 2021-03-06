package httpapi

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/jeffail/gabs"
)

const api_host = "http://my.app"                   // 修改为API域名  http://58jinrongyun.com
const api_key = "3bddc47e7cc05e1d8f488f2562969a33" // 修改为你的API key
	
func Post(module string, api string, params *gabs.Container) *gabs.Container {
	auth_str := "dyyadmin:" + api_key
	url := api_host + "/api/" + module + "/" + api
	jsonStr := []byte(params.String())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/application/json; charset=UTF-8")
	req.Header.Set("Authorization", auth_str)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		panic(err)
	}

	return jsonParsed
}
