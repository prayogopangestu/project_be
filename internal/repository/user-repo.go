package repository

import (
	"fmt"
	"log"
	"project/internal/dto"
	"project/internal/helper"
	"project/internal/models"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) (*models.User, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
	ListUserByFilter(dtoList dto.ListUserDTO) (*[]models.User, *int64, error)
}

type userConnection struct {
	connection *gorm.DB
}

// NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) (*models.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		err := db.connection.Find(&tempUser, user.ID).Error
		if err != nil {
			return nil, err
		}
		user.Password = tempUser.Password
	}

	err := db.connection.Debug().Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).
		// Preload("Dokter_m.DokterRuangan_m.RuanganM").
		Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) models.User {
	var user models.User
	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID)
	return user
}

func (db *userConnection) ListUserByFilter(dtoList dto.ListUserDTO) (*[]models.User, *int64, error) {

	limits, _ := strconv.Atoi(dtoList.Limit)
	pages, _ := strconv.Atoi(dtoList.Page)
	// Sort := "id asc"
	var count_ int64
	offset := (pages - 1) * limits
	var User []models.User
	// var x []User_try
	// queryBuider := db.connection.Limit(limits).Offset(offset).Order(Order)
	queryBuider := db.connection.Debug().Model(&models.User{})
	// result := queryBuider.Model(&master.User{}).Where(user).Find(&users)
	// queryBuider = queryBuider.Joins(`JOIN dokter_m ON dokter_ruangan_m.dokter_m_id = dokter_m."id"`).Preload("DokterM")
	// queryBuider = queryBuider.Joins(`JOIN ruangan_m ON dokter_ruangan_m.ruangan_m_id = ruangan_m."id"`).Preload("RuanganM")

	if dtoList.Search != nil {
		queryBuider = queryBuider.Where(`"user"."first_name" ILIKE ? OR "user"."last_name" ILIKE ? OR "user"."nickname" ILIKE ? `, "%"+*dtoList.Search+"%", "%"+*dtoList.Search+"%", "%"+*dtoList.Search+"%")
	}
	prosesCount := queryBuider.Debug().Count(&count_)
	if prosesCount.Error != nil {
		fmt.Println("ada error di prose count", prosesCount.Error)
		return nil, &count_, prosesCount.Error
	}
	proses := queryBuider.Limit(limits).Offset(offset).Order(dtoList.Order).Find(&User)
	if proses.Error != nil {
		fmt.Println("ada error di proses", proses.Error)
		return nil, &count_, proses.Error
	}

	return &User, &count_, nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		helper.WriteError(err.Error())
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
