package orm

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Column interface {
	Field() string
	SqlType() string
}

type StringColumn struct {
	Name   string
	Length int
}

func (sc StringColumn) Field() string {
	return sc.Name
}

func (sc StringColumn) SqlType() string {
	return fmt.Sprintf("varchar(%d)", sc.Length)
}

type IntegerColumn struct {
	Name   string
	Length int
}

func (ic IntegerColumn) Field() string {
	return ic.Name
}

func (ic IntegerColumn) SqlType() string {
	if ic.Length < 5 {
		return "smallint"
	} else if ic.Length < 9 {
		return "int"
	} else if ic.Length < 20 {
		return "bigint"
	} else {
		log.Fatal("Not supported integer length!")
	}
	return ""
}

var DB *sql.DB

func getDB() (db *sql.DB) {
	if DB == nil {
		db, err := sql.Open("postgres", "user=root host=/var/run/postgresql dbname=dating password=root123")
		if err != nil {
			log.Fatal(err)
		}
		DB = db
		return db
	} else {
		db = DB
	}
	return
}

/*
* For create table
* * CrateTable('students', {"name": string, "no": int})
 */
func CreateTable(tableName string, columns ...Column) bool {
	db := getDB()
	_, err := db.Query(rawCreateTableSql(tableName, columns...))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func DropTable(name string) bool {
	db, err := sql.Open("postgres", "user=root host=/var/run/postgresql dbname=dating password=root123")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Query(fmt.Sprintf("drop table %s", name))
	if err != nil {
		log.Fatal(err)
	}
	return true
}

/*
* "create table students(name varchar(10), no int)"
 */
func rawCreateTableSql(tableName string, columns ...Column) (query string) {
	query = "create table " + tableName + "("
	for i, column := range columns {
		if i != 0 {
			query += ", "
		}
		query += fmt.Sprintf("%s %s", column.Field(), column.SqlType())
	}
	query += ")"
	return
}
