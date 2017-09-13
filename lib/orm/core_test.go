package orm

import (
	"testing"
)

/*
* For create table
* * CrateTable('students', {"name": "string", "no": "int"})
 */
func TestRawCreateTableSql(t *testing.T) {
	nameColumn := StringColumn{"name", 10}
	var expected = "create table students(name varchar(10))"
	if expected != rawCreateTableSql("students", nameColumn) {
		t.Errorf("expected: %s, actual: %s", expected, rawCreateTableSql("students", nameColumn))
	}
}

func TestRawCreateTableWithManyStringColumn(t *testing.T) {
	nameColumns := []StringColumn{StringColumn{"name", 10}, StringColumn{"nickname", 9}}
	var expected = "create table students(name varchar(10), nickname varchar(9))"
	if actual := rawCreateTableSql("students", nameColumns[0], nameColumns[1]); expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestRawCreateTableSqlWithSmallIntegerColumn(t *testing.T) {
	nameColumn := IntegerColumn{"age", 3}
	var expected = "create table students(age smallint)"
	if acutal := rawCreateTableSql("students", nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestRawCreateTableSqlWithIntegerColumn(t *testing.T) {
	nameColumn := IntegerColumn{"count", 7}
	var expected = "create table students(count int)"
	if acutal := rawCreateTableSql("students", nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestRawCreateTableSqlWithBigIntegerColumn(t *testing.T) {
	nameColumn := IntegerColumn{"count", 11}
	var expected = "create table students(count bigint)"
	if acutal := rawCreateTableSql("students", nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestRawCreateTableWithIntegerAndString(t *testing.T) {
	countColumn := IntegerColumn{"count", 11}
	nameColumn := StringColumn{"name", 10}

	var expected = "create table students(count bigint, name varchar(10))"
	if acutal := rawCreateTableSql("students", countColumn, nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestCreateTable(t *testing.T) {
	ok := CreateTable("students", StringColumn{"name", 10})
	if !ok {
		t.Error("Create table students failed")
	}
	DropTable("students")
}

func TestCreateTable2(t *testing.T) {
	ok := CreateTable("students", StringColumn{"name", 10})
	if !ok {
		t.Error("Create table students failed")
	}
	DropTable("students")
}
