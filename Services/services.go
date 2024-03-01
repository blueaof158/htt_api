package services

import (
	database "HTTApi/Database"
	model "HTTApi/Model"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var db, err = database.ConnectDB()

// ## CarType
func GetCarTypes() ([]model.CarType, error, string) {
	var cartypes []model.CarType
	rows, err := db.Query("SELECT CarTypeID,CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM CarType;")
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var cartype model.CarType
		err = rows.Scan(&cartype.CarTypeID, &cartype.CarTypeName, &cartype.CarTypeDesctiption, &cartype.CarTypeInactive, &cartype.CreateBy, &cartype.CreateDate, &cartype.UpdateBy, &cartype.UpdateDate)
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
	err = db.QueryRow("SELECT CarTypeID,CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM CarType WHERE CarTypeID = ?;", id).Scan(&result.CarTypeID, &result.CarTypeName, &result.CarTypeDesctiption, &result.CarTypeInactive, &result.CreateBy, &result.CreateDate, &result.UpdateBy, &result.UpdateDate)
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
	stmt, err := db.Prepare("INSERT INTO CarType (CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?, ?, ?, ?, User(), NOW(), User(), NOW())")
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	result, err := stmt.Exec(
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
	query := "SELECT AwardID,AwardName,AwardDescription,AwardInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Award;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var award model.Award
		err = rows.Scan(&award.AwardID, &award.AwardName, &award.AwardDescription, &award.AwardInactive, &award.CreateBy, &award.CreateDate, &award.UpdateBy, &award.UpdateDate)
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
	err = db.QueryRow("SELECT AwardID,AwardName,AwardDescription,AwardInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Award WHERE AwardID = ?;", id).Scan(&award.AwardID, &award.AwardName, &award.AwardDescription, &award.AwardInactive, &award.CreateBy, &award.CreateDate, &award.UpdateBy, &award.UpdateDate)

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
	stmt, err := db.Prepare("INSERT INTO Award (AwardName,AwardDescription,AwardInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Award{}, err, err.Error()
	}
	result, err := stmt.Exec(
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
	query := "SELECT CarID,CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Car;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var car model.Car
		err = rows.Scan(&car.CarID, &car.CarName, &car.CarDesctiption, &car.CarTypeID, &car.CarInactive, &car.CreateBy, &car.CreateDate, &car.UpdateBy, &car.UpdateDate)
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
	err = db.QueryRow("SELECT CarID,CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Car CarID = ?;", id).Scan(&result.CarID, &result.CarName, &result.CarDesctiption, &result.CarTypeID, &result.CarInactive, &result.CreateBy, &result.CreateDate, &result.UpdateBy, &result.UpdateDate)
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
	stmt, err := db.Prepare("INSERT INTO Car (CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	result, err := stmt.Exec(
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
	query := "SELECT ContentID,ContentTitle,HyphenationTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Content;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var content model.Content
		err = rows.Scan(&content.ContentID, &content.ContentTitle, &content.HyphenationTitle, &content.ContentText, &content.Content, &content.ContentInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate)
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
	err = db.QueryRow("SELECT ContentID,ContentTitle,HyphenationTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Content ContentID = ?;", id).Scan(&content.ContentID, &content.ContentTitle, &content.HyphenationTitle, &content.ContentText, &content.Content, &content.ContentInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate)
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
	stmt, err := db.Prepare("INSERT INTO Content (ContentTitle,HyphenationTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	result, err := stmt.Exec(
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
	query := "SELECT ExecutivesID,ExecutivesFirstName,ExecutivesLastName,ExecutivesPosition,ExecutivesBio,ExecutivesInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Executives;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var executive model.Executives
		err = rows.Scan(&executive.ExecutivesID, &executive.ExecutivesFirstName, &executive.ExecutivesLastName, &executive.ExecutivesPosition, &executive.ExecutivesBio, &executive.ExecutivesInactive, &executive.CreateBy, &executive.CreateDate, &executive.UpdateBy, &executive.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		contents = append(contents, executive)
	}
	return contents, nil, "success"
}

func GetExecutive(c *fiber.Ctx) (model.Executives, error, string) {
	id, err := strconv.Atoi(c.Params("executivesid"))
	var executive model.Executives
	err = db.QueryRow("SELECT ExecutivesID,ExecutivesFirstName,ExecutivesLastName,ExecutivesPosition,ExecutivesBio,ExecutivesInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Executives WHERE ExecutivesID = ?;", id).Scan(&executive.ExecutivesID, &executive.ExecutivesFirstName, &executive.ExecutivesLastName, &executive.ExecutivesPosition, &executive.ExecutivesBio, &executive.ExecutivesInactive, &executive.CreateBy, &executive.CreateDate, &executive.UpdateBy, &executive.UpdateDate)
	if err != nil {
		return model.Executives{}, err, err.Error()
	}
	return executive, nil, "success"
}

func AddExecutive(c *fiber.Ctx) (model.Executives, error, string) {
	executive := new(model.Executives)
	if err = c.BodyParser(executive); err != nil {
		return model.Executives{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO Executives (ExecutivesFirstName,ExecutivesLastName,ExecutivesPosition,ExecutivesBio,ExecutivesInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Executives{}, err, err.Error()
	}
	result, err := stmt.Exec(
		executive.ExecutivesFirstName,
		executive.ExecutivesLastName,
		executive.ExecutivesPosition,
		executive.ExecutivesBio,
		executive.ExecutivesInactive,
	)
	if err != nil {
		return model.Executives{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.Executives{}, err, "can't get id"
	}
	var r model.Executives
	r.ExecutivesID = int(lastInsertID)
	r.ExecutivesFirstName = executive.ExecutivesFirstName
	r.ExecutivesLastName = executive.ExecutivesLastName
	r.ExecutivesPosition = executive.ExecutivesPosition
	r.ExecutivesBio = executive.ExecutivesBio
	r.ExecutivesInactive = executive.ExecutivesInactive
	return r, nil, "success"
}

func UpdateExecutive(c *fiber.Ctx) (model.Executives, error, string) {
	executivesid, err := strconv.Atoi(c.Params("executivesid"))
	content := new(model.Executives)
	c.BodyParser(content)

	data, err, _ := GetContent(c)
	if data == (model.Content{}) {
		return model.Executives{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE Executives SET ExecutivesFirstName=?, ExecutivesLastName=?, ExecutivesPosition=?,ExecutivesBio=?,ExecutivesInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE ExecutivesID=?")
	if err != nil {
		return model.Executives{}, err, err.Error()
	}
	_, err = stmt.Exec(
		content.ExecutivesFirstName,
		content.ExecutivesLastName,
		content.ExecutivesPosition,
		content.ExecutivesBio,
		content.ExecutivesInactive,
		executivesid,
	)
	if err != nil {
		return model.Executives{}, err, err.Error()
	}

	var r model.Executives
	r.ExecutivesID = executivesid
	r.ExecutivesFirstName = content.ExecutivesFirstName
	r.ExecutivesLastName = content.ExecutivesLastName
	r.ExecutivesPosition = content.ExecutivesPosition
	r.ExecutivesBio = content.ExecutivesBio
	r.ExecutivesInactive = content.ExecutivesInactive
	return r, nil, "success"
}

func DeleteExecutive(c *fiber.Ctx) (model.Executives, error, string) {
	executivesid, err := strconv.Atoi(c.Params("executivesid"))
	executiv := new(model.Executives)
	c.BodyParser(executiv)
	stmt, err := db.Prepare("DELETE FROM Executives WHERE ExecutivesID=?")
	if err != nil {
		return model.Executives{}, err, err.Error()
	}
	_, err = stmt.Exec(
		executivesid,
	)
	if err != nil {
		return model.Executives{}, err, err.Error()
	}
	return model.Executives{}, nil, "success"
}

// ## Executives

// ## BannerTop
func GetBannerTops() ([]model.BannerTop, error, string) {
	var bannertops []model.BannerTop
	query := "SELECT BannerTopID,BannerTopImagegPath,BannerTopImageLink,BannerTopInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Executives;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	for rows.Next() {
		var bannertop model.BannerTop
		err = rows.Scan(&bannertop.BannerTopID, &bannertop.BannerTopImagegPath, &bannertop.BannerTopImageLink, &bannertop.BannerTopInactive, &bannertop.CreateBy, &bannertop.CreateDate, &bannertop.UpdateBy, &bannertop.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		bannertops = append(bannertops, bannertop)
	}
	return bannertops, nil, "success"
}

func GetBannerTop(c *fiber.Ctx) (model.BannerTop, error, string) {
	id, err := strconv.Atoi(c.Params("bannertopid"))
	var bannertop model.BannerTop
	err = db.QueryRow("SELECT BannerTopID,BannerTopImagegPath,BannerTopImageLink,BannerTopInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM BannerTop WHERE ExecutivesID = ?;", id).Scan(&bannertop.BannerTopID, &bannertop.BannerTopImagegPath, &bannertop.BannerTopImageLink, &bannertop.BannerTopInactive, &bannertop.CreateBy, &bannertop.CreateDate, &bannertop.UpdateBy, &bannertop.UpdateDate)
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	return bannertop, nil, "success"
}

func AddBannerTop(c *fiber.Ctx) (model.BannerTop, error, string) {
	bannertop := new(model.BannerTop)
	if err = c.BodyParser(bannertop); err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO BannerTop (BannerTopImagegPath,BannerTopImageLink,BannerTopInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	result, err := stmt.Exec(
		bannertop.BannerTopImagegPath,
		bannertop.BannerTopImageLink,
		bannertop.BannerTopInactive,
	)
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.BannerTop{}, err, "can't get id"
	}
	var r model.BannerTop
	r.BannerTopID = int(lastInsertID)
	r.BannerTopImagegPath = bannertop.BannerTopImagegPath
	r.BannerTopImageLink = bannertop.BannerTopImageLink
	r.BannerTopInactive = bannertop.BannerTopInactive
	return r, nil, "success"
}

func UpdateBannerTop(c *fiber.Ctx) (model.BannerTop, error, string) {
	bannertopid, err := strconv.Atoi(c.Params("bannertopid"))
	content := new(model.BannerTop)
	c.BodyParser(content)

	data, err, _ := GetContent(c)
	if data == (model.Content{}) {
		return model.BannerTop{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE BannerTop SET BannerTopImagegPath=?, BannerTopImageLink=?, BannerTopInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE BannerTopID=?")
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	_, err = stmt.Exec(
		content.BannerTopImagegPath,
		content.BannerTopImageLink,
		content.BannerTopInactive,
		bannertopid,
	)
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}

	var r model.BannerTop
	r.BannerTopID = bannertopid
	r.BannerTopImagegPath = content.BannerTopImagegPath
	r.BannerTopImageLink = content.BannerTopImageLink
	r.BannerTopInactive = content.BannerTopInactive
	return r, nil, "success"
}

func DeleteBannerTop(c *fiber.Ctx) (model.BannerTop, error, string) {
	bannertopid, err := strconv.Atoi(c.Params("bannertopid"))
	executiv := new(model.Executives)
	c.BodyParser(executiv)
	stmt, err := db.Prepare("DELETE FROM BannerTop WHERE BannerTopID=?")
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	_, err = stmt.Exec(
		bannertopid,
	)
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	return model.BannerTop{}, nil, "success"
}

// ## BannerTop

// ## JobApplications

// func GetJobApplications() ([]model.JobApplications, error, string) {
// 	var jobApplications []model.JobApplications
// 	query := "SELECT JobApplicationsID,JobApplicationsName,JobApplicationsPosition,JobApplicationsDescription,JobApplicationsInactive,JobApplicationFilePath,CreateBy,CreateDate,UpdateBy,UpdateDate FROM JobApplications;"
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, err, err.Error()
// 	}
// 	for rows.Next() {
// 		var jobApplication model.JobApplications
// 		err = rows.Scan(&jobApplication.JobApplicationsID, &jobApplication.JobApplicationsName, &jobApplication.JobApplicationsPosition, &jobApplication.JobApplicationsDescription, &jobApplication.JobApplicationsInactive, &jobApplication.JobApplicationFilePath, &jobApplication.CreateBy, &jobApplication.CreateDate, &jobApplication.UpdateBy, &jobApplication.UpdateDate)
// 		if err != nil {
// 			return nil, err, err.Error()
// 		}
// 		jobApplications = append(jobApplications, jobApplication)
// 	}
// 	return jobApplications, nil, "success"
// }

// func GetJobApplication(c *fiber.Ctx) (model.JobApplications, error, string) {
// 	id, err := strconv.Atoi(c.Params("jobApplicationsid"))
// 	var jobApplication model.JobApplications
// 	err = db.QueryRow("SELECT JobApplicationsID,JobApplicationsName,JobApplicationsPosition,JobApplicationsDescription,JobApplicationsInactive,JobApplicationFilePath,CreateBy,CreateDate,UpdateBy,UpdateDate FROM JobApplications WHERE JobApplicationsID = ?;", id).Scan(&jobApplication.JobApplicationsID, &jobApplication.JobApplicationsName, &jobApplication.JobApplicationsPosition, &jobApplication.JobApplicationsDescription, &jobApplication.JobApplicationsInactive, &jobApplication.JobApplicationFilePath, &jobApplication.CreateBy, &jobApplication.CreateDate, &jobApplication.UpdateBy, &jobApplication.UpdateDate)
// 	if err != nil {
// 		return model.JobApplications{}, err, err.Error()
// 	}
// 	return jobApplication, nil, "success"
// }

// func AddJobApplication(c *fiber.Ctx) (model.JobApplications, error, string) {
// 	jobApplication := new(model.JobApplications)
// 	if err = c.BodyParser(jobApplication); err != nil {
// 		return model.JobApplications{}, err, err.Error()
// 	}
// 	stmt, err := db.Prepare("INSERT INTO JobApplications (JobApplicationsName,JobApplicationsPosition,JobApplicationsDescription,JobApplicationsInactive,JobApplicationFilePath,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,?,User(), NOW(), User(), NOW())")
// 	if err != nil {
// 		return model.JobApplications{}, err, err.Error()
// 	}
// 	result, err := stmt.Exec(
// 		jobApplication.JobApplicationsName,
// 		jobApplication.ExecutivesLastName,
// 		jobApplication.ExecutivesPosition,
// 		jobApplication.ExecutivesBio,
// 		jobApplication.ExecutivesInactive,
// 	)
// 	if err != nil {
// 		return model.JobApplications{}, err, err.Error()
// 	}
// 	lastInsertID, err := result.LastInsertId()
// 	if err != nil {
// 		return model.JobApplications{}, err, "can't get id"
// 	}
// 	var r model.JobApplications
// 	r.ExecutivesID = int(lastInsertID)
// 	r.ExecutivesFirstName = executive.ExecutivesFirstName
// 	r.ExecutivesLastName = executive.ExecutivesLastName
// 	r.ExecutivesPosition = executive.ExecutivesPosition
// 	r.ExecutivesBio = executive.ExecutivesBio
// 	r.ExecutivesInactive = executive.ExecutivesInactive
// 	return r, nil, "success"
// }

// func UpdateJobApplication(c *fiber.Ctx) (model.JobApplications, error, string) {
// 	jobapplicationsid, err := strconv.Atoi(c.Params("jobapplicationsid"))
// 	content := new(model.JobApplications)
// 	c.BodyParser(content)

// 	data, err, _ := GetContent(c)
// 	if data == (model.JobApplications{}) {
// 		return model.JobApplications{}, err, err.Error()
// 	}
// 	stmt, err := db.Prepare("UPDATE Executives SET ExecutivesFirstName=?, ExecutivesLastName=?, ExecutivesPosition=?,ExecutivesBio=?,ExecutivesInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE ExecutivesID=?")
// 	if err != nil {
// 		return model.JobApplications{}, err, err.Error()
// 	}
// 	_, err = stmt.Exec(
// 		content.ExecutivesFirstName,
// 		content.ExecutivesLastName,
// 		content.ExecutivesPosition,
// 		content.ExecutivesBio,
// 		content.ExecutivesInactive,
// 		executivesid,
// 	)
// 	if err != nil {
// 		return model.Executives{}, err, err.Error()
// 	}

// 	var r model.Executives
// 	r.ExecutivesID = executivesid
// 	r.ExecutivesFirstName = content.ExecutivesFirstName
// 	r.ExecutivesLastName = content.ExecutivesLastName
// 	r.ExecutivesPosition = content.ExecutivesPosition
// 	r.ExecutivesBio = content.ExecutivesBio
// 	r.ExecutivesInactive = content.ExecutivesInactive
// 	return r, nil, "success"
// }

// func DeleteExecutive(c *fiber.Ctx) (model.Executives, error, string) {
// 	executivesid, err := strconv.Atoi(c.Params("executivesid"))
// 	executiv := new(model.Executives)
// 	c.BodyParser(executiv)
// 	stmt, err := db.Prepare("DELETE FROM Executives WHERE ExecutivesID=?")
// 	if err != nil {
// 		return model.Executives{}, err, err.Error()
// 	}
// 	_, err = stmt.Exec(
// 		executivesid,
// 	)
// 	if err != nil {
// 		return model.Executives{}, err, err.Error()
// 	}
// 	return model.Executives{}, nil, "success"
// }

// ## JobApplications
