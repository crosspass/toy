package model

import (
  "fmt"
  "testing"
  "toy/lib/orm"
)

type Student struct {
  Name orm.Field
  Age orm.Field
}

func (s Stu {

}

func TestCreate() {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if err != nil {
		t.Error("Create table students failed: %s", err)
	}

}
