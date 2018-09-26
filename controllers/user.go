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

/**
 *
 * @api {get} /users/:id 根据id请求用户信息
 * @apiName getUser
 * @apiGroup 用户
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {Number} id   Users unique ID.
 * @apiSuccess {String} name   User name.
 *
 * @apiSuccessExample Success-Response:
 *  HTTP/1.1 200 OK
 *  {
 *      "id": 1,
 *      "name": "johny"
 *  }
 *
 * @apiUse UserNotFoundError
 *
 * @apiVersion 0.0.1
 */
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



