package handler

import (
	"washington/go_api/database"
	"washington/go_api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create a book
func CreateBook(c *fiber.Ctx) error {
	// Struck to pass json data for querying and updating associations
	type createBook struct {
		Title       string         `json:"title"`
		Subtitle    string         `json:"subtitle"`
		Description string         `json:"description"`
		PublisherID uuid.UUID      `json:"publisher-id"`
		Authors     []model.Author ``
		AuthorName  string         `json:"author-name"`
	}
	db := database.DB.Db
	createBookData := new(createBook)
	createBookErr := c.BodyParser(createBookData)
	if createBookErr != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": createBookErr})
	}
	author := new(model.Author)
	book := new(model.Book)
	db.Find(&author, "name = ?", createBookData.AuthorName).First(&author)
	if author.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Author not found", "data": nil})
	}

	// Mapping json data to book and author model for record
	book.Title = createBookData.Title
	book.Subtitle = createBookData.Subtitle
	book.Description = createBookData.Description
	book.PublisherID = createBookData.PublisherID
	book.Authors = []model.Author{*author}
	author.Books = []model.Book{*book}

	// Store the body in the book, update author and append book and return error if encountered
	err := db.Model(&author).Association("Books").Append(&book)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create book", "data": err})
	}

	// Return the created Book
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Book has created", "data": book})
}

// Get All Users from db
func GetAllBooks(c *fiber.Ctx) error {
	db := database.DB.Db
	var books []model.Book
	// find all books in the database
	db.Preload("Authors").Find(&books)
	// If no book found, return an error
	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Books not found", "data": nil})
	}
	// return books
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Books Found", "data": books})
}

// GetSingleUser from db
func GetSingleBook(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var book model.Book
	// find single book in the database by id
	db.Find(&book, "id = ?", id)
	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Book not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Book Found", "data": book})
}

// update a book in db
func UpdateBook(c *fiber.Ctx) error {
	// type updateBook struct {
	// 	Title       string       `json:"title"`
	// 	Subtitle    string       `json:"subtitle"`
	// 	Description string       `json:"description"`
	// 	Author      model.Author `json:"author"`
	// }
	db := database.DB.Db
	var book model.Book
	// get id params
	id := c.Params("id")
	// find single book in the database by id
	db.Find(&book, "id = ?", id)
	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Book not found", "data": nil})
	}
	//var updateBookData updateBook
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// book.Title = updateBookData.Title
	// book.Subtitle = updateBookData.Subtitle
	// book.Description = updateBookData.Description
	// book.Author = updateBookData.Author
	// Save the Changes
	db.Save(&book)
	// Return the updated book
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Book Found", "data": book})
}

// Delete book in db by ID
func DeleteBookByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var book model.Book
	// get id params
	id := c.Params("id")
	// find single book in the database by id
	db.Find(&book, "id = ?", id)
	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Book not found", "data": nil})
	}
	err := db.Delete(&book, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete book", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Book deleted"})
}
