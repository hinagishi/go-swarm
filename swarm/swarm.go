package swarm

type Oauth struct {
	ClientID     string
	ResponseType string
	RedirectURI  string
}

func getCheckIns(v *url.Values) *Resjson {
	url := "https://api.foursquare.com/v2/users/self?" + v.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	result := make(Resjson)
	json.Unmarshal(body, result)
}
