package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	// SSLMode  string
	// json, err := json.Marshal(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := http.Post("http://exmample.com/api/user", "application/json", bytes.NewReader(json))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// var u []User
	// decoder := json.NewDecoder(resp.Body)
	// decoder.Decode(&u)

}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn :=
		fmt.Sprintf(
			"host=%s port=%s password=%s dbname=%s ",
			config.Host, config.Port, config.Password, config.User, config.DBName,
		)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
