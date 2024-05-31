package auth

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var Token = ""

type CheckToken struct {
	Token string `json:"token"`
}

func CheckMiddleware(c *fiber.Ctx) error {

	tken := new(CheckToken)
	if err := c.BodyParser(tken); err != nil {
		return fiber.ErrUnauthorized
	}
	if tken.Token != Token {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

// var userMember = model.Login{
// 	Email:    "test@mail.com",
// 	Password: "12345@",
// }

// func Login(c *fiber.Ctx) error {

// 	user := new(model.Login)
// 	if err := c.BodyParser(user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}
// 	g := uuid.New()

// 	if user.Email != userMember.Email && user.Password != userMember.Email {
// 		return fiber.ErrUnauthorized
// 	}
// 	Token = g.String()
// 	return c.JSON(fiber.Map{
// 		"message": "Login Successful",
// 		"token":   g,
// 	})

// }
