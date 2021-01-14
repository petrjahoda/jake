package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func homepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Serving homepage")
	http.ServeFile(writer, request, "./html/homepage.html")
}

type ResponseData struct {
	Company   OneCompany
	Companies []OneCompany
	Result    string
}
type OneCompany struct {
	Name    string
	Entity  string
	Country string
	Address string
	City    string
	State   string
}

func searchOne(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Parsing incoming data")
	var data OneCompany
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing data: " + err.Error())
		var responseData ResponseData
		responseData.Result = "problem parsing data:" + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	fmt.Println("Data parsed for company " + data.Name)
	jakeDatabase, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})
	jakeDB, _ := jakeDatabase.DB()
	defer jakeDB.Close()
	if err != nil {
		fmt.Println("Problem opening database: " + err.Error())
		var responseData ResponseData
		responseData.Result = "problem opening database: " + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	var company Company
	jakeDatabase.Where("Name like ?", data.Name).First(&company)
	if len(company.Name) == 0 {
		fmt.Println("No company found")
		var responseData ResponseData
		responseData.Result = "no company found"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	fmt.Println("Company found")
	foundCompany := OneCompany{
		Name:    company.Name,
		Entity:  company.Entity,
		Country: company.Country,
		Address: company.Address,
		City:    company.City,
		State:   company.State,
	}
	var responseData ResponseData
	responseData.Result = "ok"
	responseData.Company = foundCompany
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func saveOne(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Println("Parsing incoming data")
	var data OneCompany
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing data: " + err.Error())
		var responseData ResponseData
		responseData.Result = "problem parsing data:" + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	fmt.Println("Data parsed for company " + data.Name)
	jakeDatabase, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})
	jakeDB, _ := jakeDatabase.DB()
	defer jakeDB.Close()
	if err != nil {
		fmt.Println("Problem opening database: " + err.Error())
		var responseData ResponseData
		responseData.Result = "problem opening database: " + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	newCompany := Company{
		Model:   gorm.Model{},
		Name:    data.Name,
		Entity:  data.Entity,
		Country: data.Country,
		Address: data.Address,
		City:    data.City,
		State:   data.State,
	}
	jakeDatabase.Save(&newCompany)
	var responseData ResponseData
	responseData.Result = "ok"
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func getAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jakeDatabase, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})
	jakeDB, _ := jakeDatabase.DB()
	defer jakeDB.Close()
	if err != nil {
		fmt.Println("Problem opening database: " + err.Error())
		var responseData ResponseData
		responseData.Result = "problem opening database: " + err.Error()
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	var databaseCompanies []Company
	jakeDatabase.Find(&databaseCompanies)
	fmt.Println("Found " + strconv.Itoa(len(databaseCompanies)) + " companies")
	var companies []OneCompany
	for _, company := range databaseCompanies {
		companies = append(companies, OneCompany{
			Name:    company.Name,
			Entity:  company.Entity,
			Country: company.Country,
			Address: company.Address,
			City:    company.City,
			State:   company.State,
		})
	}
	var responseData ResponseData
	responseData.Result = "ok"
	responseData.Companies = companies
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}
