package models

import (
	"net/http"
	"time"
)

type UserModel struct {
	ID        int        `json:"id" gorm:"autoIncrement"`
	FirstName string     `json:"firstName" gorm:"column:firstName"`
	LastName  string     `json:"lastName" gorm:"column:lastName"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Birth     *time.Time `json:"birth"`
	Address   string     `json:"address"`
	IsAdmin   bool       `json:"isAdmin" gorm:"column:isAdmin"`
	CreatedAt *time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

type UserListModel struct {
	Users []UserModel `json:"list"`
}

func (u *UserModel) Bind(r *http.Request) error {
	return nil
}

func (u *UserModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *UserListModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserModel() *UserModel {
	s := new(UserModel)

	s.IsAdmin = false

	return s
}
