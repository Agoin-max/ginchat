package api

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApiController struct{}

type UserBody struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
}

// UserRegister
// @Summary 用户注册
// @Tags 用户模块
// @Param object body api.UserBody false "Json请求体"
// @Param Authorization header string false "Bearer"
// @Success 200 {object} utils.ResponseStruct
// @Router /api/user/register [post]
func (con ApiController) Register(ctx *gin.Context) {
	body := UserBody{}
	if err := ctx.BindJSON(&body); err != nil {
		utils.Fail(ctx, "解析Json失败", nil)
		return
	}
	user := models.UserBasic{}
	user.Name = body.Account
	password := utils.Md5Encode(body.Password)
	user.PassWord = password
	user.Identity = utils.Md5Encode(utils.GetDate())

	// 校验
	if user.Name == "" || body.Password == "" {
		utils.Fail(ctx, "用户/密码不能为空", nil)
		return
	}

	err := utils.DB.Where("name=?", user.Name).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		utils.Fail(ctx, "用户已存在", nil)
	} else {
		utils.DB.Create(&user)
		msg := "注册成功"
		err := utils.Red.Set(ctx, user.Identity, fmt.Sprintf("%d", user.ID), 24*60*time.Minute).Err()
		if err != nil {
			panic(err)
		}
		utils.Success(ctx, msg, user)
	}
}

// UserLogin
// @Summary 用户登录
// @Tags 用户模块
// @Param object body api.UserBody false "Json请求体"
// @Success 200 {object} utils.ResponseStruct
// @Router /api/user/login [post]
func (con ApiController) Login(ctx *gin.Context) {
	body := UserBody{}
	if err := ctx.BindJSON(&body); err != nil {
		utils.Fail(ctx, "解析Json失败", nil)
		return
	}
	user := models.UserBasic{}
	user.Name = body.Account
	password := utils.Md5Encode(body.Password)

	err := utils.DB.Where("name=?", user.Name).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		if password == user.PassWord {
			utils.Success(ctx, "登录成功", user)
		} else {
			utils.Fail(ctx, "登录失败", nil)
		}
	} else {
		utils.Fail(ctx, "登录失败", nil)
	}
}

func (con ApiController) AddFriend(ctx *gin.Context) {
	user := models.UserBasic{}
	user.Name = ctx.PostForm("name")

	err := utils.DB.Where("name=?", user.Name).First(&user).Error
	if err != gorm.ErrRecordNotFound {

	}
}

// SearchFriends
// @Summary 查询联系人
// @Tags 用户模块
// @Param userId query int false "用户ID"
// @Success 200 {object} utils.ResponseStruct
// @Router /api/user/searchFriends [get]
func (con ApiController) SearchFriends(ctx *gin.Context) {
	contacts := []models.Contact{}
	userId := ctx.Query("userId")
	objIds := []uint64{}

	utils.DB.Where("owner_id=? and type=1", userId).Find(&contacts)
	for _, v := range contacts {
		fmt.Println(v)
		objIds = append(objIds, uint64(v.TargetId))
	}

	users := []models.UserBasic{}
	utils.DB.Where("id in ?", objIds).Find(&users)
	// utils.DB.Table("user_basic").Select([]string{"name", "pass_word"}).Where("id in ?", objIds).Scan(&users)
	utils.Success(ctx, "查询联系人成功", users)
}
