package models

import (
	"net/http"
	"time"
)

type BookModel struct {
	ID              int        `json:"id" gorm:"autoIncrement"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	Author          string     `json:"author"`
	Publisher       string     `json:"publisher"`
	PublicationYear int        `json:"publicationYear" gorm:"column:publicationYear"`
	CreatedAt       *time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

type BookListModel struct {
	Books []BookModel `json:"list"`
}

func (u *BookModel) Bind(r *http.Request) error {
	return nil
}

func (u *BookModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *BookListModel) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (u *BookModel) TableName() string {
	return "books"
}

func NewBookModel() *BookModel {
	s := new(BookModel)

	return s
}
