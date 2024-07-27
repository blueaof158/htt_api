package route

import (
	controller "HTTApi/Controller"

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
	// app.Post("/api/login", auth.Login)
	// app.Post("/api/gettoken", getToken)

	app.Get("api/frontend/bannertops", controller.FrontendGetBannerTops)

	// app.Use(auth.CheckMiddleware)

	// app.Get("/api/getenv/:name", getEnv)

	app.Get("/api/getcartypes", controller.GetCarTypes)
	app.Get("/api/getcartypeslst", controller.GetCarTypesLst)
	app.Get("/api/getcartype/:cartypeid", controller.GetCarType)
	app.Post("/api/addcartype", controller.AddCarType)
	app.Put("/api/updatecartype/:cartypeid", controller.UpdateCarType)
	app.Delete("/api/deletecartype/:cartypeid", controller.DeleteCarType)

	app.Get("/api/getawards", controller.GetAwards)
	app.Get("/api/getaward/:awardid", controller.GetAward)
	app.Post("/api/addaward", controller.AddAward)
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
	app.Post("/api/addcontent", controller.AddContent)
	app.Put("/api/updatecontent/:contentid", controller.UpdateContent)
	app.Delete("/api/deletecontent/:contentid", controller.DeleteContent)

	app.Get("/api/executives", controller.GetExecutives)
	app.Get("/api/executive/:executivesid", controller.GetExecutive)
	app.Post("/api/addexecutive", controller.AddExecutive)
	app.Put("/api/updateexecutive/:executivesid", controller.UpdateExecutive)
	app.Delete("/api/deleteexecutive/:executivesid", controller.DeleteExecutive)

	app.Get("/api/bannertops", controller.GetBannerTops)
	app.Get("/api/bannertop/:bannertopid", controller.GetBannerTop)
	app.Post("/api/addbannertop", controller.AddBannerTop)
	app.Put("/api/updatebannertop/:bannertopid", controller.UpdateBannerTop)
	app.Delete("/api/deletebannertop/:bannertopid", controller.DeleteBannerTop)

	app.Get("/api/getjobapplications", controller.GetJobApplications)
	app.Get("/api/getjobapplication/:jobapplicationsid", controller.GetJobApplication)
	app.Post("/api/addjobapplication", controller.AddJobApplication)
	app.Put("/api/updatejobapplication/:jobapplicationsid", controller.UpdateJobApplication)
	app.Delete("/api/deletejobapplication/:jobapplicationsid", controller.DeleteJobApplication)

	app.Get("/api/getusers", controller.GetUsers)
	app.Get("/api/getuser/:userid", controller.GetUser)
	app.Post("/api/adduser", controller.AddUser)
	app.Put("/api/updateuser/:userid", controller.UpdateUser)
	app.Delete("/api/deleteuser/:userid", controller.DeleteUser)

	app.Post("/api/updateconfig", controller.UpdateConfig)
	app.Get("/api/getconfig/", controller.GetConfig)

	app.Get("api/frontend/bannertopshtt", controller.FrontendGetBannerTops)
	app.Get("/api/frontend/getawardshtt", controller.FrontendGetAwards)
	app.Get("/api/frontend/getconfightt", controller.FrontendGetConfig)
	app.Get("/api/frontend/getcartypeshtt", controller.FrontendGetCarTypes)
	app.Get("/api/frontend/getcartypehtt/:cartypeid", controller.FrontendGetCarType)
	app.Get("/api/frontend/getcontents", controller.FrontendGetContents)
	app.Get("/api/frontend/getcontent/:contentid", controller.FrontendGetContent)
	app.Get("/api/frontend/getcars/:cartypeid", controller.FrontendGetCars)
	app.Get("/api/frontend/getjobshtt", controller.FrontendGetJobApplications)
	app.Get("/api/frontend/getjobhtt/:jobapplicationsid", controller.FrontendGetJobApplication)

	app.Post("/api/checkauth", controller.CheckAuth)

	app.Post("/api/addimages", controller.AddImage)
	app.Get("/api/getimages/:imagetype/:id", controller.GetImage)

	app.Listen(":8080")

}
