package models

import (
	"time"
)

type Article_sections struct {
	Section_number      	int 	`json:"part_number" gorm:"primary_key"`
	Section_description 	string  	`json:"part_description"`
	Uid        			string  	`json:"part_uuid"`
	Create_date			time.Time 	`json:"create_date"`
}