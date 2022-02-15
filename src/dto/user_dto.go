package dto

/// UserDto model info
// @Description User account information
// @Description with user id and username
type UserDto struct {
	Id           string `json:"id"`
	UserName     string `json:"userName"`
	UserWrorkId  string `json:"userWrorkId"`
	UserPassword string `json:"userPassword"`
	Login        bool   `json:"login"`
	LoginKey     string `json:"loginKey"`
}
