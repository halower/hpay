package databse

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *sql.DB


func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:Abc12345@tcp(118.24.27.45:3306)/hpay?parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}