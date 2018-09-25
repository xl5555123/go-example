package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id int `orm: "pk;auto"`
	Name string
}

var userOrm = orm.NewOrm()

func init()  {
	orm.RegisterModel(new(User))
}

func AddUser(name string) User {
	var user = User{Name:name}
	id, _ := userOrm.Insert(user)
	var result User
	userOrm.QueryTable("user").Filter("Id", id).One(&result)
	return result
}

func GetAllUsers() []User {
	var users []User
	userOrm.QueryTable("user").All(&users)
	return users
}

func GetUser(id int) User  {
	var user User
	userOrm.QueryTable("user").One(&user)
	return user
}

func UpdateUser(user User) User {
	userOrm.Update(user)
	var result User
	userOrm.QueryTable("user").Filter("Id", user.Id).One(&result)
	return result
}

func RemoveUser(user User)  {
	userOrm.Delete(user)
}

func GetDefaultUser() *User {
	var user = User { Id: 1, Name: "example"}
	return &user
}

func GetNUsers(n int) []User {
	var list []User = make([]User, n)
	var i int
	for i = 0; i < n; i++ {
		list[i] = *GetDefaultUser()
	}
	return list
}

