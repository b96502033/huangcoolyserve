package paramDto

// @Description 使用者登入資訊
// @Description with user id and username
type ParamUserLoginDto struct {
	UserWorkId   string `json:"userWorkId,omitempty"`
	UserPassword string `json:"userPassword,omitempty"`
}
