package service

import (
	"huangcoolyserver/src/dao"
	"huangcoolyserver/src/model"
	"huangcoolyserver/src/paramDto"
)

type EmployeeService struct {
}

func (em *EmployeeService) DoAddEmployee(paramEmployeeDto *paramDto.ParamEmployeeDto) {
	employee := new(model.Employee)

	employee.Age = paramEmployeeDto.Age
	employee.Id = paramEmployeeDto.Id
	employee.Name = paramEmployeeDto.Name
	employeeDao := dao.NewEmployeeDao()
	employeeDao.Insertmploye(employee)

}
