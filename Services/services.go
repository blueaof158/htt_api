package services

import (
	database "HTTApi/Database"
	model "HTTApi/Model"
	"strconv"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var db, err = database.ConnectDB()

func main() {
	if err != nil {
		log.Fatal("Cant Connect Database")
	}
}

func GetCarTypes() ([]model.CarType, error, string) {
	var cartypes []model.CarType
	query := "SELECT CarTypeID,CarGUID,CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM CarType;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var cartype model.CarType
		err = rows.Scan(&cartype.CarTypeID, &cartype.CarGUID, &cartype.CarTypeName, &cartype.CarTypeDesctiption, &cartype.CarTypeInactive, &cartype.CreateBy, &cartype.CreateDate, &cartype.UpdateBy, &cartype.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		cartypes = append(cartypes, cartype)
	}
	return cartypes, nil, "success"
}

func GetCarType(c *fiber.Ctx) (model.CarType, error, string) {
	id, err := strconv.Atoi(c.Params("cartypeid"))
	var result model.CarType
	err = db.QueryRow("SELECT CarTypeID,CarGUID,CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM CarType WHERE CarTypeID = ?;", id).Scan(&result.CarTypeID, &result.CarGUID, &result.CarTypeName, &result.CarTypeDesctiption, &result.CarTypeInactive, &result.CreateBy, &result.CreateDate, &result.UpdateBy, &result.UpdateDate)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	return result, nil, "success"
}

func GetAwards() ([]model.Award, error, string) {
	var awards []model.Award
	query := "SELECT AwardID,AwardGUID,AwardName,AwardDescription,AwardInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Award;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var award model.Award
		err = rows.Scan(&award.AwardID, &award.AwardGUID, &award.AwardName, &award.AwardDescription, &award.AwardInactive, &award.CreateBy, &award.CreateDate, &award.UpdateBy, &award.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		awards = append(awards, award)
	}
	return awards, nil, "success"
}

func GetAward(c *fiber.Ctx) (model.Award, error, string) {
	id, err := strconv.Atoi(c.Params("awardid"))
	var award model.Award
	err = db.QueryRow("SELECT AwardID,AwardGUID,AwardName,AwardDescription,AwardInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Award WHERE AwardID = ?;", id).Scan(&award.AwardID, &award.AwardGUID, &award.AwardName, &award.AwardDescription, &award.AwardInactive, &award.CreateBy, &award.CreateDate, &award.UpdateBy, &award.UpdateDate)

	if err != nil {
		return model.Award{}, err, err.Error()
	}

	return award, nil, "success"
}

func GetCars() ([]model.Car, error, string) {
	var cars []model.Car
	query := "SELECT CarID,CarGUID,CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Car;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var car model.Car
		err = rows.Scan(&car.CarID, &car.CarGUID, &car.CarName, &car.CarDesctiption, &car.CarTypeID, &car.CarInactive, &car.CreateBy, &car.CreateDate, &car.UpdateBy, &car.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		cars = append(cars, car)
	}
	return cars, nil, "success"
}

func GetCar(c *fiber.Ctx) (model.Car, error, string) {
	id, err := strconv.Atoi(c.Params("carid"))
	var result model.Car
	err = db.QueryRow("SELECT CarID,CarGUID,CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Car CarID = ?;", id).Scan(&result.CarID, &result.CarGUID, &result.CarName, &result.CarDesctiption, &result.CarTypeID, &result.CarInactive, &result.CreateBy, &result.CreateDate, &result.UpdateBy, &result.UpdateDate)
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	return result, nil, "success"
}
