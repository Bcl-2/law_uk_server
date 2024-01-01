package models

import (
	"time"
)

type Article_parts struct {
	Part_number      	float64 	`json:"part_number"`
	Part_description 	string  	`json:"part_description"`
	Uid        			string		`gorm:"primaryKey"`
	Create_date			time.Time 	`json:"create_date"`
	Article_number		float64 	`json:"article_number"`
}