package swarm

type Resjson struct {
	Res Response `json:"response"`
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
}
