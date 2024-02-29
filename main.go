package main

import (
	auth "HTTApi/Auth"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// ##Open Model

// ##Close Model

var db *sql.DB

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env failed")
	}
	app := fiber.New()
	route(app)
}

func getEnv(c *fiber.Ctx) error {
	name := c.Params("name")
	addressdb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("db_user"), os.Getenv("db_pwd"), os.Getenv("database_ip"), os.Getenv("database_port"), os.Getenv("database_name"))
	addressdba := os.Getenv("db_user") + ":" + os.Getenv("db_pwd") + "@tcp(" + os.Getenv("database_ip") + ":" + os.Getenv("database_port") + "/" + os.Getenv("database_name")

	return c.JSON(fiber.Map{
		"env":  os.Getenv(name),
		"name": name,
		"db":   addressdb,
		"db2":  addressdba,
		"aa":   time.Now().Format("2006-01-02 15:04:05"),
	})
}

func getToken(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"token": auth.Token,
	})
}
