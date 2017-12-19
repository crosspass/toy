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

type Field struct {
	Name  string
	Value interface{}
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
		db, err := sql.Open("postgres", "user=wu host=/private/tmp dbname=dating password=wu123")
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
func CreateTable(tb string, columns ...Column) error {
	db := getDB()
	_, err := db.Query(rawCreateTableSql(tb, columns...))
	return err
}

func DropTable(name string) error {
	db := getDB()
	_, err := db.Query(fmt.Sprintf("drop table %s", name))
	return err
}

func AddColumns(tb string, columns ...Column) error {
	db := getDB()
	var sql string
	for i, column := range columns {
		if i != 0 {
			sql += ","
		}
		sql += fmt.Sprintf("add %s %v", column.Field(), column.SqlType())
	}

	_, err := db.Query(fmt.Sprintf("alter table %s %s", tb, sql))
	if err != nil {
		log.Println(err)
	}
	return err
}

/*
* "create table students(name varchar(10), no int)"
 */
func rawCreateTableSql(tb string, columns ...Column) (query string) {
	if len(columns) == 0 {
		return fmt.Sprintf("create table %s(id SERIAL)", tb)
	}
	columsStr := rawColumnsStr(columns...)
	return fmt.Sprintf("create table %s(id SERIAL PRIMARY KEY, %s)", tb, columsStr)
}

func rawColumnsStr(columns ...Column) string {
	var columsStr string
	for i, column := range columns {
		if i != 0 {
			columsStr += ", "
		}
		columsStr += fmt.Sprintf("%s %s", column.Field(), column.SqlType())
	}
	return columsStr
}

func CreateRecord(tb string, fields ...Field) (*sql.Rows, error) {
	db := getDB()

	var colStr, valStr string
	for i, field := range fields {
		if i != 0 {
			colStr += ","
			valStr += ","
		}
		colStr += field.Name
		valStr += fmt.Sprintf("'%v'", field.Value)
	}
	str := fmt.Sprintf("insert into %s (%s) values (%s) returning *", tb, colStr, valStr)
	return db.Query(str)
}

func UpdateRecord(tb string, whereFields []Field, fields ...Field) (*sql.Rows, error) {
	db := getDB()
	where := parseWhere(whereFields...)
	var colStr string
	for i, field := range fields {
		if i != 0 {
			colStr += ","
		}
		colStr += fmt.Sprintf("%v = '%v'", field.Name, field.Value)
	}
	str := fmt.Sprintf("update %s set %s  %s returning *", tb, colStr, where)
	return db.Query(str)
}

func CountRecord(tb string, fields ...Field) *sql.Row {
	db := getDB()
	var where = parseWhere(fields...)
	str := fmt.Sprintf("select count(*) from %s %s", tb, where)
	return db.QueryRow(str)
}

func FetchRecords(tb string, fields ...Field) (*sql.Rows, error) {
	db := getDB()
	var where = parseWhere(fields...)
	str := fmt.Sprintf("select * from %s %s", tb, where)
	fmt.Println(str)
	return db.Query(str)
}

// Find record by columns as condition
func FindRecord(tb string, fields ...Field) (*sql.Rows, error) {
	db := getDB()
	var where = parseWhere(fields...)
	str := fmt.Sprintf("select * from %s %s limit 1", tb, where)
	return db.Query(str)
}

func parseWhere(fields ...Field) string {
	if len(fields) == 0 {
		return ""
	}

	var where string
	for i, f := range fields {
		if i != 0 {
			where += " and "
		}
		where += fmt.Sprintf("%s = '%v'", f.Name, f.Value)
	}
	return fmt.Sprintf("where(%s)", where)
}
