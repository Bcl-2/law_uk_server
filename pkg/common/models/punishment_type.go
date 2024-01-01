package models

type Punishments_type struct {
	Id              int    `json:"id" gorm:"primary_key"`
	Punishment_type string `json:"punishment_type"`
}

func (Punishments_type) TableName() string {
	return "punishments_type"
}