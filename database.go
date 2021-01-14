package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Company struct {
	gorm.Model
	Name    string
	Entity  string
	Country string
	Address string
	City    string
	State   string
}

func CheckDatabase() {
	databaseCheck := false
	for !databaseCheck {
		jakeDatabase, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})
		jakeDB, _ := jakeDatabase.DB()
		defer jakeDB.Close()
		if err != nil {
			fmt.Println("Problem opening database, looks like it does not exist")
			database, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
			sqlDB, _ := database.DB()
			if err != nil {
				fmt.Println("Problem opening main mysql database: " + err.Error())
				continue
			}
			fmt.Println("Creating jake database")
			database.Exec("CREATE DATABASE jake;")
			sqlDB.Close()
		}
		fmt.Println("Jake database already exists")
		if !jakeDatabase.Migrator().HasTable(&Company{}) {
			fmt.Println("Creating table Company")
			err := jakeDatabase.Migrator().CreateTable(&Company{})
			if err != nil {
				fmt.Println("Cannot create table: " + err.Error())
				return
			}
			jakeDatabase.Raw("SELECT create_hypertable('benchmark_data', 'created_at');")
		} else {
			fmt.Println("Updating table Company")
			err := jakeDatabase.Migrator().AutoMigrate(&Company{})
			if err != nil {
				fmt.Println("Cannot update table: " + err.Error())
				return
			}
		}
		databaseCheck = true
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Checking database done")
}
