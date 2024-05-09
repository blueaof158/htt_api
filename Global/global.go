package global

import (
	database "HTTApi/Database"
	model "HTTApi/Model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
