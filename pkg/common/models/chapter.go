package models

import "time"

type Article_Chapters struct {
	Chapter_number      int       `db:"chapter_number" gorm:"primary_key"`
	Uid                 string    `db:"uid"`
	Chapter_description string    `db:"chapter_description"`
	Create_date         time.Time `db:"create_date"`
}