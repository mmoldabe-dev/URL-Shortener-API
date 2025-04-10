package models

import (
	"time"
)

type URL struct {
	Id          int       `json:"id"`
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	TTLSecond   int       `json:"ttl_second"`
}
