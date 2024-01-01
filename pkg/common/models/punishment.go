package models

import (
	"time"
)

type Article_punishments struct {
	Punishment_uid            	string           `json:"punishment_uuid"`
	Create_date               	time.Time        `json:"create_date"`
	Section_ref 				Article_sections `json:"section" gorm:"foreignKey:Section_number" references:"Section_number"`
	Chapter_ref                 Article_Chapters `json:"chapter" gorm:"foreignKey:Chapter_number" references:"Chapter_number"`
	Article_ref 				Article 		 `json:"article" gorm:"foreignKey:Article_number" references:"Article_number"`
	Part_ref 					Article_parts    `json:"part"    gorm:"foreignKey:Part_uid"       references:"Uid"`
	Punishment_ref 				Punishments_type `json:"punishment_type" gorm:"foreignKey:Article_punishment" references:"Id"`
	Punishment_min_period     	int              `json:"punishment_min_period"`
	Punishment_min_unit_ref 	Punishment_units `json:"punishment_min_unit" gorm:"foreignKey:Punishment_min_type_period" references:"Id"` 
	Punishment_max_period     	int              `json:"punishment_max_period"`
	Punishment_max_unit_ref  	Punishment_units `json:"punishment_max_unit" gorm:"foreignKey:Punishment_min_type_period" references:"Id"` 
	Danger_ref 					Article_dangers  `json:"article_danger" gorm:"foreignKey:Article_danger" references:"Id"`
	Section_number            	int				 `json:"-"`
	Article_danger            	int				 `json:"-"`
	Chapter_number            	int				 `json:"-"`
	Article_number            	float64			 `json:"-"`
	Part_uid                  	string			 `json:"-"`
	Article_punishment        	int				 `json:"-"`
	Punishment_min_type_period  int				 `json:"-"`
	Punishment_max_type_period  int  			 `json:"-"`
}