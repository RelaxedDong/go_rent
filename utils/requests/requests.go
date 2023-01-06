package requests

import (
	"encoding/json"
	"net/http"
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
