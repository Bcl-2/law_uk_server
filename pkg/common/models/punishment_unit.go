package models

type Punishment_units struct {
	Id               int    `json:"id" gorm:"primary_key"`
	Unit_description string `json:"unit_description"`
}