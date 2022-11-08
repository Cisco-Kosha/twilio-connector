package httpclient

import (
	"bytes"
	"encoding/base64"

	"io/ioutil"
	"net/http"
	"net/url"
	// "fmt"
	// "strconv"
)

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func makeHttpReq(username string, password string, req *http.Request, params url.Values) []byte {
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	req.URL.RawQuery = params.Encode()
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func GetAccount(username string, password string, params url.Values) string {
	req, err := http.NewRequest("GET", "https://api.twilio.com/2010-04-01"+"/Accounts", nil)
	if err != nil {
		return ""
	}

	res := makeHttpReq(username, password, req, params)
	return string(res)
}

func SendWhatsAppMessage(username string, password string, body []byte, params url.Values) (string, error) {
	req, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01"+"/Accounts/"+username+"/Messages.json", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}

func MakePhoneCall(username string, password string, body []byte, params url.Values) (string, error) {
	req, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01"+"/Accounts/"+username+"/Calls.json", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return "", err
	}
	return string(makeHttpReq(username, password, req, params)), nil
}
