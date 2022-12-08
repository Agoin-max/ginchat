package admin

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(ctx *gin.Context) {
	fmt.Println("xxx")
	sTime := utils.UnixToTime(1629788418)
	fmt.Println(sTime)
	con.success(ctx)
}

func (con UserController) List(ctx *gin.Context) {
	ctx.String(200, "用户列表")
}

func (con UserController) Article(ctx *gin.Context) {
	ctx.String(200, "新闻列表")
}

func (con UserController) UserList(ctx *gin.Context) {
	// 查询数据库
	userList := []models.UserBasic{}
	utils.DB.Find(&userList)
	fmt.Println(userList)

	userL := []models.UserBasic{}
	utils.DB.Where("name=?", "申专").First(&userL)
	fmt.Println(userL)

	ctx.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

func (con UserController) Add(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.Query("name")
	password := utils.Md5Encode(ctx.Query("password"))
	user.PassWord = password
	utils.DB.Create(&user)
	msg := "新增用户成功"
	utils.Success(ctx, msg, user)
}

func (con UserController) Edit(ctx *gin.Context) {
	// 方式一
	// user := models.UserBasic{
	// 	PassWord: "123",
	// }
	// utils.DB.Find(&user)
	// user.Email = "goin@qq.com"
	// utils.DB.Save(&user)

	// 方式二
	// ins := models.UserBasic{}
	// utils.DB.Model(&ins).Where("id=?", 4).Update("name", "Sophie")

	// 方式三
	user := models.UserBasic{}
	id := ctx.PostForm("id")
	utils.DB.Where("id=?", id).Find(&user)
	user.Name = "Sophie"
	utils.DB.Save(&user)

	ctx.String(200, "修改用户")
}

func (con UserController) Delete(ctx *gin.Context) {
	user := models.UserBasic{}
	utils.DB.Where("name=?", "").Delete(&user)

	ctx.String(200, "删除用户")
}
