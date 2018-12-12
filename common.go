package d3auth

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

func HttpGet(geturl string) (string, error) {
	resp, err := http.Get(geturl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func HttpPost(geturl string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, err := http.NewRequest("POST", geturl, nil)

	if err != nil {
		return "", err
	}

	response, _ := client.Do(reqest)

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
