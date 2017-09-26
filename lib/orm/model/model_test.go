package model

import (
	"testing"
	"toy/lib/orm"
)

type Student struct {
	Name orm.Field
	Age  orm.Field
}

func (stu Student) TbName() string {
	return "students"
}

func TestCreate(t *testing.T) {
	err := orm.CreateTable("students", orm.StringColumn{"name", 10}, orm.IntegerColumn{"age", 3})
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	stu := Student{orm.Field{"name", "bob"}, orm.Field{"age", 17}}
	Create(stu)
	orm.DropTable("students")
}

func TestUpdate(t *testing.T) {
	err := orm.CreateTable("students", orm.StringColumn{"name", 10}, orm.IntegerColumn{"age", 3})
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	stu := Student{orm.Field{"name", "bob"}, orm.Field{"age", 17}}
	Create(stu)
	stu.Name = orm.Field{"name", "mike"}
	Update(stu)
	if v, ok := stu.Name.Value.(string); ok {
		if v != "mike" {
			t.Errorf("expected: %s, actual: %v", "mike", stu.Name.Value)
		}
	}
	orm.DropTable("students")
}
