package model

import (
	"time"
)

type Model struct {
	Id          int       `json:"id" `
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	CreatedAT   time.Time `json:"created_at"`
	TTLSecond   int       `json:"ttl_second"`
}
