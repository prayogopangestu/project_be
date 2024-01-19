package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRouters(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_books/:id", r.DeleteBook)
	api.Get("/get_books/id", r.GetBookById)
	api.Get("/books", r.GetBooks)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := controller.NewConnection(config)

	if err := nil {
		log.Fatal("tidak dapat load database")
	}


	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRouters(app)
	app.Listen(":8080")
}
