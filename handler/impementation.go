package handler

import "go-rest-api-starter/controller"

//Handler is a struct
type Handler struct {
	Controller controller.Controller
}

//NewHandler is function to return struct of handler
func NewHandler(controller controller.Controller) *Handler {
	return &Handler{
		Controller: controller,
	}
}
