package dao

import (
	"context"
	"huangcoolyserver/src/model"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeDao struct {
	MgoBaseDao
	collection *mongo.Collection
	onceC      sync.Once
}

func NewEmployeeDao() *EmployeeDao {
	employeeDao := new(EmployeeDao)
	employeeDao.onceC.Do(func() {
		collection := getCollection("empolyee")
		employeeDao.collection = collection
	})

	return employeeDao
}

func (e *EmployeeDao) InsertmployeeList(employeeList []*model.Employee) {

	var t []interface{}
	for _, employee := range employeeList {
		t = append(t, employee)
	}

	_, err := e.collection.InsertMany(context.TODO(), t)
	if err != nil {
		panic(err)
	}

}

func (e *EmployeeDao) Insertmploye(employee *model.Employee) {
	_, err := e.collection.InsertOne(context.TODO(), employee)
	if err != nil {
		panic(err)
	}

}
