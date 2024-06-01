package links

import "time"

type Link struct {
	Short    string    `json:"short"`
	Url      string    `json:"Url"`
	HitCount uint      `json:"hitCount"`
	LastHit  time.Time `json:"lastHit"`
}
