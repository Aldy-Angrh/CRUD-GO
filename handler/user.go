package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest-api-starter/model"
	"net/http"
	"strconv"
)

func (h *Handler)CreateUser(c *gin.Context)    {
	var req model.RequestUser
	var response model.Response
	err := c.BindJSON(&req)
	if err != nil{
		response.ResponseCode = "400"
		response.Deskripsi = "Bad Request"
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	if req.Pekerjaan !="PNS" && req.Pekerjaan !="Wirausaha" && req.Pekerjaan !="Wiraswasta"{
		response.ResponseCode = "400"
		response.Deskripsi = "Baduest"
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	if req.Pendidikan != model.SD && req.Pendidikan != model.SMP && req.Pendidikan != model.SMA &&
		req.Pendidikan != model.Sarjana && req.Pendidikan != model.Magister && req.Pendidikan != model.Doktor{
		response.ResponseCode = "400"
		response.Deskripsi = "Gagal"
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	err = h.Controller.CreateUser(req.UserName, req.BirthDate, req.NoKtp, req.Pekerjaan, req.Pendidikan)
	fmt.Println()
	if err != nil {
		response.ResponseCode = "400"
		response.Deskripsi = err.Error()
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	response.ResponseCode = "200"
	response.Deskripsi = "Success Create User"
	response.Data= err
	c.JSON(http.StatusOK, response)
	return
}
func (h *Handler)GetUserById(c *gin.Context)  {
	var response model.Response
	userID:= c.Query("user_id")
	userIdIN,err := strconv.Atoi(userID)
	if err != nil {
		response.ResponseCode = "400"
		response.Deskripsi = err.Error()
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	err,res := h.Controller.GetUserById(userIdIN)
	if err != nil {
		response.ResponseCode = "400"
		response.Deskripsi = err.Error()
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	response.ResponseCode = "200"
	response.Deskripsi = "Get Data Success"
	response.Data = res
	c.JSON(http.StatusOK, response)
	return
}
func (h *Handler)GetAllUsers(c *gin.Context)  {
	var response model.Response
	r,res := h.Controller.GetAllUser()
	if r != nil {
		response.ResponseCode = "500"
		response.Deskripsi = r.Error()
		c.JSON(http.StatusInternalServerError, response)
		c.Abort()
		return
	}
	response.ResponseCode = "200"
	response.Deskripsi = "Get Data Success"
	response.Data = res
	c.JSON(http.StatusOK, response)
	return
}
func (h *Handler)UpdateUser(c *gin.Context)  {
	var response model.Response
	var req model.RequestUser
	err := c.BindJSON(&req)
	if err != nil{
		response.ResponseCode = "400"
		response.Deskripsi = "Bad Request"
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	userID:= c.Query("user_id")
	userIdIN,err := strconv.Atoi(userID)
	if err != nil {
		response.ResponseCode = "400"
		response.Deskripsi = err.Error()
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	err = h.Controller.UpdateUser(userIdIN, req.UserName, req.BirthDate, req.Pekerjaan, req.NoKtp, req.Pendidikan)
	if err != nil {
		response.ResponseCode = "400"
		response.Deskripsi = err.Error()
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}
	response.ResponseCode = "200"
	response.Deskripsi = " Update Success"
	c.JSON(http.StatusOK, response)
	return
}