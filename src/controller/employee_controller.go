package controller

import (
	"strconv"

	"huangcoolyserver/src/model"
	"huangcoolyserver/src/paramDto"
	"huangcoolyserver/src/service"

	"github.com/gin-gonic/gin"
)

var employee *model.Employee

type EmployeeController struct {
}

// GetEmployee example
// @Summary 取得Employee人員
// @Tags Employee
// @Accept  json
// @Produce  json
// @Param id query int true "User Id"
// @Router /huangcoolyserve/empolyee/getEmployee [get]
func (c *BaseController) GetEmployee(ctx *gin.Context) {

	ctx.JSON(200, employee)
}

// AddUser example
// @Summary Add a new user to the store
// @Tags Employee
// @Accept  json
// @Produce  json
// @Param message body paramDto.ParamEmployeeDto true "User Data"
// @Router /huangcoolyserver/empolyee/addEmployee [post]
func (c *BaseController) AddEmployee(ctx *gin.Context) {
	var paramEmployeeDto paramDto.ParamEmployeeDto
	err := ctx.Bind(&paramEmployeeDto)
	if err != nil {
		return
	}
	empolyeeService := new(service.EmployeeService)
	empolyeeService.DoAddEmployee(&paramEmployeeDto)

	ctx.JSON(200, paramEmployeeDto)
}

// UpdateUser example
// @Summary Add a new user to the store
// @Tags Employee
// @Accept  json
// @Produce  json
// @Param name query string true "Employee name"
// @Param age query int false "沒填寫預設則為0"
// @Router /huangcoolyserver/empolyee/updateEmployee [put]
func (c *BaseController) UpdateEmployee(ctx *gin.Context) {
	name := ctx.Query("name")
	stringAge := ctx.DefaultQuery("age", "0") // 若沒輸入age參數則預設為"0"
	age, _ := strconv.Atoi(stringAge)
	employee.Name = name
	employee.Age = age
	ctx.JSON(200, "success")
}
