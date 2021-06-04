package services

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/ariefsn/book-store/auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// Initiate service and register connection
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

	// initiate admin user
	admin := models.DefaultAdminUser()

	user, err := GetUserByEmail(admin.Email)

	if err != nil {
		return err
	}

	if user.Email != admin.Email {
		CreateUser(admin)
	}

	return nil
}

// Find user by id
func GetUserByID(id int) (*models.UserModel, error) {
	user := models.NewUserModel()

	res := db.Table(user.TableName()).Where("id = ?", id).First(&user)

	return user, res.Error
}

// Find user by email
func GetUserByEmail(email string) (*models.UserModel, error) {
	user := models.NewUserModel()

	res := db.Table(user.TableName()).Where("email = ?", email).First(&user)

	return user, res.Error
}

// Find all users
func GetUsers() ([]models.UserModel, error) {
	users := []models.UserModel{}

	res := db.Table(models.NewUserModel().TableName()).Find(&users)

	return users, res.Error
}

// Create new user
func CreateUser(user *models.UserModel) (int64, error) {
	res := db.Table(user.TableName()).Create(&user)

	return res.RowsAffected, res.Error
}

// Update user
func UpdateUser(id int, data *models.UserModel) int64 {
	res := db.Where("id = ?", id).Save(&data)

	return res.RowsAffected
}

// Delete user
func DeleteUser(data *models.UserModel) int64 {
	res := db.Delete(&data)

	return res.RowsAffected
}
