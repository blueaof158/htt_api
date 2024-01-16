package controller

import (
	services "HTTApi/Services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

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
