package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/kardianos/service"
	"net/http"
	"os"
)

const connection = "root:password@tcp(jake_database:3306)/mysql?charset=utf8&parseTime=True&loc=Local"
const databaseConnection = "root:password@tcp(jake_database:3306)/jake?charset=utf8&parseTime=True&loc=Local"
const serviceName = "Jake web application"
const serviceDescription = "Jake web application"

type program struct{}

func main() {
	fmt.Println(serviceName + " starting...")
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot start: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start: " + err.Error())
	}
}

func (p *program) Start(service.Service) error {
	fmt.Println(serviceName + " started")
	go p.run()
	return nil
}

func (p *program) Stop(service.Service) error {
	fmt.Println(serviceName + " stopped")
	return nil
}

func (p *program) run() {
	go CheckDatabase()
	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.GET("/", homepage)
	router.GET("/get_all", getAll)
	router.POST("/search_one", searchOne)
	router.POST("/save_one", saveOne)
	err := http.ListenAndServe(":80", router)
	if err != nil {
		fmt.Println("Problem starting service: " + err.Error())
		os.Exit(-1)
	}
	fmt.Println(serviceName + " running")
}
