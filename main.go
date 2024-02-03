package main

import (
	_ "context"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/models"
	"project/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	// err := r.DB.Create(&book).Error
	err = r.DB.Create(&book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "book succes add"})
	return err
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})

	return err
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}
	err := r.DB.Delete(bookModels, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "cannot delete book",
		})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books delete successfully",
	})
	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}

	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	fmt.Println("ini id", id)
	err := r.DB.Where("id = ?", id).First(bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "cannot get id book",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "success get books id ",
		"data":    bookModels,
	})
	return nil
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

	config := &storage.Config{
		Host:     os.Getenv("HOST_HOST"),
		Port:     os.Getenv("HOST_PORT"),
		Password: os.Getenv("HOST_PASSWORD"),
		User:     os.Getenv("HOST_USER"),
		DBName:   os.Getenv("HOST_DBNAME"),
		// SSLMode:  os.Getenv("HOST_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("tidak dapat load database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("tidak bisa migrate database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRouters(app)
	app.Listen(":8080")
}
