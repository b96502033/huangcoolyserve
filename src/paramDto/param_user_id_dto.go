package paramDto

/// Employee model info
// @Description User account information
// @Description with user id and username
type ParamUserIdDto struct {
	UserId       string `json:"userId,omitempty"`
	UserPassword string `json:"userPassword,omitempty"`
}
