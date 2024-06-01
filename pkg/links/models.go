package links

import "time"

type Link struct {
	Short     string    `json:"short"`
	Url       string    `json:"url"`
	HitCount  uint      `json:"hitCount"`
	LastHit   time.Time `json:"lastHit"`
	CreatedBy string    `json:"creator"`
	Apikey    string    `json:"apikey"`
	IPaddr    string    `json:"ipaddr"`
}

type JsonLink struct {
	Short     string `json:"short"`
	Url       string `json:"url"`
	HitCount  uint   `json:"hitCount"`
	LastHit   string `json:"lastHit"`
	CreatedBy string `json:"creator"`
}
