package dto

/// UserDto model info
// @Description User account information
// @Description with user id and username
type UserLoginDto struct {
	Login            bool        `json:"login,omitempty"`
	LoginKey         string      `json:"loginKey,omitempty"`
	NavbarDtoList    []NavbarDto `json:"navbarDtoList,omitempty"`
	ResponseStatusId int         `json:"responseStatusId"`
	UserId           string      `json:"userId,omitempty"`
	UserName         string      `json:"userName,omitempty"`
	UserWorkId       string      `json:"userWorkId,omitempty"`
}
