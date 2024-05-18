package services

import (
	constant "HTTApi/Constant"
	database "HTTApi/Database"
	global "HTTApi/Global"
	model "HTTApi/Model"
	"database/sql"
	"reflect"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var db, err = database.ConnectDB()

// ## CarType
func GetCarTypes() ([]model.CarType, error, string) {
	var cartypes []model.CarType
	rows, err := db.Query("SELECT car.CarTypeID, car.CarTypeName, car.CarTypeDesctiption, car.CarTypeInactive, img.ImagePath, car.CreateBy, car.CreateDate, car.UpdateBy, car.UpdateDate FROM CarType car LEFT JOIN Images img ON car.CarTypeID = img.CarTypeID")
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var cartype model.CarType
		var imagePath sql.NullString
		err := rows.Scan(&cartype.CarTypeID, &cartype.CarTypeName, &cartype.CarTypeDesctiption, &cartype.CarTypeInactive, &imagePath, &cartype.CreateBy, &cartype.CreateDate, &cartype.UpdateBy, &cartype.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		if imagePath.Valid {
			cartype.Image64 = imagePath.String
		} else {
			cartype.Image64 = ""
		}
		cartypes = append(cartypes, cartype)
	}
	if err := rows.Err(); err != nil {
		return nil, err, err.Error()
	}
	return cartypes, nil, "success"
}

func GetCarTypesLst() ([]model.CarTypeLst, error, string) {
	var cartypeslst []model.CarTypeLst
	rows, err := db.Query("SELECT CarTypeID, CarTypeName FROM CarType")
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var cartype model.CarTypeLst
		err := rows.Scan(&cartype.CarTypeID, &cartype.CarTypeName)
		if err != nil {
			return nil, err, err.Error()
		}
		cartypeslst = append(cartypeslst, cartype)
	}
	if err := rows.Err(); err != nil {
		return nil, err, err.Error()
	}
	return cartypeslst, nil, "success"
}

func GetCarType(c *fiber.Ctx) (model.CarType, error, string) {
	id, err := strconv.Atoi(c.Params("cartypeid"))
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	var cartype model.CarType
	err = db.QueryRow("SELECT car.CarTypeID,car.CarTypeName,car.CarTypeDesctiption,car.CarTypeInactive,img.ImagePath,car.CreateBy,car.CreateDate,car.UpdateBy,car.UpdateDate FROM CarType car LEFT JOIN Images img ON car.CarTypeID = img.CarTypeID WHERE car.CarTypeID = ?;", id).Scan(&cartype.CarTypeID, &cartype.CarTypeName, &cartype.CarTypeDesctiption, &cartype.CarTypeInactive, &cartype.Image64, &cartype.CreateBy, &cartype.CreateDate, &cartype.UpdateBy, &cartype.UpdateDate)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	return cartype, nil, "success"
}

