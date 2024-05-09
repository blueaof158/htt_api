package route

import (
	auth "HTTApi/Auth"
	controller "HTTApi/Controller"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Route(app *fiber.App) {

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
	app.Get("/api/getcartypeslst", controller.GetCarTypesLst)
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
	app.Post("/api/addcar", controller.AddCar)
	app.Put("/api/updatecar/:carid", controller.UpdateCar)
	app.Delete("/api/deletecar/:carid", controller.DeleteCar)

	app.Get("/api/contents", controller.GetContents)
	app.Get("/api/content/:contentid", controller.GetContent)
	app.Post("/api/addcontent/:contentid", controller.AddContent)
	app.Put("/api/updatecontent/:contentid", controller.UpdateContent)
	app.Delete("/api/deletecontent/:contentid", controller.DeleteContent)

	app.Get("/api/executives", controller.GetExecutives)
	app.Get("/api/executive/:executivesid", controller.GetExecutive)
	app.Post("/api/addexecutive/:executivesid", controller.AddExecutive)
	app.Put("/api/updateexecutive/:executivesid", controller.UpdateExecutive)
	app.Delete("/api/deleteexecutive/:executivesid", controller.DeleteExecutive)

	app.Get("/api/bannertops", controller.GetBannerTops)
	app.Get("/api/bannertop/:bannertopid", controller.GetBannerTop)
	app.Post("/api/addbannertop", controller.AddBannerTop)
	app.Put("/api/updatebannertop/:bannertopid", controller.UpdateBannerTop)
	app.Delete("/api/deletebannertop/:bannertopid", controller.DeleteBannerTop)

	app.Post("/api/addimages/", controller.AddImage)
	app.Get("/api/getimages/:imagetype/:id", controller.GetImage)

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
		"aa":   time.Now().Format("2006-01-02 15:04:05"),
	})
}

func getToken(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"token": auth.Token,
	})
}
