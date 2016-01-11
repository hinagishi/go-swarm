package swarm

type Options struct {
	Limit           int
	Offset          int
	sort            string
	AfterTimestamp  int
	BeforeTimestamp int
}

type MetaInfo struct {
	Code int `json:"code"`
}

type Resjson struct {
	Meta MetaInfo `json:"meta"`
	Res  Response `json:"response"`
}

type Response struct {
	Checkins Checkin `json:"checkins"`
}

type Checkin struct {
	Count int    `json:"count"`
	Items []Item `json:"items"`
}

type Item struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Type      string `json:"type"`
	Shout     string `json:"shout"`
	TimeZone  int    `json:"timeZoneOffset"`
}
