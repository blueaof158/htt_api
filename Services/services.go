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

// ## CarType
func GetCarTypes() ([]model.CarType, error, string) {
	var cartypes []model.CarType
	rows, err := db.Query("SELECT CarTypeID,CarGUID,CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM CarType;")
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
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	var result model.CarType
	err = db.QueryRow("SELECT CarTypeID,CarGUID,CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM CarType WHERE CarTypeID = ?;", id).Scan(&result.CarTypeID, &result.CarGUID, &result.CarTypeName, &result.CarTypeDesctiption, &result.CarTypeInactive, &result.CreateBy, &result.CreateDate, &result.UpdateBy, &result.UpdateDate)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	return result, nil, "success"
}

func AddCarType(c *fiber.Ctx) (model.CarType, error, string) {
	cartype := new(model.CarType)
	if err = c.BodyParser(cartype); err != nil {
		return model.CarType{}, err, err.Error()
	}
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
	var r model.CarType
	r.CarTypeID = int(lastInsertID)
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

	data, err, _ := GetCarType(c)
	if data == (model.CarType{}) {
		return model.CarType{}, err, err.Error()
	}
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

// ## CarType

// ## Awards
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

func AddAward(c *fiber.Ctx) (model.Award, error, string) {
	award := new(model.Award)
	if err = c.BodyParser(award); err != nil {
		return model.Award{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO Award (AwardGUID,AwardName,AwardDescription,AwardInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Award{}, err, err.Error()
	}
	result, err := stmt.Exec(
		award.AwardGUID,
		award.AwardName,
		award.AwardDescription,
		award.AwardInactive,
	)
	if err != nil {
		return model.Award{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.Award{}, err, "can't get id"
	}
	var r model.Award
	r.AwardID = int(lastInsertID)
	r.AwardGUID = award.AwardGUID
	r.AwardName = award.AwardName
	r.AwardDescription = award.AwardDescription
	r.AwardInactive = award.AwardInactive
	return r, nil, "success"
}

func UpdateAward(c *fiber.Ctx) (model.Award, error, string) {
	awardid, err := strconv.Atoi(c.Params("awardid"))
	award := new(model.Award)
	c.BodyParser(award)

	data, err, _ := GetAward(c)
	if data == (model.Award{}) {
		return model.Award{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE Award SET AwardName=?, AwardDescription=?, AwardInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE AwardID=?")
	if err != nil {
		return model.Award{}, err, err.Error()
	}
	_, err = stmt.Exec(
		award.AwardName,
		award.AwardDescription,
		award.AwardInactive,
		awardid,
	)
	if err != nil {
		return model.Award{}, err, err.Error()
	}

	var r model.Award
	r.AwardID = awardid
	r.AwardGUID = award.AwardGUID
	r.AwardName = award.AwardName
	r.AwardDescription = award.AwardDescription
	r.AwardInactive = award.AwardInactive
	return r, nil, "success"
}

func DeleteAward(c *fiber.Ctx) (model.CarType, error, string) {
	awardid, err := strconv.Atoi(c.Params("awardid"))
	award := new(model.CarType)
	c.BodyParser(award)
	stmt, err := db.Prepare("DELETE FROM  Award WHERE AwardID=?")
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	_, err = stmt.Exec(
		awardid,
	)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	return model.CarType{}, nil, "success"
}

// ## Awards

// ## Car
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

func AddCar(c *fiber.Ctx) (model.Car, error, string) {
	car := new(model.Car)
	if err = c.BodyParser(car); err != nil {
		return model.Car{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO Car (CarGUID,CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	result, err := stmt.Exec(
		car.CarGUID,
		car.CarName,
		car.CarDesctiption,
		car.CarInactive,
	)
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.Car{}, err, "can't get id"
	}
	var r model.Car
	r.CarID = int(lastInsertID)
	r.CarGUID = car.CarGUID
	r.CarName = car.CarName
	r.CarTypeID = car.CarTypeID
	r.CarDesctiption = car.CarDesctiption
	r.CarInactive = car.CarInactive
	return r, nil, "success"
}

func UpdateCar(c *fiber.Ctx) (model.Car, error, string) {
	carid, err := strconv.Atoi(c.Params("carid"))
	car := new(model.Car)
	c.BodyParser(car)

	data, err, _ := GetCar(c)
	if data == (model.Car{}) {
		return model.Car{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE Car SET CarName=?, CarDesctiption=?, CarTypeID=?,CarInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE CarID=?")
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	_, err = stmt.Exec(
		car.CarName,
		car.CarDesctiption,
		car.CarTypeID,
		car.CarInactive,
		carid,
	)
	if err != nil {
		return model.Car{}, err, err.Error()
	}

	var r model.Car
	r.CarID = carid
	r.CarGUID = car.CarGUID
	r.CarName = car.CarName
	r.CarDesctiption = car.CarDesctiption
	r.CarTypeID = car.CarTypeID
	r.CarInactive = car.CarInactive
	return r, nil, "success"
}

func DeleteCar(c *fiber.Ctx) (model.Car, error, string) {
	carid, err := strconv.Atoi(c.Params("carid"))
	car := new(model.Car)
	c.BodyParser(car)
	stmt, err := db.Prepare("DELETE FROM Car WHERE CarID=?")
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	_, err = stmt.Exec(
		carid,
	)
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	return model.Car{}, nil, "success"
}

// ## Car

// ## Content

func GetContents() ([]model.Content, error, string) {
	var contents []model.Content
	query := "SELECT ContentID,ContentGUID,ContentTitle,HyphenationTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Content;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var content model.Content
		err = rows.Scan(&content.ContentID, &content.ContentGUID, &content.ContentTitle, &content.HyphenationTitle, &content.ContentText, &content.Content, &content.ContentInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		contents = append(contents, content)
	}
	return contents, nil, "success"
}

func GetContent(c *fiber.Ctx) (model.Content, error, string) {
	id, err := strconv.Atoi(c.Params("carid"))
	var content model.Content
	err = db.QueryRow("SELECT ContentID,ContentGUID,ContentTitle,HyphenationTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Content ContentID = ?;", id).Scan(&content.ContentID, &content.ContentGUID, &content.ContentTitle, &content.HyphenationTitle, &content.ContentText, &content.Content, &content.ContentInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate)
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	return content, nil, "success"
}

func AddContent(c *fiber.Ctx) (model.Content, error, string) {
	content := new(model.Content)
	if err = c.BodyParser(content); err != nil {
		return model.Content{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO Content (ContentGUID,ContentTitle,HyphenationTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	result, err := stmt.Exec(
		content.ContentGUID,
		content.ContentTitle,
		content.HyphenationTitle,
		content.ContentText,
		content.Content,
		content.ContentInactive,
	)
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.Content{}, err, "can't get id"
	}
	var r model.Content
	r.ContentID = int(lastInsertID)
	r.ContentGUID = content.ContentGUID
	r.ContentTitle = content.ContentTitle
	r.HyphenationTitle = content.HyphenationTitle
	r.ContentText = content.ContentText
	r.Content = content.Content
	r.ContentInactive = content.ContentInactive
	return r, nil, "success"
}

func UpdateContent(c *fiber.Ctx) (model.Content, error, string) {
	contentid, err := strconv.Atoi(c.Params("contentid"))
	content := new(model.Content)
	c.BodyParser(content)

	data, err, _ := GetContent(c)
	if data == (model.Content{}) {
		return model.Content{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE Content SET ContentTitle=?, HyphenationTitle=?, ContentText=?,Content=?,ContentInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE ContentID=?")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	_, err = stmt.Exec(
		content.ContentTitle,
		content.HyphenationTitle,
		content.ContentText,
		content.Content,
		content.ContentInactive,
		contentid,
	)
	if err != nil {
		return model.Content{}, err, err.Error()
	}

	var r model.Content
	r.ContentID = contentid
	r.ContentTitle = content.ContentTitle
	r.HyphenationTitle = content.HyphenationTitle
	r.ContentText = content.ContentText
	r.Content = content.Content
	r.ContentInactive = content.ContentInactive
	return r, nil, "success"
}

func DeleteContent(c *fiber.Ctx) (model.Content, error, string) {
	contentid, err := strconv.Atoi(c.Params("contentid"))
	content := new(model.Content)
	c.BodyParser(content)
	stmt, err := db.Prepare("DELETE FROM Content WHERE ContentID=?")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	_, err = stmt.Exec(
		contentid,
	)
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	return model.Content{}, nil, "success"
}

// ## Content

// ## Executives

func GetExecutives() ([]model.Executives, error, string) {
	var contents []model.Executives
	query := "SELECT ExecutivesID,ExecutivesGUID,ExecutivesFirstName,ExecutivesLastName,ExecutivesPosition,ExecutivesBio,ExecutivesInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Executives;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var content model.Executives
		err = rows.Scan(&content.ExecutivesID, &content.ExecutivesGUID, &content.ExecutivesFirstName, &content.ExecutivesLastName, &content.ExecutivesPosition, &content.ExecutivesBio, &content.ExecutivesInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		contents = append(contents, content)
	}
	return contents, nil, "success"
}

func GetExecutive(c *fiber.Ctx) (model.Executives, error, string) {
	id, err := strconv.Atoi(c.Params("carid"))
	var content model.Executives
	err = db.QueryRow("SELECT ExecutivesID,ExecutivesGUID,ExecutivesFirstName,ExecutivesLastName,ExecutivesPosition,ExecutivesBio,ExecutivesInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Executives WHERE ExecutivesID = ?;", id).Scan(&content.ExecutivesID, &content.ExecutivesGUID, &content.ExecutivesFirstName, &content.ExecutivesLastName, &content.ExecutivesPosition, &content.ExecutivesBio, &content.ExecutivesInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate)
	if err != nil {
		return model.Executives{}, err, err.Error()
	}
	return content, nil, "success"
}

func AddContent(c *fiber.Ctx) (model.Content, error, string) {
	content := new(model.Content)
	if err = c.BodyParser(content); err != nil {
		return model.Content{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO Content (ContentGUID,ContentTitle,HyphenationTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	result, err := stmt.Exec(
		content.ContentGUID,
		content.ContentTitle,
		content.HyphenationTitle,
		content.ContentText,
		content.Content,
		content.ContentInactive,
	)
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.Content{}, err, "can't get id"
	}
	var r model.Content
	r.ContentID = int(lastInsertID)
	r.ContentGUID = content.ContentGUID
	r.ContentTitle = content.ContentTitle
	r.HyphenationTitle = content.HyphenationTitle
	r.ContentText = content.ContentText
	r.Content = content.Content
	r.ContentInactive = content.ContentInactive
	return r, nil, "success"
}

func UpdateContent(c *fiber.Ctx) (model.Content, error, string) {
	contentid, err := strconv.Atoi(c.Params("contentid"))
	content := new(model.Content)
	c.BodyParser(content)

	data, err, _ := GetContent(c)
	if data == (model.Content{}) {
		return model.Content{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE Content SET ContentTitle=?, HyphenationTitle=?, ContentText=?,Content=?,ContentInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE ContentID=?")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	_, err = stmt.Exec(
		content.ContentTitle,
		content.HyphenationTitle,
		content.ContentText,
		content.Content,
		content.ContentInactive,
		contentid,
	)
	if err != nil {
		return model.Content{}, err, err.Error()
	}

	var r model.Content
	r.ContentID = contentid
	r.ContentTitle = content.ContentTitle
	r.HyphenationTitle = content.HyphenationTitle
	r.ContentText = content.ContentText
	r.Content = content.Content
	r.ContentInactive = content.ContentInactive
	return r, nil, "success"
}

func DeleteContent(c *fiber.Ctx) (model.Content, error, string) {
	contentid, err := strconv.Atoi(c.Params("contentid"))
	content := new(model.Content)
	c.BodyParser(content)
	stmt, err := db.Prepare("DELETE FROM Content WHERE ContentID=?")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	_, err = stmt.Exec(
		contentid,
	)
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	return model.Content{}, nil, "success"
}

// ## Executives
