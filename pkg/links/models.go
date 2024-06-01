package links

import "time"

type Link struct {
	Short    string    `json:"short"`
	Url      string    `json:"url"`
	HitCount uint      `json:"hitCount"`
	LastHit  time.Time `json:"lastHit"`
	Apikey   string    `json:"apikey"`
	IPaddr   string    `json:"ipaddr"`
}
