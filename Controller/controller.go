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

// func AddContent(c *fiber.Ctx) error {
// 	data, err, msg := services.AddContent(c)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":  fiber.StatusBadRequest,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"data":    data,
// 		"status":  fiber.StatusOK,
// 		"message": msg,
// 	})
// }

// func UpdateContent(c *fiber.Ctx) error {
// 	data, err, msg := services.UpdateContent(c)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":  fiber.StatusBadRequest,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"data":    data,
// 		"status":  fiber.StatusOK,
// 		"message": msg,
// 	})
// }

// func DeleteContent(c *fiber.Ctx) error {
// 	data, err, msg := services.DeleteContent(c)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"status":  fiber.StatusBadRequest,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.JSON(fiber.Map{
// 		"data":    data,
// 		"status":  fiber.StatusOK,
// 		"message": msg,
// 	})
// }
// ## Executives
