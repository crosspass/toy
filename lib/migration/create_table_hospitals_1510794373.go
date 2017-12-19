package main

import (
	"github.com/toy/lib/orm"
)

func main() {
	orm.CreateTable("hospitals", orm.StringColumn{"name", 20})
}
