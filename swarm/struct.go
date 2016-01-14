package swarm

type Options struct {
	Limit           int
	Offset          int
	sort            string
	AfterTimestamp  int
	BeforeTimestamp int
}

type Resjson struct {
	Meta  MetaInfo       `json:"meta"`
	Notif []Notification `json:"notifications"`
	Res   Response       `json:"response"`
}

type MetaInfo struct {
	Code int `json:"code"`
}

type Notification struct {
	Type string    `json:"type"`
	Item NotifItem `json:"item"`
}

type NotifItem struct {
	UnreadCount int `json:"unreadCount"`
}

type Response struct {
	Checkins Checkin `json:"checkIns"`
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
