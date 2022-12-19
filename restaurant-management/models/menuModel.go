package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bean/primitive"
)

type Menu struct {
	ID 				primite.ObjectID		`bson:"_id"`
	Name			string 					`json:"name" validate:"required"`
	Category		string 					`json:"category" validate:"required"`
	Start_Data		*time.Time 				`json:"start_date"`
	End_Date		*time.Time 				`json:"end_Date"`
	Created_at		time.Time 				`json:"created_at"`
	Updated_at		time.Time 				`json:"updated_at"`
	Menu_id			string 					`json:"menu_id" validate:"required"`
}