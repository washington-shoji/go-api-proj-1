package router

import (
	"washington/go_api/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// grouping
	api := app.Group("/api")

	user := api.Group("/user")
	book := api.Group("/book")

	// routes
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUserByID)

	book.Get("/", handler.GetAllBooks)
	book.Get("/:id", handler.GetSingleBook)
	book.Post("/", handler.CreateBook)
	book.Put("/:id", handler.UpdateBook)
	book.Delete("/:id", handler.DeleteBookByID)

}