func AddCarType(c *fiber.Ctx) (model.CarType, error, string) {
	cartype := new(model.CarType)
	if err = c.BodyParser(cartype); err != nil {
		return model.CarType{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO CarType (CarTypeName,CarTypeDesctiption,CarTypeInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?, ?, ?, User(), NOW(), User(), NOW())")
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
	lastInsertID, _ := result.LastInsertId()
	err = global.InsertImage(cartype.Image64, lastInsertID, constant.CarTypeImage)
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

	err = global.DeleteImage(int64(cartypeid), constant.CarTypeImage)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}
	err = global.InsertImage(cartype.Image64, int64(cartypeid), constant.CarTypeImage)
	if err != nil {
		return model.CarType{}, err, err.Error()
	}

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
func FrontendGetAwards() ([]model.FrontendAward, error, string) {
	var awards []model.FrontendAward
	query := "SELECT award.AwardName,award.AwardDescription,img.ImagePath FROM Award award LEFT JOIN Images img ON award.AwardID = img.AwardID WHERE award.AwardInactive != 1;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var award model.FrontendAward
		err = rows.Scan(&award.AwardName, &award.AwardDescription, &award.Image64)
		if err != nil {
			return nil, err, err.Error()
		}
		awards = append(awards, award)
	}
	return awards, nil, "success"
}

func GetAwards() ([]model.Award, error, string) {
	var awards []model.Award
	query := "SELECT award.*,img.ImagePath FROM Award award LEFT JOIN Images img ON award.AwardID = img.AwardID;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var award model.Award
		err = rows.Scan(&award.AwardID, &award.AwardName, &award.AwardDescription, &award.AwardInactive, &award.CreateBy, &award.CreateDate, &award.UpdateBy, &award.UpdateDate, &award.Image64)
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
	err = db.QueryRow("SELECT award.*,img.ImagePath FROM Award award LEFT JOIN Images img ON award.AwardID = img.AwardID WHERE award.AwardID = ? ORDER BY img.ImagesID LIMIT 1;", id).Scan(&award.AwardID, &award.AwardName, &award.AwardDescription, &award.AwardInactive, &award.CreateBy, &award.CreateDate, &award.UpdateBy, &award.UpdateDate, &award.Image64)

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
	stmt, err := db.Prepare("INSERT INTO Award (AwardName,AwardDescription,AwardInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,User(), NOW(), User(), NOW())")
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
	global.InsertImage(award.Image64, lastInsertID, constant.AwardImage)
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
	global.DeleteImage(int64(awardid), constant.AwardImage)
	global.InsertImage(award.Image64, int64(awardid), constant.AwardImage)
	var r model.Award
	r.AwardID = awardid
	r.AwardName = award.AwardName
	r.AwardDescription = award.AwardDescription
	r.AwardInactive = award.AwardInactive
	return r, nil, "success"
}

func DeleteAward(c *fiber.Ctx) (model.CarType, error, string) {
	awardid, err := strconv.Atoi(c.Params("awardid"))
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
	// query := "SELECT CarID,CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM Car;"
	query := "SELECT car.*,img.ImagePath,cartype.CarTypeName FROM Car car LEFT JOIN (SELECT * FROM Images img WHERE img.ImagesID = ( SELECT MIN(inner_img.ImagesID) FROM Images inner_img WHERE inner_img.CarID = img.CarID )) img ON car.CarID = img.CarID LEFT JOIN CarType cartype ON car.CarTypeID = cartype.CarTypeID ORDER BY img.ImagesID;"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var car model.Car
		err = rows.Scan(&car.CarID, &car.CarName, &car.CarDesctiption, &car.CarTypeID, &car.CarInactive, &car.CreateBy, &car.CreateDate, &car.UpdateBy, &car.UpdateDate, &car.Image64, &car.CarTypeName)
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
	err = db.QueryRow("SELECT car.*,img.ImagePath,cartype.CarTypeName FROM Car car LEFT JOIN Images img ON car.CarID = img.CarID LEFT JOIN CarType cartype ON car.CarTypeID = cartype.CarTypeID WHERE car.CarID = ? ORDER BY img.ImagesID  LIMIT 1;", id).Scan(&result.CarID, &result.CarName, &result.CarDesctiption, &result.CarTypeID, &result.CarInactive, &result.CreateBy, &result.CreateDate, &result.UpdateBy, &result.UpdateDate, &result.Image64, &result.CarTypeName)
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
	stmt, err := db.Prepare("INSERT INTO Car (CarName,CarDesctiption,CarTypeID,CarInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Car{}, err, err.Error()
	}
	result, err := stmt.Exec(
		car.CarName,
		car.CarDesctiption,
		car.CarTypeID,
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
	err = global.DeleteImage(int64(carid), constant.CarImage)
	return model.Car{}, nil, "success"
}

// ## Car

// ## Content

func GetContents() ([]model.Content, error, string) {
	var contents []model.Content
	query := "SELECT c.* , i.ImagePath FROM Content c LEFT JOIN Images i ON c.ContentID = i.ContentID;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var content model.Content
		err = rows.Scan(&content.ContentID, &content.ContentTitle, &content.ContentText, &content.Content, &content.ContentInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate, &content.Image64)
		if err != nil {
			return nil, err, err.Error()
		}
		contents = append(contents, content)
	}
	return contents, nil, "success"
}

func GetContent(c *fiber.Ctx) (model.Content, error, string) {
	id, err := strconv.Atoi(c.Params("contentid"))
	var content model.Content
	err = db.QueryRow("SELECT c.* , i.ImagePath FROM Content c LEFT JOIN Images i ON c.ContentID = i.ContentID WHERE c.ContentID = ? LIMIT 1;", id).Scan(&content.ContentID, &content.ContentTitle, &content.ContentText, &content.Content, &content.ContentInactive, &content.CreateBy, &content.CreateDate, &content.UpdateBy, &content.UpdateDate, &content.Image64)
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
	stmt, err := db.Prepare("INSERT INTO Content (ContentTitle,ContentText,Content,ContentInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	result, err := stmt.Exec(
		content.ContentTitle,
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

	global.InsertImage(content.Image64, int64(lastInsertID), constant.ContentImage)
	var r model.Content
	r.ContentID = int(lastInsertID)
	r.ContentTitle = content.ContentTitle
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
	stmt, err := db.Prepare("UPDATE Content SET ContentTitle=?, ContentText=?,Content=?,ContentInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE ContentID=?")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	_, err = stmt.Exec(
		content.ContentTitle,
		content.ContentText,
		content.Content,
		content.ContentInactive,
		contentid,
	)
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	global.DeleteImage(int64(contentid), constant.ContentImage)
	global.InsertImage(content.Image64, int64(contentid), constant.ContentImage)

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
	stmt, err := db.Prepare("DELETE FROM Content WHERE ContentID=?")
	if err != nil {
		return model.Content{}, err, err.Error()
	}
	_, err = stmt.Exec(
		contentid,
	)
	global.DeleteImage(int64(contentid), constant.ContentImage)
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
	defer rows.Close()
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
	data, err, _ := GetExecutive(c)
	if data == (model.Executives{}) {
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
		executivesid)
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
func FrontendGetBannerTops() ([]model.FrontendBannerTop, error, string) {
	var bannertops []model.FrontendBannerTop
	query := "SELECT BT.BannerTopImageLink,Img.ImagePath FROM Htt.BannerTop BT LEFT JOIN Htt.Images Img ON BT.BannerTopID = Img.BannerTopID WHERE BT.BannerTopInactive = 0;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var bannertop model.FrontendBannerTop
		err = rows.Scan(&bannertop.BannerTopImageLink, &bannertop.ImagePath)
		if err != nil {
			return nil, err, err.Error()
		}
		bannertops = append(bannertops, bannertop)
	}
	return bannertops, nil, "success"
}

func GetBannerTops() ([]model.BannerTop, error, string) {
	var bannertops []model.BannerTop
	query := "SELECT BT.BannerTopID,BT.BannerTopImageLink,BT.BannerTopInactive,Img.ImagePath,BT.CreateBy,BT.CreateDate,BT.UpdateBy,BT.UpdateDate FROM Htt.BannerTop BT LEFT JOIN Htt.Images Img ON BT.BannerTopID = Img.BannerTopID ORDER BY BT.BannerTopID ASC;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var bannertop model.BannerTop
		err = rows.Scan(&bannertop.BannerTopID, &bannertop.BannerTopImageLink, &bannertop.BannerTopInactive, &bannertop.ImagePath, &bannertop.CreateBy, &bannertop.CreateDate, &bannertop.UpdateBy, &bannertop.UpdateDate)
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

	err = db.QueryRow("SELECT BT.BannerTopID,BT.BannerTopImageLink,BT.BannerTopInactive,Img.ImagePath,BT.CreateBy,BT.CreateDate,BT.UpdateBy,BT.UpdateDate FROM Htt.BannerTop BT LEFT JOIN Htt.Images Img ON BT.BannerTopID = Img.BannerTopID WHERE BT.BannerTopID = ?;", id).Scan(&bannertop.BannerTopID, &bannertop.BannerTopImageLink, &bannertop.BannerTopInactive, &bannertop.ImagePath, &bannertop.CreateBy, &bannertop.CreateDate, &bannertop.UpdateBy, &bannertop.UpdateDate)
	// err = db.QueryRow("SELECT BannerTopID,BannerTopImagegPath,BannerTopImageLink,BannerTopInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM BannerTop WHERE ExecutivesID = ?;", id).Scan(&bannertop.BannerTopID, &bannertop.BannerTopImagegPath, &bannertop.BannerTopImageLink, &bannertop.BannerTopInactive, &bannertop.CreateBy, &bannertop.CreateDate, &bannertop.UpdateBy, &bannertop.UpdateDate)
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
	stmt, err := db.Prepare("INSERT INTO BannerTop (BannerTopImageLink,BannerTopInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	result, err := stmt.Exec(
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
	global.InsertImage(bannertop.ImagePath, int64(lastInsertID), constant.BannerImage)
	var r model.BannerTop
	r.BannerTopID = int(lastInsertID)
	r.BannerTopImageLink = bannertop.BannerTopImageLink
	r.BannerTopInactive = bannertop.BannerTopInactive
	return r, nil, "success"
}

func UpdateBannerTop(c *fiber.Ctx) (model.BannerTop, error, string) {
	bannertopid, err := strconv.Atoi(c.Params("bannertopid"))
	content := new(model.BannerTop)
	c.BodyParser(content)
	data, err, _ := GetBannerTop(c)
	if data == (model.BannerTop{}) {
		return model.BannerTop{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE BannerTop SET BannerTopImageLink=?, BannerTopInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE BannerTopID=?")
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	_, err = stmt.Exec(
		content.BannerTopImageLink,
		content.BannerTopInactive,
		bannertopid,
	)

	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	global.DeleteImage(int64(bannertopid), constant.BannerImage)
	global.InsertImage(content.ImagePath, int64(bannertopid), constant.BannerImage)

	var r model.BannerTop
	r.BannerTopID = bannertopid
	r.BannerTopImageLink = content.BannerTopImageLink
	r.BannerTopInactive = content.BannerTopInactive
	return r, nil, "success"
}

func DeleteBannerTop(c *fiber.Ctx) (model.BannerTop, error, string) {
	bannertopid, err := strconv.Atoi(c.Params("bannertopid"))
	stmt, err := db.Prepare("DELETE FROM BannerTop WHERE BannerTopID=?")
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	_, err = stmt.Exec(
		bannertopid,
	)
	global.DeleteImage(int64(bannertopid), constant.BannerImage)
	if err != nil {
		return model.BannerTop{}, err, err.Error()
	}
	return model.BannerTop{}, nil, "success"
}

// ## BannerTop

// ## JobApplications

func GetJobApplications() ([]model.JobApplications, error, string) {
	var jobApplications []model.JobApplications
	query := "SELECT JobApplicationsID,JobApplicationsName,JobApplicationsPosition,JobApplicationsInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM JobApplications;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var jobApplication model.JobApplications
		err = rows.Scan(&jobApplication.JobApplicationsID, &jobApplication.JobApplicationsName, &jobApplication.JobApplicationsPosition, &jobApplication.JobApplicationsInactive, &jobApplication.CreateBy, &jobApplication.CreateDate, &jobApplication.UpdateBy, &jobApplication.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		jobApplications = append(jobApplications, jobApplication)
	}
	return jobApplications, nil, "success"
}

func GetJobApplication(c *fiber.Ctx) (model.JobApplications, error, string) {
	id, err := strconv.Atoi(c.Params("jobApplicationsid"))
	var jobApplication model.JobApplications
	err = db.QueryRow("SELECT JobApplicationsID,JobApplicationsName,JobApplicationsPosition,JobApplicationsDescription,JobApplicationsInactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM JobApplications WHERE JobApplicationsID = ?;", id).Scan(&jobApplication.JobApplicationsID, &jobApplication.JobApplicationsName, &jobApplication.JobApplicationsPosition, &jobApplication.JobApplicationsDescription, &jobApplication.JobApplicationsInactive, &jobApplication.CreateBy, &jobApplication.CreateDate, &jobApplication.UpdateBy, &jobApplication.UpdateDate)
	if err != nil {
		return model.JobApplications{}, err, err.Error()
	}
	return jobApplication, nil, "success"
}

func AddJobApplication(c *fiber.Ctx) (model.JobApplications, error, string) {
	jobApplication := new(model.JobApplications)
	if err = c.BodyParser(jobApplication); err != nil {
		return model.JobApplications{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO JobApplications (JobApplicationsName,JobApplicationsPosition,JobApplicationsDescription,JobApplicationsInactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.JobApplications{}, err, err.Error()
	}
	result, err := stmt.Exec(
		jobApplication.JobApplicationsName,
		jobApplication.JobApplicationsPosition,
		jobApplication.JobApplicationsDescription,
		jobApplication.JobApplicationsInactive,
	)
	if err != nil {
		return model.JobApplications{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.JobApplications{}, err, "can't get id"
	}
	var r model.JobApplications
	r.JobApplicationsID = int(lastInsertID)
	r.JobApplicationsName = jobApplication.JobApplicationsName
	r.JobApplicationsPosition = jobApplication.JobApplicationsPosition
	r.JobApplicationsDescription = jobApplication.JobApplicationsDescription
	r.JobApplicationsInactive = jobApplication.JobApplicationsInactive
	return r, nil, "success"
}

func UpdateJobApplication(c *fiber.Ctx) (model.JobApplications, error, string) {
	jobapplicationsid, err := strconv.Atoi(c.Params("jobapplicationsid"))
	content := new(model.JobApplications)
	c.BodyParser(content)

	data, err, _ := GetJobApplication(c)
	if data == (model.JobApplications{}) {
		return model.JobApplications{}, err, err.Error()
	}
	stmt, err := db.Prepare("UPDATE JobApplications SET JobApplicationsName=?, JobApplicationsPosition=?, JobApplicationsDescription=?,JobApplicationsInactive=?, UpdateBy=User(), UpdateDate=NOW() WHERE JobApplicationsID=?")
	if err != nil {
		return model.JobApplications{}, err, err.Error()
	}
	_, err = stmt.Exec(
		content.JobApplicationsName,
		content.JobApplicationsPosition,
		content.JobApplicationsDescription,
		content.JobApplicationsInactive,
		jobapplicationsid,
	)
	if err != nil {
		return model.JobApplications{}, err, err.Error()
	}

	var r model.JobApplications
	r.JobApplicationsID = jobapplicationsid
	r.JobApplicationsName = content.JobApplicationsName
	r.JobApplicationsPosition = content.JobApplicationsPosition
	r.JobApplicationsDescription = content.JobApplicationsDescription
	r.JobApplicationsInactive = content.JobApplicationsInactive
	return r, nil, "success"
}

func DeleteJobApplication(c *fiber.Ctx) (model.JobApplications, error, string) {
	jobapplicationsid, err := strconv.Atoi(c.Params("jobapplicationsid"))
	stmt, err := db.Prepare("DELETE FROM JobApplications WHERE JobApplicationsID=?")
	if err != nil {
		return model.JobApplications{}, err, err.Error()
	}
	_, err = stmt.Exec(
		jobapplicationsid,
	)
	if err != nil {
		return model.JobApplications{}, err, err.Error()
	}
	return model.JobApplications{}, nil, "success"
}

// ## JobApplications

// ## User

func GetUsers() ([]model.User, error, string) {
	var users []model.User
	query := "SELECT UserID,User,Password,Inactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM User;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.UserID, &user.User, &user.Password, &user.Inactive, &user.CreateBy, &user.CreateDate, &user.UpdateBy, &user.UpdateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		users = append(users, user)
	}
	return users, nil, "success"
}

func GetUser(c *fiber.Ctx) (model.User, error, string) {
	id, err := strconv.Atoi(c.Params("userid"))
	var user model.User
	err = db.QueryRow("SELECT UserID,User,Inactive,CreateBy,CreateDate,UpdateBy,UpdateDate FROM User WHERE UserID = ?;", id).Scan(&user.UserID, &user.User, &user.Inactive, &user.CreateBy, &user.CreateDate, &user.UpdateBy, &user.UpdateDate)
	if err != nil {
		return model.User{}, err, err.Error()
	}
	return user, nil, "success"
}

func AddUser(c *fiber.Ctx) (model.User, error, string) {
	user := new(model.User)
	if err = c.BodyParser(user); err != nil {
		return model.User{}, err, err.Error()
	}
	stmt, err := db.Prepare("INSERT INTO User (User,Password,Inactive,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return model.User{}, err, err.Error()
	}
	password, _ := global.HashPassword(user.Password)
	result, err := stmt.Exec(
		user.User,
		password,
		user.Inactive,
	)
	if err != nil {
		return model.User{}, err, err.Error()
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return model.User{}, err, "can't get id"
	}
	var r model.User
	r.UserID = int(lastInsertID)
	r.User = user.User
	r.Inactive = user.Inactive
	return r, nil, "success"
}

func UpdateUser(c *fiber.Ctx) (model.User, error, string) {
	userid, err := strconv.Atoi(c.Params("userid"))
	content := new(model.User)
	c.BodyParser(content)

	data, err, _ := GetUser(c)
	if data == (model.User{}) {
		return model.User{}, err, err.Error()
	}
	password, _ := global.HashPassword(content.Password)
	stmt, err := db.Prepare("UPDATE User SET Password=?, Inactive=? WHERE UserID=?")
	if err != nil {
		return model.User{}, err, err.Error()
	}
	_, err = stmt.Exec(
		password,
		content.Inactive,
		userid,
	)
	if err != nil {
		return model.User{}, err, err.Error()
	}
	var r model.User
	r.UserID = userid
	r.Inactive = content.Inactive
	return r, nil, "success"
}

func DeleteUser(c *fiber.Ctx) (model.User, error, string) {
	userid, err := strconv.Atoi(c.Params("userid"))
	stmt, err := db.Prepare("DELETE FROM User WHERE UserID=?")
	if err != nil {
		return model.User{}, err, err.Error()
	}
	_, err = stmt.Exec(
		userid,
	)
	if err != nil {
		return model.User{}, err, err.Error()
	}
	return model.User{}, nil, "success"
}

// ## User

func AddImage(c *fiber.Ctx) (error, string) {
	img := new(model.ImageJson)
	if err = c.BodyParser(img); err != nil {
		return err, err.Error()
	}
	global.DeleteImage(int64(img.Id), img.ImageType)
	for _, img64 := range img.Image64 {
		err := global.InsertImage(img64, int64(img.Id), img.ImageType)
		if err != nil {
			global.DeleteImage(int64(img.Id), img.ImageType)
			return err, err.Error()
		}
	}
	return nil, "success"
}

func GetImage(c *fiber.Ctx) (model.ImageJson, error, string) {
	id, err := strconv.Atoi(c.Params("id"))
	imagetype := c.Params("imagetype")
	data, err := global.GetImages(int64(id), imagetype)
	if err != nil {
		return model.ImageJson{}, err, err.Error()
	}
	return data, nil, "success"
}

func UpdateConfig(c *fiber.Ctx) (error, string) {
	companyInfo := new(model.CompanyInfo)
	if err = c.BodyParser(companyInfo); err != nil {
		return err, err.Error()
	}
	err, errstring := global.DeleteConfig()
	if err != nil {
		return err, errstring
	}
	global.AddConfig("Advise", companyInfo.Advise)
	global.AddConfig("CompanyAddress", companyInfo.CompanyAddress)
	global.AddConfig("CompanyName", companyInfo.CompanyName)
	global.AddConfig("ContactUs", companyInfo.ContactUs)
	global.AddConfig("AboutUs", companyInfo.AboutUs)
	global.AddConfig("Image641", companyInfo.Image641)
	global.AddConfig("Image642", companyInfo.Image642)
	global.AddConfig("Image643", companyInfo.Image643)
	global.AddConfig("JuristicID", companyInfo.JuristicID)
	global.AddConfig("OurService", companyInfo.OurService)
	global.AddConfig("ServiceLocation", companyInfo.ServiceLocation)
	global.AddConfig("WhyUseUs1", companyInfo.WhyUseUs1)
	global.AddConfig("WhyUseUs2", companyInfo.WhyUseUs2)
	global.AddConfig("WhyUseUs3", companyInfo.WhyUseUs3)
	global.AddConfig("WhyUseUsTitl1", companyInfo.WhyUseUsTitl1)
	global.AddConfig("WhyUseUsTitl2", companyInfo.WhyUseUsTitl2)
	global.AddConfig("WhyUseUsTitl3", companyInfo.WhyUseUsTitl3)
	global.AddConfig("ContactNumber", companyInfo.ContactNumber)
	global.AddConfig("FacebookLink", companyInfo.FacebookLink)
	global.AddConfig("LineLink", companyInfo.LineLink)
	global.AddConfig("GoogleMap", companyInfo.GoogleMap)

	return nil, "success"
}

func GetConfig(c *fiber.Ctx) (model.CompanyInfo, error, string) {
	// configname, err := strconv.Atoi(c.Params("configname"))
	var companyInfo model.CompanyInfo
	rows, err := db.Query("SELECT ConfigName, Value FROM OtherConfigs ;")
	if err != nil {
		return model.CompanyInfo{}, err, err.Error()
	}
	v := reflect.ValueOf(&companyInfo).Elem()
	t := v.Type()
	defer rows.Close()
	for rows.Next() {
		var configName, value string
		if err := rows.Scan(&configName, &value); err != nil {
			return model.CompanyInfo{}, err, err.Error()
		}
		// Find the corresponding field in the struct
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.Name == configName {
				v.FieldByName(field.Name).SetString(value)
				break
			}
		}
	}
	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return model.CompanyInfo{}, err, err.Error()
	}
	return companyInfo, nil, "success"
}

func FrontendGetConfig(c *fiber.Ctx) (model.CompanyInfo, error, string) {
	var companyInfo model.CompanyInfo
	rows, err := db.Query("SELECT ConfigName, Value FROM OtherConfigs;")
	if err != nil {
		return model.CompanyInfo{}, err, err.Error()
	}
	v := reflect.ValueOf(&companyInfo).Elem()
	t := v.Type()
	defer rows.Close()
	for rows.Next() {
		var configName, value string
		if err := rows.Scan(&configName, &value); err != nil {
			return model.CompanyInfo{}, err, err.Error()
		}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.Name == configName {
				v.FieldByName(field.Name).SetString(value)
				break
			}
		}
	}
	if err := rows.Err(); err != nil {
		return model.CompanyInfo{}, err, err.Error()
	}
	return companyInfo, nil, "success"
}

func FrontendGetCarTypes(c *fiber.Ctx) ([]model.FrontendCarType, error, string) {
	var cartypes []model.FrontendCarType
	rows, err := db.Query("SELECT car.CarTypeID, car.CarTypeName, car.CarTypeDesctiption, car.CarTypeInactive, img.ImagePath FROM CarType car LEFT JOIN Images img ON car.CarTypeID = img.CarTypeID WHERE car.CarTypeInactive = 0")
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var cartype model.FrontendCarType
		var imagePath sql.NullString
		err := rows.Scan(&cartype.CarTypeID, &cartype.CarTypeName, &cartype.CarTypeDesctiption, &cartype.CarTypeInactive, &imagePath)
		if err != nil {
			return nil, err, err.Error()
		}
		if imagePath.Valid {
			cartype.Image64 = imagePath.String
		} else {
			cartype.Image64 = ""
		}
		cartypes = append(cartypes, cartype)
	}
	if err := rows.Err(); err != nil {
		return nil, err, err.Error()
	}
	return cartypes, nil, "success"
}

func FrontendGetContents(c *fiber.Ctx) ([]model.FrontendContent, error, string) {
	var contents []model.FrontendContent
	query := "SELECT c.ContentID,c.ContentTitle,c.ContentText , c.Content, i.ImagePath , c.CreateDate FROM Content c LEFT JOIN Images i ON c.ContentID = i.ContentID WHERE c.ContentInactive = 0 ORDER BY c.ContentID DESC;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var content model.FrontendContent
		err = rows.Scan(&content.ContentID, &content.ContentTitle, &content.ContentText, &content.Content, &content.Image64, &content.CreateDate)
		if err != nil {
			return nil, err, err.Error()
		}
		contents = append(contents, content)
	}
	return contents, nil, "success"
}

func FrontendGetContent(c *fiber.Ctx) (model.FrontendContent, error, string) {
	id, err := strconv.Atoi(c.Params("contentid"))
	if err != nil {
		return model.FrontendContent{}, err, err.Error()
	}
	var content model.FrontendContent
	err = db.QueryRow("SELECT c.ContentID,c.ContentTitle,c.ContentText , c.Content, i.ImagePath , c.CreateDate FROM Content c LEFT JOIN Images i ON c.ContentID = i.ContentID WHERE c.ContentID = ?;", id).Scan(&content.ContentID, &content.ContentTitle, &content.ContentText, &content.Content, &content.Image64, &content.CreateDate)
	if err != nil {
		return model.FrontendContent{}, err, err.Error()
	}
	return content, nil, "success"
}

func FrontendGetCarType(c *fiber.Ctx) (model.FrontendCarType, error, string) {
	id, err := strconv.Atoi(c.Params("cartypeid"))
	if err != nil {
		return model.FrontendCarType{}, err, err.Error()
	}
	var cartype model.FrontendCarType
	err = db.QueryRow("SELECT car.CarTypeID,car.CarTypeName,car.CarTypeDesctiption,img.ImagePath FROM CarType car LEFT JOIN Images img ON car.CarTypeID = img.CarTypeID WHERE  car.CarTypeInactive = 0 AND car.CarTypeID = ?;", id).Scan(&cartype.CarTypeID, &cartype.CarTypeName, &cartype.CarTypeDesctiption, &cartype.Image64)
	if err != nil {
		return model.FrontendCarType{}, err, err.Error()
	}
	return cartype, nil, "success"
}

func FrontendGetCars(c *fiber.Ctx) ([]model.FrontendCar, error, string) {
	id, err := strconv.Atoi(c.Params("cartypeid"))
	var cars []model.FrontendCar
	query := "SELECT c.CarID , c.CarName , c.CarDesctiption , c.CarTypeID   FROM Htt.CarType ct LEFT JOIN Car c ON ct.CarTypeID = c.CarTypeID WHERE ct.CarTypeID = ? AND c.CarInactive = 0"
	rows, err := db.Query(query, id)
	if err != nil {
		return []model.FrontendCar{}, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var car model.FrontendCar
		// var imagePath sql.NullString
		err = rows.Scan(&car.CarID, &car.CarName, &car.CarDesctiption, &car.CarTypeID)
		if err != nil {
			return []model.FrontendCar{}, err, err.Error()
		}
		img, err := global.GetImages(int64(car.CarID), constant.CarImage)
		if err != nil {
			return []model.FrontendCar{}, err, err.Error()
		}
		car.Image64 = img.Image64
		cars = append(cars, car)
	}
	return cars, nil, "success"
}

func FrontendGetJobApplications() ([]model.FrontendJobApplications, error, string) {
	var jobApplications []model.FrontendJobApplications
	query := "SELECT JobApplicationsID,JobApplicationsName,JobApplicationsPosition FROM JobApplications WHERE JobApplicationsInactive = 0;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err, err.Error()
	}
	defer rows.Close()
	for rows.Next() {
		var jobApplication model.FrontendJobApplications
		err = rows.Scan(&jobApplication.JobApplicationsID, &jobApplication.JobApplicationsName, &jobApplication.JobApplicationsPosition)
		if err != nil {
			return nil, err, err.Error()
		}
		jobApplications = append(jobApplications, jobApplication)
	}
	return jobApplications, nil, "success"
}

func FrontendGetJobApplication(c *fiber.Ctx) (model.FrontendJobApplications, error, string) {
	id, err := strconv.Atoi(c.Params("jobapplicationsid"))
	if err != nil {
		return model.FrontendJobApplications{}, err, err.Error()
	}
	var job model.FrontendJobApplications
	err = db.QueryRow("SELECT JobApplicationsID,JobApplicationsName,JobApplicationsPosition,JobApplicationsDescription FROM JobApplications WHERE JobApplicationsInactive = 0 AND JobApplicationsID = ?;", id).Scan(&job.JobApplicationsID, &job.JobApplicationsName, &job.JobApplicationsPosition, &job.JobApplicationsDescription)
	if err != nil {
		return model.FrontendJobApplications{}, err, err.Error()
	}
	return job, nil, "success"
}
