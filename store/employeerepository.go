package store

import (
	"fmt"
	"testdbNewMetods/internal/app/model"
)

// нужна для запросов к базе данных и все такое
//здесь будут описа11ны запросы которые мы будем дергать уже в api методах

type EmployeRepository struct {
	store *Store
}

func (employe *EmployeRepository) CreateEmployee(emp *model.Employee) (*model.Employee, error) {
	//store.USER()
	//зеачт ошибка не тут ?
	//fmt.Printf("Вижу такое подключение к db :%s", employe.store.db)
	fmt.Println(employe.store.Open())
	if err := employe.store.db.QueryRow(
		"INSERT INTO turnixSchem.employees (login,password) VALUES ($1,$2) RETURNING id ",
		emp.Login, //
		emp.Password,
	).Scan(&emp.ID); err != nil {
		return nil, err
	}

	return nil, nil

}
