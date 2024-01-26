package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

func (pc *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["ID"]

}
