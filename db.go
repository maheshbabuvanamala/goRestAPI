package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Getconnection(username string, password string, dbName string) {
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	stmt, err1 := db.Prepare("CREATE TABLE IF NOT EXISTS TEST1 (ID VARCHAR(100),NAME VARCHAR(200))")
	if err1 != nil {
		fmt.Println(err1.Error())
		panic(err1.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}
}
