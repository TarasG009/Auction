package DataBase

//import (
//	"github.com/go-sql-driver/mysql"
//	"database/sql"
//	"fmt"
//	"log"
//	"os"
//)
//
//func Connect() {
//
//	var db *sql.DB
//	var err error
//
//	cfg := mysql.
//		User:   os.Getenv("user"),
//		Passwd: os.Getenv("password"),
//		Net:    "tcp",
//		Addr:   "127.0.0.1:3306",
//		DBName: "goods",
//	}
//
//	db, err = sql.Open("mysql", cfg.FormatDSN())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	pingErr := db.Ping()
//	if pingErr != nil {
//		log.Fatal(pingErr)
//	}
//	fmt.Println("Connected!")
//}
