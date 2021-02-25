package controller

import (
	"fmt"
	"go-rest-api-starter/model"
)

func  (c *Controller)CreateUser(UserName ,BirthDate, NoKtp, Pekerjaan, Pendidikan string) (r error)  {
	r=c.DB.Table("user").Create(model.RequestUser{UserName: UserName,BirthDate:
		BirthDate,Pekerjaan: Pekerjaan, NoKtp: NoKtp, Pendidikan: Pendidikan}).Error
	if r != nil{
		return
	}
	return
}
func (c *Controller)GetUserById(UserId int)(r error, user model.User)  {
	r=c.DB.Table("user").Where("id=?",UserId).Find(&user).Error
	if r != nil {
		return
	}
	return
}

func (c *Controller) GetAllUser() ( error,  []model.User){
	var userList  []model.User
	rows,r := c.DB.Table("user").Select("id,userName,birthDate,Pekerjaan,noKtp,Pendidikan").Rows()
	fmt.Println(r)
	if r !=nil{
		return r,[]model.User{}
	}
	defer rows.Close()
	for rows.Next(){
		var user model.User
		r := rows.Scan(&user.UserID,&user.UserName,&user.BirthDate,&user.Pekerjaan,&user.NoKtp,&user.Pendidikan)
	if r != nil {
		return  r,[]model.User{}
	}
	userList = append(userList,user)
	}

	return nil, userList
}
func (c *Controller)UpdateUser(userId int, UserName ,BirthDate, NoKtp, Pekerjaan, Pendidikan string)(r error)  {
	r=c.DB.Table("user").Where("id=?",userId).Updates(map[string]interface{}{"userName": UserName,"birthDate":
	BirthDate,"Pekerjaan": Pekerjaan, "noKtp": NoKtp, "Pendidikan": Pendidikan}).Error
	if r != nil{
		return
	}
	return
}
