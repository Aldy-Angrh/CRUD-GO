package controller

import "github.com/jinzhu/gorm"

//Controller is a struct
type Controller struct {
	DB *gorm.DB
}

//NewController is function
func NewController(db *gorm.DB) *Controller {
	return &Controller{
		DB: db,
	}
}
