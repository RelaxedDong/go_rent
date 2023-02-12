package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 5 * time.Second}

func Get(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func JsonPostGetBody(url string, jsonData interface{}) (result []byte, err error) {
	data, _ := json.Marshal(jsonData)
	fmt.Println("string(data)", string(data))
	b := strings.NewReader(string(data))
	r, err := myClient.Post(url, "application/json", b)
	if err != nil {
		return result, err
	}
	defer r.Body.Close()
	result, err = ioutil.ReadAll(r.Body)
	return result, err
}
