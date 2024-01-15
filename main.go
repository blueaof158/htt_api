package main

import (
	auth "HTTApi/Auth"
	controller "HTTApi/Controller"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Post("/api/login", auth.Login)
	app.Post("/api/gettoken", getToken)
	// app.Use(auth.CheckMiddleware)

	app.Get("/api/getenv/:name", getEnv)

	app.Get("/api/getcartypes", controller.GetCarTypes)
	app.Get("/api/getcartype/:cartypeid", controller.GetCarType)
	// app.Post("/api/addcartype/:cartypeid", controller.GetCarType)
	// app.Put("/api/updatecartype/:cartypeid", controller.UpdateCarType)
	// app.Delete("/api/deletecartype/:cartypeid", controller.DeleteCarType)

	app.Get("/api/getawards", controller.GetAwards)
	app.Get("/api/getaward/:awardid", controller.GetAward)
	// app.Post("/api/addawards/:cartypeid", controller.GetCarType)
	// app.Put("/api/updateawards/:cartypeid", controller.UpdateCarType)
	// app.Delete("/api/deleteawards/:cartypeid", controller.DeleteCarType)
	app.Get("/api/updateaward/:awardid", controller.GetAward)

	app.Get("/api/cars", controller.GetCars)
	app.Get("/api/car/:carid", controller.GetCar)
	// app.Post("/api/addcar/:carid", controller.GetCarType)
	// app.Put("/api/updatecar/:cartypeid", controller.UpdateCarType)
	// app.Delete("/api/deletecar/:cartypeid", controller.DeleteCarType)
	app.Get("/api/updateacar/:awardid", controller.GetAward)

	// app.Get("/api/contents", controller.Getcontents)
	// app.Get("/api/content/:contentid", controller.Getcontent)
	// app.Post("/api/addcontent/:contentid", controller.GetcontentType)
	// app.Put("/api/updatecontent/:contenttypeid", controller.UpdatecontentType)
	// app.Delete("/api/deletecontent/:contenttypeid", controller.DeletecontentType)
	// app.Get("/api/updateacontent/:awardid", controller.GetAward)

	app.Listen(":8080")

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
	})
}

func getToken(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"token": auth.Token,
	})
}
