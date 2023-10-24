package epmloyeeService

import (
	"strconv"
	"testdbNewMetods/internal/app/model"
	"testdbNewMetods/store"
)

type MyService struct {
}

// Конструктор сервиса вроде так правильно
func NewMyServiceEmployee() *MyService {
	return &MyService{}
}

// Метод над myService
func (s *MyService) SayHello() string {
	return "!!!!!!Hi"
}

func (s *MyService) TestApi() string {
	return "This test api"
}
func (s *MyService) CreateEmployee() string {
	AdminMock := &model.Employee{Login: "TestNow", Password: "123123123"}
	repository := store.EmployeRepository{}
	AdminMock, _ = repository.CreateEmployee(AdminMock)

	return strconv.Itoa(AdminMock.ID)
}
