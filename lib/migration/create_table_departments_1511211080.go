package main

import (
	"toy/lib/orm"
)

func main() {
	orm.CreateTable("departments", orm.StringColumn{"name", 20}, orm.IntegerColumn{"department_category_id", 9}, orm.IntegerColumn{"hospital_id", 9})
}
