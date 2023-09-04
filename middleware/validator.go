package middleware

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var Validator = validator.New()

type IError struct {
	Field string
	Tag   string
	Value string
}

type LoginInput struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthorInput struct {
	Name string `json:"name" validate:"required"`
}

type BookInput struct {
	Title       string    `json:"title" validate:"required"`
	Subtitle    string    `json:"subtitle" validate:"required"`
	Description string    `json:"description" validate:"required"`
	PublisherID uuid.UUID `json:"publisher-id" validate:"required"`
	AuthorName  string    `json:"author-name" validate:"required"`
}

type PublisherInput struct {
	Name string `json:"name" validate:"required"`
}

func ValidateUserLogin(c *fiber.Ctx) error {
	var errors []*IError
	user := new(LoginInput)
	c.BodyParser(&user)

	err := Validator.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = fmt.Sprint(err.Value())
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func ValidateBookInput(c *fiber.Ctx) error {
	var errors []*IError
	book := new(BookInput)
	c.BodyParser(&book)

	err := Validator.Struct(book)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = fmt.Sprint(err.Value())
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func ValidateAuthorInput(c *fiber.Ctx) error {
	var errors []*IError
	author := new(AuthorInput)
	c.BodyParser(&author)

	err := Validator.Struct(author)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = fmt.Sprint(err.Value())
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}

func ValidatePublisherInput(c *fiber.Ctx) error {
	var errors []*IError
	publisher := new(PublisherInput)
	c.BodyParser(&publisher)

	err := Validator.Struct(publisher)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = fmt.Sprint(err.Value())
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
