package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var _HOST = os.Getenv("OTRS_DB_HOST")
var _DB = os.Getenv("OTRS_DB_NAME")
var _USER = os.Getenv("OTRS_DB_USER")
var _PASS = os.Getenv("OTRS_DB_PASS")

var DB *sqlx.DB

func init() {
	if _HOST == "" {
		_HOST = "192.168.88.254:9001"
	}
	if _DB == "" {
		_DB = "otrs"
	}
	if _USER == "" {
		_USER = "otrs"
	}
	if _PASS == "" {
		_PASS = "suP3rp244w0rds910tr4"
	}
	urlLog := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true",
		_USER, "*****", _HOST, _DB)
	url := fmt.Sprintf("%s:%s@(%s)/%s?parseTime=true",
		_USER, _PASS, _HOST, _DB)
	log.Println("try DB", urlLog)
	var err error
	DB, err = sqlx.Connect("mysql", url)
	if err != nil {
		panic(err)
	}
	// force a connection and test that it worked
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}

// Close Close DB
func Close() {
	DB.Close()
}
