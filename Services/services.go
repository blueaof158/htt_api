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

func AddCarType(c *fiber.Ctx) (model.CarType, error, string) {
	cartype := new(model.CarType)
	c.BodyParser(cartype)
	stmt, err := db.Prepare("INSERT INTO CarType (CarGUID,CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?, ?, ?, ?, User(), NOW(), User(), NOW())")
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	result, err := stmt.Exec(
		cartype.CarGUID,
		cartype.CarTypeName,
		cartype.CarTypeDesctiption,
		cartype.CarTypeInactive,
	)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.CarType{}, err, "can't get id"
	}
	cartype.CarTypeID = int(lastInsertID)
	var r model.CarType
	r.CarGUID = cartype.CarGUID
	r.CarTypeName = cartype.CarTypeName
	r.CarTypeDesctiption = cartype.CarTypeDesctiption
	r.CarTypeInactive = cartype.CarTypeInactive
	return r, nil, "success"
}

func UpdateCarType(c *fiber.Ctx) (model.CarType, error, string) {
	cartypeid, err := strconv.Atoi(c.Params("cartypeid"))
	cartype := new(model.CarType)
	c.BodyParser(cartype)
	stmt, err := db.Prepare("UPDATE CarType SET CarTypeName=?, CarTypeDesctiption=?, CarTypeInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE CarTypeID=?")
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	_, err = stmt.Exec(
		cartype.CarTypeName,
		cartype.CarTypeDesctiption,
		cartype.CarTypeInactive,
		cartypeid,
	)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}

	var r model.CarType
	r.CarTypeID = cartypeid
	r.CarGUID = cartype.CarGUID
	r.CarTypeName = cartype.CarTypeName
	r.CarTypeDesctiption = cartype.CarTypeDesctiption
	r.CarTypeInactive = cartype.CarTypeInactive
	return r, nil, "success"
}

func DeleteCarType(c *fiber.Ctx) (model.CarType, error, string) {
	cartypeid, err := strconv.Atoi(c.Params("cartypeid"))
	cartype := new(model.CarType)
	c.BodyParser(cartype)
	stmt, err := db.Prepare("DELETE FROM  CarType WHERE CarTypeID=?")
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	_, err = stmt.Exec(
		cartypeid,
	)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}

	return model.CarType{}, nil, "success"
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
