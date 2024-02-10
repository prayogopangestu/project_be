package tx

import (
	"project/pkg/err"

	"gorm.io/gorm"
)

func CommitOrRollback(db *gorm.DB) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if erro := tx.Error; erro != nil {
		// return err
		err.PanicIfError(erro)
	}

	err := recover()

	if err != nil {
		// errorRollback := tx.Rollback()
		// PanicIfError(errorRollback)
		// panic(err)
		tx.Rollback()
	} else {
		// errorCommit := tx.Commit()
		// PanicIfError(errorCommit)
		tx.Commit()
	}
}
