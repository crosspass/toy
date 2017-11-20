package orm

import (
	"testing"
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
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	stu := Student{Name: "bob", Age: 17}
	Create(&stu)
	if stu.Id != 1 {
		t.Errorf("expected: %d, actual: %d", 1, stu.Id)
	}
	defer DropTable("students")
}

func TestFindBy(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	defer DropTable("students")
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	stu := Student{Name: "bob", Age: 17}
	Create(&stu)
	var expectedStu = new(Student)
	err = FindBy(expectedStu, Field{"name", "bob"})
	if err != nil || expectedStu.Age != 17 {
		t.Error("Find failed!", err)
	}
}

func TestFindOrCreateBy(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	defer DropTable("students")
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	var expectedStu = new(Student)
	err = FindOrCreateBy(expectedStu, Field{"name", "bob"})
	if err != nil || expectedStu.Name != "bob" {
		t.Errorf("expected: %v, err: %v ", expectedStu, err)
	}

}

func TestCount(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	defer DropTable("students")
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	actual, err := Count(new(Student))
	if err != nil || actual != 0 {
		t.Errorf("expected: %d, actual: %d, err: %s", 0, actual, err)
	}
}

func TestUpdate(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	defer DropTable("students")
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	stu := Student{Name: "bob", Age: 17}
	Create(&stu)
	err = Update(&stu, Field{"name", "mike"})
	if err != nil || stu.Name != "mike" {
		t.Error("Save failed!", err)
	}
}

func TestSave(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	defer DropTable("students")
	if err != nil {
		t.Errorf("Create table students failed: %s", err)
	}
	stu := Student{Name: "bob", Age: 17}
	Create(&stu)
	stu.Name = "mike"
	err = Save(&stu)
	if err != nil || stu.Name != "mike" {
		t.Error("Save failed!", err)
	}
}
