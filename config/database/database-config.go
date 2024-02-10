package database

import (
	"fmt"
	"log"
	"os"
	"project/internal/models"

	"github.com/joho/godotenv"

	_ "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

// SetupDatabaseConnection is creating a new connection to our database
func (server *Server) SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbHost := os.Getenv("DB_HOST")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
	Dbdriver := os.Getenv("DB_DRIVER")
	var err error
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	//nanti kita isi modelnya di sini
	server.InitMigrate()
	return server.DB
}

func (server *Server) InitMigrate() {
	server.DB.Migrator().AutoMigrate(
		&models.RmCodeM{},
	)
}

// CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
