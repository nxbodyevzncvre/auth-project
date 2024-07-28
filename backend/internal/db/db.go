package db

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nxbodyevzncvre/mypackage/internal/config"
)

var DB *sql.DB
func Init() error{
	conf := config.GetConfig()
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Dbname)

	log.Printf("Connecting to database with DSN: %s", dsn)

	DB, err = sql.Open("postgres", dsn)
	
	if err != nil{
		log.Printf("error opening db %v", err)
	}

	if err = DB.Ping(); err !=nil{
		fmt.Printf("error pingign db %s", err)
		return err
	}
	log.Println("success")
	return nil
}