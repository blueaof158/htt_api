package global

import (
	database "HTTApi/Database"
	model "HTTApi/Model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db, _ = database.ConnectDB()

func InsertImage(imgBase64 string, id int64, imageType string) error {
	query := fmt.Sprintf("INSERT INTO Images (%s, ImagePath, CreateBy, CreateDate, UpdateBy, UpdateDate) VALUES (?,? ,User(), NOW(), User(), NOW())", imageType)
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, imgBase64)
	if err != nil {
		return err
	}
	return nil
}

func DeleteImage(id int64, imageType string) error {
	query := fmt.Sprintf("DELETE FROM Images WHERE %s = ?", imageType)
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func GetImages(id int64, imageType string) (model.ImageJson, error) {
	var imageJson model.ImageJson
	query := fmt.Sprintf("SELECT ImagePath FROM Images WHERE %s = ?", imageType)
	rows, err := db.Query(query, id)
	if err != nil {
		return imageJson, err
	}
	for rows.Next() {
		var imageName string
		err := rows.Scan(&imageName)
		if err != nil {
			return imageJson, err
		}
		imageJson.Image64 = append(imageJson.Image64, imageName)
	}

	imageJson.Id = int(id)
	imageJson.ImageType = imageType

	return imageJson, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddConfig(configname string, value string) string {
	stmt, err := db.Prepare("INSERT INTO OtherConfigs (ConfigName,Value,CreateBy,CreateDate,UpdateBy,UpdateDate) VALUES (?,?,User(), NOW(), User(), NOW())")
	if err != nil {
		return err.Error()
	}
	_, err = stmt.Exec(
		configname,
		value)
	return ""
}

func DeleteConfig() (error, string) {
	stmt, err := db.Prepare("DELETE FROM OtherConfigs;")
	if err != nil {
		return err, err.Error()
	}
	_, err = stmt.Exec()
	return nil, ""

}
