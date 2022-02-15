package model

/// Employee model info
// @Description User account information
// @Description with user id and username
type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
