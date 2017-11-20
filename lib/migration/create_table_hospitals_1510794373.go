package main

import (
	"toy/lib/orm"
)

func main() {
	orm.CreateTable("hospitals", orm.StringColumn{"name", 20})
}
