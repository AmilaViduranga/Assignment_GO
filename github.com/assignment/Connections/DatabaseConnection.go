package Connections

import (
	"database/sql"
	"fmt"
	"github.com/assignment/Constants"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CreateConnection() *sql.DB {

	if (Constants.ACTIVE_DATA_BASE == "MYSQL") {

		db, err := sql.Open(Constants.DRIVER_NAME, Constants.USERNAME+":"+Constants.PASSWORD+"@"+Constants.HOST+"/"+Constants.DATABASE)

		if err != nil {
			fmt.Println("Mysql Error")
			panic(err.Error())
		} else {
			fmt.Println("Mysql Connection Successfully")
			return db

		}

	} else {
		log.Print("Active Database Error")

	}
	return nil
}
