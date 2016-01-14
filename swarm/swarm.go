package swarm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Oauth struct {
	ClientID     string
	ResponseType string
	RedirectURI  string
}

func GetCheckIns(v *url.Values) *Resjson {
	url := "https://api.foursquare.com/v2/users/self/checkins?" + v.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	fmt.Println(url)
	result := new(Resjson)
	json.Unmarshal(body, result)

	return result
}
