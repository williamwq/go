
package models

import (
	orm "api/database"
)

type User struct {
	ID       int64  `json:"id"`       // 列名为 `id`
	Username string `json:"username"` // 列名为 `username`
	Password string `json:"password"` // 列名为 `password`
	Address string   `json:"address"` // 地址
	Token  string  `json:"token"`
	Age   int        `json:"age"`   //年龄
}

var Users []User

//添加
func (user User) Insert() (id int64, err error) {

	//添加数据
	result := orm.Eloquent.Create(&user)
	id =user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

//列表
func (user *User) Find(id int64) (users []User, err error) {
	if err = orm.Eloquent.Where("id = ?",id).First(&users).Error; err != nil {
		return
	}
	//fmt.Println(users)
	return
}
//通过 用户名查找用户
func (user *User)FindByName(username string) (result []User,err error) {
	if err = orm.Eloquent.Where("username =?",username).First(&result).Error; err !=nil{
		return
	}
	return
}
//查找是否存在
func (user *User)FindByNameExist(username string) (bools bool,err error) {
	if err = orm.Eloquent.Where("username = ?",username).First(&bools).Error; err !=nil{
		panic("连接失败")
	}
	return
}
//修改
func (user *User) Update(id int64) (updateUser User, err error) {

	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *User) Destroy(id int64) (Result User, err error) {

	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}