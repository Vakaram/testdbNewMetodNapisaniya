package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"testdbNewMetods/store"
)

// 1
type APIServer struct {
	config *Config        //
	logger *logrus.Logger //добавляем логгер в apiserver
	router *mux.Router    // Все входящие запросы будут отправлены на маршрутизатор.
	store  *store.Store   // добавили указатель на stor
}

// 2 функция возвращает указатель на api сервер а в нутри инициализирует его чтобы вернуть
// для старта сервера нужно передать конфиг чтобы ему было куда подключаться и тд
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),    // logrus.New() - это встроенно в логрус а не нами написано
		router: mux.NewRouter(), // возвращает новый экземпляр маршрутизатора
	}
}

//3 функция старт нудна запускать http сервер и подключаться к базе данных
//могут быть ошибки потом обработать их

func (s *APIServer) Start() error {
	if err := s.configurateLogger(); err != nil { // ниже мы создали логгер с указанием уровня логирования
		return err
	}
	//s.configRouter() //Добавляем вызов конфигурации маршрутов

	s.logger.Info("Запустили сервер Vakaram ")
	s.configRouter() // это мы пережали конфиграцию которая будет отслеживать тот набор команд которые в ней есть еще можно добавлять и расширять её

	//передадим в appserver наше хранилище чтобы у него всегда был доступ
	if err := s.configStore(); err != nil {
		return err
	}

	return http.ListenAndServe(s.config.Server.Address, s.router)

	//return http.ListenAndServe(":8080", s.router)
}

// функция для определния поведения логера *конфиг логер типо куда что сохраняем уровень логирование и  тд
func (s *APIServer) configurateLogger() error {
	leverl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	//взяли обратились к сткрутре у которой есть функция в которой левл записан как
	//debug и потом уже достали это значение и передали его s.logger.SetLevel(leverl)
	s.logger.SetLevel(leverl)
	return nil
}

// запросы будут обрабатываться тут
func (s *APIServer) configRouter() {
	s.router.HandleFunc("/hello", s.HandleHello) // сюда придем посмотрим а уж потом вызовим функцию которая ниже =)
	s.router.HandleFunc("/testapi", s.TestApi)
	s.router.HandleFunc("/createemployee", s.CreateEmployee)
}

func (s *APIServer) configStore() error {
	st := store.New(s.config.Store) // сюда должны передать строку с подключением
	//пробуем открыть хранилище
	if err := st.Open(); err != nil {
		return err
	}

	//если ошибки не было значит эту переменную записываем в в АПИсервер
	s.store = st
	return nil
}

func (s *APIServer) HandleHello(w http.ResponseWriter, r *http.Request) {
	//service := epmloyeeService.NewMyService() // определили сервис
	response := s.config.Servis.SayHello() // определили нужный метод
	fmt.Fprintf(w, response)               // передали нужный метод
}

func (s *APIServer) TestApi(w http.ResponseWriter, r *http.Request) {
	response := s.config.Servis.TestApi() // определили нужный метод // не совсем понятно что для чего
	fmt.Fprintf(w, response)              // передали нужный метод
}

func (s *APIServer) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	response := s.config.Servis.CreateEmployee()
	fmt.Fprintf(w, response) // передали нужный метод

}

//func (s *APIServer) HandleHello() http.HandlerFunc { // важно без s *APIServer не работает
//	// если возвращать http.HandlerFunc  то код который будет тут внутри он будет супер локальный и не будет захламления
//	return func(w http.ResponseWriter, r *http.Request) {
//		io.WriteString(w, "Helo!!!!!")
//		//if _, err := io.WriteString(os.Stdout, "Hello World!!!!"); err != nil {
//		//	log.Fatal(err)
//		//} // осталось вызвать функцию в старте кода
//	}
//}
