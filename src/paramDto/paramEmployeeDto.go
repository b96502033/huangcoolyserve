package paramDto

/// ParamEmployeeDto info
// @Description User account information
// @Description with user id and username
type ParamEmployeeDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
