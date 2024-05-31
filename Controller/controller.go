package controller

import (
	services "HTTApi/Services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

// ## CarType
func GetCarTypes(c *fiber.Ctx) error {
	data, err, msg := services.GetCarTypes()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}
func GetCarTypesLst(c *fiber.Ctx) error {
	data, err, msg := services.GetCarTypesLst()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}
func GetCarType(c *fiber.Ctx) error {
	data, err, msg := services.GetCarType(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddCarType(c *fiber.Ctx) error {
	data, err, msg := services.AddCarType(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateCarType(c *fiber.Ctx) error {
	data, err, msg := services.UpdateCarType(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteCarType(c *fiber.Ctx) error {
	data, err, msg := services.DeleteCarType(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## CarType

// ## Award
func FrontendGetAwards(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetAwards()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetAwards(c *fiber.Ctx) error {
	data, err, msg := services.GetAwards()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetAward(c *fiber.Ctx) error {
	data, err, msg := services.GetAward(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddAward(c *fiber.Ctx) error {
	data, err, msg := services.AddAward(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateAward(c *fiber.Ctx) error {
	data, err, msg := services.UpdateAward(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteAward(c *fiber.Ctx) error {
	data, err, msg := services.DeleteAward(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## Award

// ## Car
func GetCars(c *fiber.Ctx) error {
	data, err, msg := services.GetCars()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetCar(c *fiber.Ctx) error {
	data, err, msg := services.GetCar(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddCar(c *fiber.Ctx) error {
	data, err, msg := services.AddCar(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateCar(c *fiber.Ctx) error {
	data, err, msg := services.UpdateCar(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteCar(c *fiber.Ctx) error {
	data, err, msg := services.DeleteCar(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## Car

// ## Content
func GetContents(c *fiber.Ctx) error {
	data, err, msg := services.GetContents()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetContent(c *fiber.Ctx) error {
	data, err, msg := services.GetContent(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddContent(c *fiber.Ctx) error {
	data, err, msg := services.AddContent(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateContent(c *fiber.Ctx) error {
	data, err, msg := services.UpdateContent(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteContent(c *fiber.Ctx) error {
	data, err, msg := services.DeleteContent(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## Content

// ## Executives
func GetExecutives(c *fiber.Ctx) error {
	data, err, msg := services.GetExecutives()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetExecutive(c *fiber.Ctx) error {
	data, err, msg := services.GetExecutive(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddExecutive(c *fiber.Ctx) error {
	data, err, msg := services.AddExecutive(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateExecutive(c *fiber.Ctx) error {
	data, err, msg := services.UpdateExecutive(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteExecutive(c *fiber.Ctx) error {
	data, err, msg := services.DeleteExecutive(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## Executives

// ## BannerTop
func FrontendGetBannerTops(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetBannerTops()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetBannerTops(c *fiber.Ctx) error {
	data, err, msg := services.GetBannerTops()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetBannerTop(c *fiber.Ctx) error {
	data, err, msg := services.GetBannerTop(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddBannerTop(c *fiber.Ctx) error {
	data, err, msg := services.AddBannerTop(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateBannerTop(c *fiber.Ctx) error {
	data, err, msg := services.UpdateBannerTop(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteBannerTop(c *fiber.Ctx) error {
	data, err, msg := services.DeleteBannerTop(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## BannerTop

// ## JobApplication
func GetJobApplications(c *fiber.Ctx) error {
	data, err, msg := services.GetJobApplications()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetJobApplication(c *fiber.Ctx) error {
	data, err, msg := services.GetJobApplication(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddJobApplication(c *fiber.Ctx) error {
	data, err, msg := services.AddJobApplication(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateJobApplication(c *fiber.Ctx) error {
	data, err, msg := services.UpdateJobApplication(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteJobApplication(c *fiber.Ctx) error {
	data, err, msg := services.DeleteJobApplication(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## JobApplication

// ## User
func GetUsers(c *fiber.Ctx) error {
	data, err, msg := services.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetUser(c *fiber.Ctx) error {
	data, err, msg := services.GetUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func AddUser(c *fiber.Ctx) error {
	data, err, msg := services.AddUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	data, err, msg := services.UpdateUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	data, err, msg := services.DeleteUser(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

// ## User

func AddImage(c *fiber.Ctx) error {
	err, msg := services.AddImage(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetImage(c *fiber.Ctx) error {
	data, err, msg := services.GetImage(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func UpdateConfig(c *fiber.Ctx) error {
	err, msg := services.UpdateConfig(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{

		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func GetConfig(c *fiber.Ctx) error {
	data, err, msg := services.GetConfig(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}

func FrontendGetConfig(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetConfig(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}
func FrontendGetCarTypes(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetCarTypes(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}
func FrontendGetCarType(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetCarType(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})
}
func FrontendGetContents(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetContents(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})

}
func FrontendGetContent(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetContent(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})

}
func FrontendGetCars(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetCars(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})

}

func FrontendGetJobApplications(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetJobApplications()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})

}

func FrontendGetJobApplication(c *fiber.Ctx) error {
	data, err, msg := services.FrontendGetJobApplication(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})

}

func CheckAuth(c *fiber.Ctx) error {
	data, err, msg := services.CheckAuth(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":    data,
		"status":  fiber.StatusOK,
		"message": msg,
	})

}
