package handler

import (
	"washington/go_api/database"
	"washington/go_api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create a user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

// Get All Users from db
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User
	// find all users in the database
	db.Find(&users)
	// If no user found, return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	// return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

// GetSingleUser from db
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var user model.User
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	db := database.DB.Db
	var user model.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	user.Username = updateUserData.Username
	user.Email = updateUserData.Email
	user.Password = updateUserData.Password
	// Save the Changes
	db.Save(&user)
	// Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users Found", "data": user})
}

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

// Create a book
func CreateBook(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(model.Book)
	// Store the body in the book and return error if encountered
	err := c.BodyParser(book)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	err = db.Create(&book).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create book", "data": err})
	}

	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Book has created", "data": book})
}

// Get All Users from db
func GetAllBooks(c *fiber.Ctx) error {
	db := database.DB.Db
	var books []model.Book
	// find all books in the database
	db.Find(&books)
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

// Create a user
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
	db.Find(&authors)
	// If no user found, return an error
	if len(authors) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Authors not found", "data": nil})
	}
	// return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Authors Found", "data": authors})
}
