package db

import (
	"fmt"

	"github.com/nxbodyevzncvre/mypackage/internal/config"
)


func Init() {
	conf := config.GetConfig()

	fmt.Println(conf.DB.Host)
	fmt.Println(conf.DB.Port)
	fmt.Println(conf.DB.User)
	fmt.Println(conf.DB.Password)
	fmt.Println(conf.DB.Dbname)

}