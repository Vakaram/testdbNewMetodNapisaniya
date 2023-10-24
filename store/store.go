package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Store struct {
	config             *Config
	db                 *sql.DB
	employeeRepository *EmployeRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}

}

// при инициализации при попытке подключения к бд будет использоваться
func (s *Store) Open() error {
	fmt.Println(s.config.DatabaseURL + "sssssssss")
	db, err := sql.Open("postgres", s.config.DatabaseURL) // 	да правильная строка
	// db, err := sql.Open("postgres", NewConfig())
	//db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=admin dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	//fmt.Println()
	logrus.Info("Пинг прошел все хорошо")
	//fmt.Println(s.config.DatabaseURL)
	// попробуем тут вставит тестовые данные админа в бд
	//это работает значит к бд можно обращаться
	//if err := db.QueryRow(
	//	"INSERT INTO turnixSchem.employees (login,password) VALUES ($1,$2)",
	//	"111", //
	//	"222",
	//); err != nil {
	//	return err.Err()
	//}

	//сам додумался начинаю въезжать
	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
	// пока пусто
}

// чтобы нельзя было заюзать наш репозиторий в обход хранилища ??? не особо понял сам
// будем вызывать вот так store.Employee().Create()
func (s *Store) Employee() *EmployeRepository {
	if s.employeeRepository != nil {
		return s.employeeRepository
	}
	//если его нет инициализируем его
	s.employeeRepository = &EmployeRepository{
		store: s,
	}
	return s.employeeRepository

}
