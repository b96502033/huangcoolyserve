package controller

import (
	"huangcoolyserver/src/paramDto"
	"huangcoolyserver/src/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// @Summary 取得User
// @Tags User
// @Accept  json
// @Produce  json
// @Param userId query string true "User Id"
// @Router /huangcoolyserver/user/getUser [get]
func (c *BaseController) GetUser(ctx *gin.Context) {
	userId := ctx.Query("userId")

	userService := new(service.UserService)
	userDto := userService.GetUserById(userId)

	ctx.JSON(200, userDto)
}

// @Summary 使用者進行登入
// @Tags User
// @Accept  json
// @Produce  json
// @Param message body paramDto.ParamUserLoginDto true "使用者資訊"
// @Router /huangcoolyserver/user/login [post]
func (c *BaseController) Login(ctx *gin.Context) {
	var paramUserLoginDto paramDto.ParamUserLoginDto
	err := ctx.Bind(&paramUserLoginDto)
	if err != nil {
		return
	}
	userService := new(service.UserService)
	usrLoginDto := userService.Login(&paramUserLoginDto)

	ctx.JSON(200, usrLoginDto)
}

// @Summary 驗證使用者的登陸狀態
// @Tags User
// @Accept  json
// @Produce  json
// @Param message body paramDto.ParamUserIdDto true "使用者資訊"
// @Router /huangcoolyserver/user/checkUserLoginStatus [post]
func (c *BaseController) CheckUserLoginStatus(ctx *gin.Context) {
	var paramUserIdDto paramDto.ParamUserIdDto
	err := ctx.Bind(&paramUserIdDto)
	if err != nil {
		return
	}
	userService := new(service.UserService)
	usrLoginDto := userService.CheckUserLoginStatus(&paramUserIdDto)

	ctx.JSON(200, usrLoginDto)
}

// @Summary 取得使用者可以使用的Navbar
// @Tags User
// @Accept  json
// @Produce  json
// @Param userId path string true "User Id"
// @Router /huangcoolyserver/user/getRoleAccessSystem/{userId} [get]
func (c *BaseController) GetRoleAccessSystem(ctx *gin.Context) {
	userId := ctx.Param("userId")
	userService := new(service.UserService)
	usrLoginDto := userService.GetRoleAccessSystem(userId)

	ctx.JSON(200, usrLoginDto)
}
