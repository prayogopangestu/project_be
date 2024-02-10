package repository

import (
	// "time"

	"fmt"
	"project/internal/dto"
	"project/internal/models"
	"strconv"

	"gorm.io/gorm"
	// "github.com/google/uuid"
)

type RmCodeRepository interface {
	InsertRmCode(b models.RmCodeM) (*models.RmCodeM, error)
	UpdateRmCode(b models.RmCodeM) (*models.RmCodeM, error)
	DeleteRmCode(b models.RmCodeM) error
	AllRmCode() (*[]models.RmCodeM, error)
	FindRmCodeByID(RmCodeID uint64) (*models.RmCodeM, error)
	FindRmCode(RmCodeID string) (*[]models.RmCodeM, error)
	FindRmCodeByFilter(listRmCodeDTO dto.ListRmCodeDTO) (*[]models.RmCodeM, *int64, error)
}

type RmCodeConnection struct {
	connection *gorm.DB
}

// NewRmCodeRepository creates an instance RmCodeRepository
func NewRmCodeRepository(dbConn *gorm.DB) RmCodeRepository {
	return &RmCodeConnection{
		connection: dbConn,
	}
}

func (db *RmCodeConnection) InsertRmCode(b models.RmCodeM) (*models.RmCodeM, error) {

	prosesRmCode := db.connection.Debug().Create(&b)

	if prosesRmCode.Error != nil {
		fmt.Println("ada error di proses", prosesRmCode.Error)
		return nil, prosesRmCode.Error
	}

	db.connection.Preload("CreateUser").Find(&b)

	return &b, nil
}

func (db *RmCodeConnection) UpdateRmCode(b models.RmCodeM) (*models.RmCodeM, error) {

	prosesRmCode := db.connection.Debug().Updates(&b)
	if prosesRmCode.Error != nil {
		fmt.Println("ada error di proses", prosesRmCode.Error)
		return nil, prosesRmCode.Error
	}
	db.connection.Preload("UpdateUser").Find(&b)
	return &b, nil
}

func (db *RmCodeConnection) DeleteRmCode(b models.RmCodeM) error {

	trx := db.connection.Begin()

	defer func() {
		if r := recover(); r != nil {
			trx.Rollback()
		}
	}()

	if err := trx.Error; err != nil {
		return err
	}
	if b.DeleteUserID != nil {
		prosesRmCode := trx.Model(b).Where(models.RmCodeM{Model: gorm.Model{ID: b.ID}}).Update("delete_user_id", b.DeleteUserID)
		if prosesRmCode.Error != nil {
			trx.Rollback()
			fmt.Println("ada error di proses", prosesRmCode.Error)
			return prosesRmCode.Error
		}
	}

	prosesRmCode := trx.Delete(&b)
	if prosesRmCode.Error != nil {
		trx.Rollback()
		fmt.Println("ada error di proses", prosesRmCode.Error)
		return prosesRmCode.Error
	}
	trx.Commit()
	return nil
}

func (db *RmCodeConnection) FindRmCodeByID(RmCodeID uint64) (*models.RmCodeM, error) {
	var RmCode models.RmCodeM
	prosesRmCode := db.connection.Preload("UpdateUser").Preload("CreateUser").Find(&RmCode, RmCodeID)
	if prosesRmCode.Error != nil {
		fmt.Println("ada error di proses", prosesRmCode.Error)
		return nil, prosesRmCode.Error
	}

	return &RmCode, nil
}

func (db *RmCodeConnection) FindRmCode(RmCodeID string) (*[]models.RmCodeM, error) {
	var RmCode []models.RmCodeM
	prosesRmCode := db.connection.Where("nama ILIKE ? ", "%"+RmCodeID+"%").Find(&RmCode)
	db.connection.Find(&RmCode, RmCodeID)
	if prosesRmCode.Error != nil {
		fmt.Println("ada error di proses", prosesRmCode.Error)
		return nil, prosesRmCode.Error
	}
	return &RmCode, nil
}

func (db *RmCodeConnection) AllRmCode() (*[]models.RmCodeM, error) {
	var RmCodes []models.RmCodeM
	prosesRmCode := db.connection.Find(&RmCodes)
	if prosesRmCode.Error != nil {
		fmt.Println("ada error di proses", prosesRmCode.Error)
		return nil, prosesRmCode.Error
	}
	return &RmCodes, nil
}

func (db *RmCodeConnection) FindRmCodeByFilter(listRmCodeDTO dto.ListRmCodeDTO) (*[]models.RmCodeM, *int64, error) {

	limits, _ := strconv.Atoi(listRmCodeDTO.Limit)
	pages, _ := strconv.Atoi(listRmCodeDTO.Page)
	// Sort := "id asc"
	var count_ int64
	offset := (pages - 1) * limits
	var RmCode []models.RmCodeM
	// var x []RmCode_try
	// queryBuider := db.connection.Limit(limits).Offset(offset).Order(Order)
	queryBuider := db.connection.Debug().Model(&models.RmCodeM{})

	// queryBuider = queryBuider.Joins("") //wip sampe sini
	// result := queryBuider.Model(&models.User{}).Where(user).Find(&users)
	if listRmCodeDTO.Search != nil {
		queryBuider = queryBuider.Debug().Where("rm_code_nama  ILIKE ? OR rm_code_kode ILIKE ?  OR rm_code_deskripsi ILIKE ?", "%"+*listRmCodeDTO.Search+"%", "%"+*listRmCodeDTO.Search+"%", "%"+*listRmCodeDTO.Search+"%")
	}
	prosesCount := queryBuider.Debug().Count(&count_)
	if prosesCount.Error != nil {
		fmt.Println("ada error di prose count", prosesCount.Error)
		return nil, &count_, prosesCount.Error
	}
	proses := queryBuider.Debug().Limit(limits).Offset(offset).Order(listRmCodeDTO.Order).Find(&RmCode)
	if proses.Error != nil {
		fmt.Println("ada error di proses", proses.Error)
		return nil, &count_, proses.Error
	}

	return &RmCode, &count_, nil
}
