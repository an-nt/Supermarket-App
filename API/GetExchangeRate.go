package API

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	//domain = "http://192.168.255.2:8080"
	//domain = "http://www.mysupermarket.com:8080"
	domain = "http://host.docker.internal:8080"
	id     = 1
	pass   = "123456"
)

type LoginFormat struct {
	ID   uint
	Pass string
}
type Response struct {
	Message string      `json: "message"`
	Detail  interface{} `json: "detail"`
}

func (a *Api) GetExchangeRate() (uint, error) {
	token, err := authenticate()
	if err != nil {
		return 0, err
	}

	URL, err := url.Parse(fmt.Sprintf("%s/v1/exchangerates/USDVND", domain))
	if err != nil {
		return 0, err
	}
	req := &http.Request{
		Method: "GET",
		URL:    URL,
		Header: map[string][]string{
			"Content-type":  {"application/json"}, //input format
			"Accept":        {"application/json"}, //output format
			"Authorization": {token},
		},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil
	}

	resBody, _, _ := bufio.NewReader(res.Body).ReadLine()

	var resp Response
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return 0, err
	}

	rate, err := strconv.Atoi(fmt.Sprint(resp.Detail))
	if err != nil {
		return 0, err
	}
	return uint(rate), nil
}

func authenticate() (string, error) {
	URL, err := url.Parse(fmt.Sprintf("%s/v1/login", domain))
	if err != nil {
		return "", err
	}

	credentials := LoginFormat{
		ID:   id,
		Pass: pass,
	}

	reqBody, err := json.Marshal(credentials)
	if err != nil {
		return "", err
	}

	req := &http.Request{
		Method: "POST",
		URL:    URL,
		Header: map[string][]string{
			"Content-type": {"application/json"},
			"Accept":       {"application/json"},
		},
		Body: ioutil.NopCloser(bytes.NewBuffer(reqBody)),
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resBody, _, _ := bufio.NewReader(res.Body).ReadLine()

	var resp Response
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return "", err
	}

	token := fmt.Sprint(resp.Detail)
	return token, nil
}
