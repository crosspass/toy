package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	// "flag"
)

/*
* create table students
* Example:
*     create_table_students
* output file: ./create_table_students
 */
func main() {
	if len(os.Args) == 1 {
		usage()
	}

	sqls := strings.Split(os.Args[1], "_")
	if len(sqls) < 3 {
		usage()
	}

	var colstr string
	for _, arg := range os.Args[2:] {
		column := strings.Split(arg, ":")
		switch len(column) {
		case 1:
			colstr += fmt.Sprintf("orm.StringColumn{ \"%s\", 20 }", column[0])
		case 2:
			switch column[1] {
			case "string":
				colstr += fmt.Sprintf("orm.StringColumn{ \"%s\", 20 }", column[0])
			case "integer":
				colstr += fmt.Sprintf("orm.IntegerColumn{ \"%s\", 9 }", column[0])
			}
		case 3:
			switch column[1] {
			case "string":
				colstr += fmt.Sprintf("orm.StringColumn{ \"%s\", %s }", column[0], column[2])
			case "integer":
				colstr += fmt.Sprintf("orm.IntegerColumn{ \"%s\", %s }", column[0], column[2])
			}
		default:
			usage()
		}
	}

	var sql string
	switch sqls[0] {
	case "create":
		sql = fmt.Sprintf("orm.CreateTable(\"%s\", %s)", sqls[2], colstr)
	case "add":
		sql = fmt.Sprintf("orm.CreateTable(\"%s\", %s)", sqls[2], colstr)
	case "change":
		sql = fmt.Sprintf("orm.CreateTable(\":%s\", %s)", sqls[2], colstr)
	default:
		usage()
	}

	content := `
package main

import (
  "toy/lib/orm"
)

func main() {
  %s
}
`

	name := fmt.Sprintf("%s_%d.go", os.Args[1], time.Now().Unix())
	file, err := os.Create(name)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	file.WriteString(fmt.Sprintf(content, sql))
	// os.Args[1]
	// file.write(fmt.Sprintf(content, tb))
}

func usage() {
	str := `
  Useage:
    Examle: migrate create_table_sudents name:string age:integer
  `
	fmt.Println(str)
	os.Exit(1)
}
