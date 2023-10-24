package main

import (
	"fmt"
	"log"
	"testdbNewMetods/epmloyeeService"
	"testdbNewMetods/internal/app/apiserver"
	"testdbNewMetods/store"
)

func main() {
	//5 создадим наш конфиг - это стркутура которая содержит адреса пароли localhosn 8080 и подобное
	config := apiserver.NewConfig()
	//fmt.Println(config)
	fmt.Printf("Вижу в main %s \n", config)

	service := epmloyeeService.NewMyServiceEmployee()
	config.Servis = service
	instansStorage := store.NewConfig()
	config.Store = instansStorage

	//4 теперь нужно создать сервер
	s := apiserver.New(config)
	fmt.Printf("Вижу в main %s \n", config)
	//db :=store.New(store.NewConfig())
	//config.Server = db
	fmt.Printf("После Вижу в main %s \n", config)
	//теперь вызываем метод старт у него
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
