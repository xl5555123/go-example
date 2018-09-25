package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"quickstart/models"
	"strconv"
)


type UserController struct {
	beego.Controller
}

// @router /users [get]
func (c *UserController) GetAll() {
	users := models.GetAllUsers()
	c.Data["json"] = users
	c.ServeJSON()
}

// @router /users/:id [get]
func (c *UserController) Get() {
	userId, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	fmt.Println(err)
	user := models.GetUser(userId)
	c.Data["json"] = user
	c.ServeJSON()
}

// @router /users [post]
func (c *UserController) Post() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	result := models.AddUser(user.Name)

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /users/:id [put]
func (c *UserController) Put() {
	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	user.Id = userId
	result := models.UpdateUser(user)

	c.Data["json"] = result
	c.ServeJSON()
}

// @router /users/:id [delete]
func (c *UserController) Delete() {
	userId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	var user models.User
	user.Id = userId
	models.RemoveUser(user)
	c.Data["json"] = "success"
	c.ServeJSON()
}



