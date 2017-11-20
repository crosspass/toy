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
	var expected = "create table students(id SERIAL PRIMARY KEY, name varchar(10))"
	if expected != rawCreateTableSql("students", nameColumn) {
		t.Errorf("expected: %s, actual: %s", expected, rawCreateTableSql("students", nameColumn))
	}
}

func TestRawCreateTableWithManyStringColumn(t *testing.T) {
	nameColumns := []StringColumn{StringColumn{"name", 10}, StringColumn{"nickname", 9}}
	var expected = "create table students(id SERIAL PRIMARY KEY, name varchar(10), nickname varchar(9))"
	if actual := rawCreateTableSql("students", nameColumns[0], nameColumns[1]); expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestRawCreateTableSqlWithSmallIntegerColumn(t *testing.T) {
	nameColumn := IntegerColumn{"age", 3}
	var expected = "create table students(id SERIAL PRIMARY KEY, age smallint)"
	if acutal := rawCreateTableSql("students", nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestRawCreateTableSqlWithIntegerColumn(t *testing.T) {
	nameColumn := IntegerColumn{"count", 7}
	var expected = "create table students(id SERIAL PRIMARY KEY, count int)"
	if acutal := rawCreateTableSql("students", nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestRawCreateTableSqlWithBigIntegerColumn(t *testing.T) {
	nameColumn := IntegerColumn{"count", 11}
	var expected = "create table students(id SERIAL PRIMARY KEY, count bigint)"
	if acutal := rawCreateTableSql("students", nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestRawCreateTableWithIntegerAndString(t *testing.T) {
	countColumn := IntegerColumn{"count", 11}
	nameColumn := StringColumn{"name", 10}

	var expected = "create table students(id SERIAL PRIMARY KEY, count bigint, name varchar(10))"
	if acutal := rawCreateTableSql("students", countColumn, nameColumn); expected != acutal {
		t.Errorf("expected: %s, acutal: %s", expected, acutal)
	}
}

func TestCreateTable(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10})
	if err != nil {
		t.Error("Create table students failed")
	}
	DropTable("students")
}

func TestCreateTableWithMutiColumns(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if err != nil {
		t.Error(err)
	}
	DropTable("students")
}

func TestRawColumnsStr(t *testing.T) {
	expected := "name varchar(10), age smallint"
	actual := rawColumnsStr(StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestAddColumn(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10})
	if err != nil {
		t.Error(err)
	}
	err = AddColumns("students", IntegerColumn{"age", 3})
	if err != nil {
		t.Error(err)
	}
	DropTable("students")
}

func TestAddMutiColumns(t *testing.T) {
	err := CreateTable("students")
	if err != nil {
		t.Error(err)
	}
	err = AddColumns("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if err != nil {
		t.Error(err)
	}
	DropTable("students")
}

func TestCreateRecord(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if err != nil {
		t.Error("Create table students failed: %s", err)
	}
	rows, err := CreateRecord("students", Field{"name", "bob"}, Field{"age", 20})
	if err != nil {
		t.Error(err)
	}
	var (
		id   int
		name string
		age  int
	)
	if rows.Next() {
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			t.Error(err)
		}
		if e_id, e_name, e_age := 1, "bob", 20; e_id != id || e_name != name || e_age != age {
			t.Errorf("expected: %d, %s, %d actual: %d, %s, %d", e_id, e_name, e_age, id, name, age)
		}
	}
	DropTable("students")
}

func TestUpdateRecord(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if err != nil {
		t.Error("Create table students failed: %s", err)
	}
	CreateRecord("students", Field{"name", "bob"}, Field{"age", 20})
	_, err = UpdateRecord("students", []Field{{"name", "bob"}}, Field{"name", "mike"})
	if err != nil {
		t.Error(err)
	}
	DropTable("students")
}

func TestFetchRecord(t *testing.T) {
	err := CreateTable("students", StringColumn{"name", 10}, IntegerColumn{"age", 3})
	if err != nil {
		t.Error("Create table students failed: %s", err)
	}
	CreateRecord("students", Field{"name", "bob"}, Field{"age", 20})
	rows, err := FetchRecords("students", []Field{{"name", "bob"}})
	defer rows.Close()
	if err != nil {
		t.Error(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		if name != "bob" || age != 20 {
			t.Errorf("expected: name=%s, age=%d, acutal: name=%s, age=%d", "bob", 20, name, age)
		}
	}
	DropTable("students")
}
