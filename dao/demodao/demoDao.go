package demodao

import (
	"goTemplate/entity"
	"goTemplate/utils/db"
)

func InsertDemo(thisEntity *entity.Demo) error {
	mysql := db.GetDb()
	result := mysql.Create(thisEntity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
