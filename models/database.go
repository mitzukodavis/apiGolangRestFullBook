package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/mitzukodavis/apirestgolang/config"
)

var db *sql.DB
var debug bool

func init()  {
	CreateConnection()
	debug = config.Debug()

}

func CreateConnection()  {

	if GetConnection() != nil{
		return
	}
	url := config.UrlDatabase()
	if connention , err := sql.Open("mysql",url); err != nil{
		panic(err)
	}else {
		db = connention
	}
}

func CreateTables()  {
	createTable("users", userSchema)
}

func createTable(tableName, schema string)  {
	if !exitsTable(tableName){
		Exec(schema)
	}else{
		truncateTable(tableName)
	}
}

func truncateTable(tableName string)  {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

func exitsTable(tableName string) bool{
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, _:=Query(sql)
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil && !debug{
		log.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error){
	rows, err :=db.Query(query, args...)
	if err != nil && !debug {
		log.Println(err)
	}
	return  rows, err
}

func InsertData(query string, args ...interface{}) (int64, error) {
	if result, err :=Exec(query, args...); err != nil{
		return int64(0), err
	} else {
		id, err := result.LastInsertId()
		return id, err
	}
 }

func GetConnection()  *sql.DB{
	return db
}

func Ping()  {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func CloseConnection()  {
	db.Close()
}

