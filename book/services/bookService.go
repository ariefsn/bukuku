package services

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ariefsn/book-store/book/helper"
	"github.com/ariefsn/book-store/book/models"
	"github.com/imroc/req"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var baseUrl = "http://" + os.Getenv("URL_AUTH")

// Init service and register new connection
func InitService(sqlDb *sql.DB) (err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	return nil
}

// Find user by id
func GetUserByID(r *http.Request, id int) (map[string]interface{}, int, error) {
	header := req.Header{
		"Accept": "application/json",
		"Claims": r.Header.Get("Claims"),
	}

	req := req.New()

	idS := strconv.Itoa(id)

	res, err := req.Get(baseUrl+"/user/"+idS, header)

	if err != nil || res.Response().StatusCode != 200 {
		return nil, res.Response().StatusCode, errors.New(res.Response().Status)
	}

	newRes := helper.ResponseModel{}

	res.ToJSON(&newRes)

	if !newRes.Success {
		return nil, newRes.HTTPStatusCode, errors.New(newRes.Message)
	}

	user := newRes.Data.(map[string]interface{})

	return user, 200, nil
}

// Find book by id
func GetBookByID(id int) (*models.BookModel, error) {
	user := models.NewBookModel()

	res := db.Table(user.TableName()).Where("id = ?", id).First(&user)

	return user, res.Error
}

// Find book by email
func GetBookByEmail(email string) (*models.BookModel, error) {
	user := models.NewBookModel()

	res := db.Table(user.TableName()).Where("email = ?", email).First(&user)

	return user, res.Error
}

// Find all books
func GetBooks() ([]models.BookModel, error) {
	users := []models.BookModel{}

	res := db.Table(models.NewBookModel().TableName()).Find(&users)

	return users, res.Error
}

// Create new book
func CreateBook(user *models.BookModel) (int64, error) {
	res := db.Table(user.TableName()).Create(&user)

	return res.RowsAffected, res.Error
}

// Update book
func UpdateBook(id int, data *models.BookModel) int64 {
	res := db.Where("id = ?", id).Save(&data)

	return res.RowsAffected
}

// Delete book
func DeleteBook(data *models.BookModel) int64 {
	res := db.Delete(&data)

	return res.RowsAffected
}
