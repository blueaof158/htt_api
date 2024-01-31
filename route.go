package main

import (
	auth "HTTApi/Auth"
	controller "HTTApi/Controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// ##Open Model

// ##Close Model

func route(app *fiber.App) {

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
	app.Post("/api/addcartype", controller.AddCarType)
	app.Put("/api/updatecartype/:cartypeid", controller.UpdateCarType)
	app.Delete("/api/deletecartype/:cartypeid", controller.DeleteCarType)

	app.Get("/api/getawards", controller.GetAwards)
	app.Get("/api/getaward/:awardid", controller.GetAward)
	app.Post("/api/addaward/", controller.AddAward)
	app.Put("/api/updateawards/:awardid", controller.UpdateAward)
	app.Delete("/api/deleteawards/:awardid", controller.DeleteAward)
	app.Get("/api/updateaward/:awardid", controller.GetAward)

	app.Get("/api/cars", controller.GetCars)
	app.Get("/api/car/:carid", controller.GetCar)
	app.Post("/api/addcar/:carid", controller.AddCar)
	app.Put("/api/updatecar/:carid", controller.UpdateCar)
	app.Delete("/api/deletecar/:carid", controller.DeleteCarType)

	app.Get("/api/contents", controller.GetContents)
	app.Get("/api/content/:contentid", controller.GetContent)
	app.Post("/api/addcontent/:contentid", controller.AddContent)
	app.Put("/api/updatecontent/:contentid", controller.UpdateContent)
	app.Delete("/api/deletecontent/:contentid", controller.DeleteContent)

	app.Listen(":8080")

}
