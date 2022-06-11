package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// usersLoginInfo 使用map数据结构去存储用户信息，在demo中键值对的key是用户名+密码

// user data will be cleared every time the server starts
// 每次启动服务，用户数据将被清空

// test data: username=zhanglei, password=douyin
// 预先设定的测试数据，
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user User

	//判断数据库中是否出现这个用户
	if rowsAffected := DB.Where("name = ?", username).First(&user).RowsAffected; rowsAffected != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := User{
			Name:     username,
			Password: password,
		}
		DB.Create(&newUser)
		//Create的作用相当于insert
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   newUser.Id,
			Token:    username,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	var user User
	if rowsAffected := DB.Where("name = ?", username).First(&user).RowsAffected; rowsAffected != 0 {
		if user.Password == password {
			//测试是否进来
			fmt.Println("11111111111111111111")
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    user.Name,
			})
		} else { //密码错误
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Incorrect password"},
			})
		}

	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//token里面装的是username,判断username存不不存在
	var user User
	if rowsAffected := DB.Where("name = ?", token).First(&user).RowsAffected; rowsAffected != 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist2"},
		})
	}
}
