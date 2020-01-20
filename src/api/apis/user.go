package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/williamwq/go/src/api/commons"
	"github.com/williamwq/go/src/api/database"
	model "github.com/williamwq/go/src/api/models"
	"net/http"
	"strconv"
)

//列表数据
func Users(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = c.Request.FormValue("password")
	if user.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名不能为空",
		})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "密码不能为空",
		})
		return
	}
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

//添加数据
func Store(c *gin.Context) {
	var user model.User
	user.Username = c.Request.FormValue("username")
	user.Password = commons.Md5v(c.Request.FormValue("password"))
	if user.Username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名不能为空",
		})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "密码不能为空",
		})
		return
	}
	id, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    id,
	})
}

//修改数据
func Update(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	user.Password = c.Request.FormValue("password")
	result, err := user.Update(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "修改成功",
	})
}

//删除数据
func Destroy(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := user.Destroy(id)
	if err != nil || result.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
	})
}

func InsertRedis(c *gin.Context) {
	var user model.User
	id, err := strconv.ParseInt(c.Request.FormValue("id"), 10, 64)
	result, err := user.Find(id)

	err = database.SetJson("redis", result)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
	})

}

func UserRegister(c *gin.Context) {
	var user model.User
	user.Age, _ = strconv.Atoi(c.Request.FormValue("age"))
	username := c.Request.FormValue("username")
	user.Username = username
	result, erro := user.FindByName(username)
	if erro != nil {
		fmt.Println("查询出错")
		return
	}
	if len(result) > 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "已经有人注册此用户名",
			"data": "",
		})
		return
	}
	password := commons.Md5v(c.Request.FormValue("password"))
	user.Password = password
	user.Address = c.Request.FormValue("address")
	user.Token = commons.Md5v(username + password)
	id, erro := user.Insert()
	if erro != nil {
		c.JSON(1, gin.H{
			"code": 1,
			"msg":  "注册失败",
			"data": "",
		})
		return
	}
	c.JSON(0, gin.H{
		"code":  http.StatusOK,
		"id":    id,
		"msg":   "注册成功",
		"token": user.Token,
	})

}

func Login(c *gin.Context) {
	var user model.User
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user.Username = username
	result, erro := user.FindByNames(username)
	if erro != nil {
		fmt.Println("查询出错")
		return
	}
	if result[0].Password != commons.Md5Encrypt(password) {
		c.JSON(0, gin.H{
			"code":  http.StatusOK,
			"msg":   "注册成功",
			"token": result[0],
		})
		errs := database.SetJson(result[0].Username, result)
		if errs != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "添加失败",
			})
			return
		}
	}

}
