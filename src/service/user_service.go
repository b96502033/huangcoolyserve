package service

import (
	"fmt"
	"huangcoolyserver/src/dao"
	"huangcoolyserver/src/dto"
	"huangcoolyserver/src/paramDto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
}

func (us *UserService) GetUserById(userId string) *dto.UserDto {

	comUserDao := dao.NewComUserDao()
	objID, _ := primitive.ObjectIDFromHex(userId)
	comUser, _ := comUserDao.GetComUserByUserId(objID)

	userDto := &dto.UserDto{
		Id:           comUser.Id.Hex(),
		UserName:     comUser.UserName,
		UserWrorkId:  comUser.UserWrorkId,
		UserPassword: comUser.UserPassword,
		Login:        comUser.Login,
		LoginKey:     comUser.LoginKey,
	}

	return userDto
}

func (us *UserService) Login(paramUserLoginDto *paramDto.ParamUserLoginDto) *dto.UserLoginDto {

	usrLoginDto := new(dto.UserLoginDto)
	comUserDao := dao.NewComUserDao()
	checkComUser, _ := comUserDao.GetComUserByUserWorkId(paramUserLoginDto.UserWorkId)
	if checkComUser == nil {
		usrLoginDto.ResponseStatusId = -2
		return usrLoginDto
	}
	if checkComUser.UserPassword != paramUserLoginDto.UserPassword {
		usrLoginDto.ResponseStatusId = -3
		return usrLoginDto
	}
	err := comUserDao.UpdateComUserLogin(checkComUser.Id, true)

	if err != nil {
		usrLoginDto.ResponseStatusId = -3
		return usrLoginDto
	}
	comUser, _ := comUserDao.GetComUserByUserWorkId(paramUserLoginDto.UserWorkId)

	usrLoginDto.ResponseStatusId = 0
	usrLoginDto.Login = comUser.Login
	usrLoginDto.UserId = comUser.Id.Hex()
	usrLoginDto.UserName = comUser.UserName
	usrLoginDto.LoginKey = comUser.LoginKey
	return usrLoginDto
}

func (us *UserService) CheckUserLoginStatus(paramUserIdDto *paramDto.ParamUserIdDto) *dto.UserLoginDto {

	fmt.Println(paramUserIdDto.UserId)
	usrLoginDto := new(dto.UserLoginDto)
	comUserDao := dao.NewComUserDao()
	objID, _ := primitive.ObjectIDFromHex(paramUserIdDto.UserId)
	comUser, _ := comUserDao.GetComUserByUserId(objID)

	usrLoginDto.ResponseStatusId = 0
	usrLoginDto.Login = comUser.Login
	usrLoginDto.UserId = comUser.Id.Hex()
	usrLoginDto.UserName = comUser.UserName
	usrLoginDto.LoginKey = comUser.LoginKey

	return usrLoginDto
}

func (us *UserService) GetRoleAccessSystem(userId string) *dto.UserLoginDto {

	usrLoginDto := new(dto.UserLoginDto)
	comUserDao := dao.NewComUserDao()
	objID, _ := primitive.ObjectIDFromHex(userId)
	comUser, _ := comUserDao.GetComUserByUserId(objID)

	usrLoginDto.ResponseStatusId = 0
	usrLoginDto.UserId = comUser.Id.Hex()
	usrLoginDto.UserName = comUser.UserName

	navbarList := new([]dto.NavbarDto)
	comUserNavbarDao := dao.NewComUserNavbarDao()
	comUserNavbarSlice, _ := comUserNavbarDao.GetComUserNavbarById(comUser.Id)
	sysNavbarDao := dao.NewSysNavbarDao()
	for _, comUserNavbar := range comUserNavbarSlice {
		sysNavbar, _ := sysNavbarDao.GetSysNavbarById(comUserNavbar.SysNavbarId)
		navbarDto := dto.NavbarDto{
			NavbarId:   sysNavbar.Id.Hex(),
			NavbarName: sysNavbar.NavbarName,
		}
		*navbarList = append(*navbarList, navbarDto)
	}
	usrLoginDto.NavbarDtoList = *navbarList

	return usrLoginDto
}
