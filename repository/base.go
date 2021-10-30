package repo

import (
	db "github.com/radhianamri/toggl-cardgame/lib/database"
)

func Save(entity interface{}) error {
	tx := db.GetConn().Begin()
	if err := tx.Create(entity).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
