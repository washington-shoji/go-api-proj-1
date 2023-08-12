package handler

import (
	"washington/go_api/database"
	"washington/go_api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create publisher
func CreatePublisher(c *fiber.Ctx) error {
	db := database.DB.Db
	publisher := new(model.Publisher)
	// Store the body in the publisher and return error if encountered
	err := c.BodyParser(publisher)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&publisher).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create publisher", "data": err})
	}
	// Return the created publisher
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Publisher has created", "data": publisher})
}

// Get All Users from db
func GetAllPublishers(c *fiber.Ctx) error {
	db := database.DB.Db
	var publisher []model.Publisher
	// find all publishers in the database
	db.Find(&publisher)
	// If no publisher found, return an error
	if len(publisher) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Publisher not found", "data": nil})
	}
	// return publishers
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Publisher Found", "data": publisher})
}

// GetSingle publisher from db
func GetSinglePublisher(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var publisher model.Publisher
	// find single publisher in the database by id
	db.Find(&publisher, "id = ?", id)
	if publisher.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Publisher not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Publisher Found", "data": publisher})
}

// update an publisher in db
func UpdatePublisher(c *fiber.Ctx) error {
	db := database.DB.Db
	var publisher model.Publisher
	// get id params
	id := c.Params("id")
	// find single publisher in the database by id
	db.Find(&publisher, "id = ?", id)
	if publisher.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Publisher not found", "data": nil})
	}
	err := c.BodyParser(&publisher)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Save the Changes
	db.Save(&publisher)
	// Return the updated publisher
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Publisher Found", "data": publisher})
}

// Delete publisher in db by ID
func DeletePublisherByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var publisher model.Publisher
	// get id params
	id := c.Params("id")
	// find single publisher in the database by id
	db.Find(&publisher, "id = ?", id)
	if publisher.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Publisher not found", "data": nil})
	}
	err := db.Delete(&publisher, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete publisher", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Publisher deleted"})
}