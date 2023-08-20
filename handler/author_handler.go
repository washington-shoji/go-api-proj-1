package handler

import (
	"washington/go_api/database"
	"washington/go_api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create a author
func CreateAuthor(c *fiber.Ctx) error {
	db := database.DB.Db
	author := new(model.Author)
	// Store the body in the user and return error if encountered
	err := c.BodyParser(author)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&author).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create author", "data": err})
	}
	// Return the created author
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Author has created", "data": author})
}

// Get All Users from db
func GetAllAuthors(c *fiber.Ctx) error {
	db := database.DB.Db
	var authors []model.Author
	// find all users in the database
	db.Preload("Books").Find(&authors)
	// If no user found, return an error
	if len(authors) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Authors not found", "data": nil})
	}
	// return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Authors Found", "data": authors})
}

// GetSingleAuthor from db
func GetSingleAuthor(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var author model.Author
	// find single book in the database by id
	db.Find(&author, "id = ?", id)
	if author.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Author not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Author Found", "data": author})
}

// update an author in db
func UpdateAuthor(c *fiber.Ctx) error {
	db := database.DB.Db
	var author model.Author
	// get id params
	id := c.Params("id")
	// find single author in the database by id
	db.Find(&author, "id = ?", id)
	if author.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Author not found", "data": nil})
	}
	err := c.BodyParser(&author)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Save the Changes
	db.Save(&author)
	// Return the updated book
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Author Found", "data": author})
}

// Delete author in db by ID
func DeleteAuthorByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var author model.Author
	// get id params
	id := c.Params("id")
	// find single author in the database by id
	db.Find(&author, "id = ?", id)
	if author.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Author not found", "data": nil})
	}
	err := db.Delete(&author, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete author", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Author deleted"})
}
