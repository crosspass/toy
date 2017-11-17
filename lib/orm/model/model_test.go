package model

import (
	"testing"
	"toy/lib/orm"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

func (stu Student) TbName() string {
	return "students"
}

func TestCreate(t *testing.T) {
	err := orm.CreateTable("students", orm.StringColumn{"name", 10}, orm.IntegerColumn{"age", 3})
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	stu := Student{Name: "bob", Age: 17}
	Create(&stu)
	if stu.Id != 1 {
		t.Error("expected: %d, actual: %d", 1, stu.Id)
	}
	orm.DropTable("students")
}

func TestUpdate(t *testing.T) {
	err := orm.CreateTable("students", orm.StringColumn{"name", 10}, orm.IntegerColumn{"age", 3})
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
  stu := Student{Name: "bob", Age: 17}
	Create(&stu)
	stu.Name = "mike"
  ok := Update(stu)
	if !ok {
    t.Error("Update failed!")
	}
	orm.DropTable("students")
}
