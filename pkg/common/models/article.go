package models

import (
	"time"
)

type Article struct {
	Article_number      float64 `json:"article_number"`
	Article_description string  `json:"article_description"`
	Uid        			string  `json:"article_uuid"`
	Create_date         time.Time `json:"create_date"`
}