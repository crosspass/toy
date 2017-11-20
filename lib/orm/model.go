package orm

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Attributes map[string]interface{}

type Model interface {
	TbName() string
}

func Find(model Model, whereFields ...Field) {
	row := FindRecord(model.TbName(), whereFields...)
	var fields []interface{}

	valOf := reflect.ValueOf(model).Elem()
	for i := 0; i < valOf.NumField(); i++ {
		switch valOf.Field(i).Kind().String() {
		case "int":
			f := new(int)
			fields = append(fields, f)
		case "rune":
			f := new(rune)
			fields = append(fields, f)
		case "int8":
			f := new(int8)
			fields = append(fields, f)
		case "int16":
			f := new(int16)
			fields = append(fields, f)
		case "int32":
			f := new(int32)
			fields = append(fields, f)
		case "int64":
			f := new(int8)
			fields = append(fields, f)
		case "float32":
			f := new(float32)
			fields = append(fields, f)
		case "float64":
			f := new(float64)
			fields = append(fields, f)
		case "string":
			f := new(string)
			fields = append(fields, f)
		default:
			panic(fmt.Sprintf("Type %s don't supported!", valOf.Field(i).Kind()))
		}
	}
	row.Scan(fields...)
	for i := 0; i < valOf.NumField(); i++ {
		switch kind := valOf.Field(i).Kind().String(); kind {
		case "int":
			l, ok := fields[i].(*int)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetInt(int64(*l))
		case "rune":
			l, ok := fields[i].(*rune)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetInt(int64(*l))
		case "int8":
			l, ok := fields[i].(*int8)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetInt(int64(*l))
		case "int16":
			l, ok := fields[i].(*int16)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetInt(int64(*l))
		case "int32":
			l, ok := fields[i].(*int32)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetInt(int64(*l))
		case "int64":
			l, ok := fields[i].(*int64)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetInt(int64(*l))
		case "float32":
			l, ok := fields[i].(*float32)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetFloat(float64(*l))
		case "float64":
			l, ok := fields[i].(*float64)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetFloat(float64(*l))
		case "string":
			l, ok := fields[i].(*string)
			if !ok {
				panic("assert failed")
			}
			valOf.Field(i).SetString(*l)
		default:
			panic(fmt.Sprintf("Type %s don't supported!", kind))
		}
	}
}

func Create(model Model) error {
	fields := getFiledsWithoutId(model)
	rows, err := CreateRecord((model).TbName(), fields...)
	if err != nil {
		return err
	}
	setFields(model, rows)
	return nil
}

func Update(model Model, fields ...Field) error {
	rows, err := UpdateRecord(model.TbName(), []Field{getId(model)}, fields...)
	if err != nil {
		return err
	}
	return setFields(model, rows)
}

func Save(model Model) error {
	fields := getFiledsWithoutId(model)
	rows, err := UpdateRecord(model.TbName(), []Field{getId(model)}, fields...)
	if err != nil {
		return err
	}
	return setFields(model, rows)
}

func getId(model Model) Field {
	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name == "Id" {
			return Field{"id", v.Field(i).Interface()}
		}
	}
	return Field{}
}

func getFileds(model Model) []Field {
	var fields []Field
	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()
	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, Field{t.Field(i).Name, v.Field(i).Interface()})
	}
	return fields
}

func getFiledsWithoutId(model Model) []Field {
	var fields []Field
	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, Field{t.Field(i).Name, v.Field(i).Interface()})
	}
	return fields
}

// Fetch data from database and set to model
func setFields(model Model, rows *sql.Rows) error {
	var err error
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	valOf := reflect.ValueOf(model).Elem()
	typeOf := reflect.TypeOf(model).Elem()
	var fields []interface{}
	var f_h = map[string]interface{}{}

	for i := 0; i < valOf.NumField(); i++ {
		for j := 0; j < len(columnTypes); j++ {
			if columnType := columnTypes[j]; strings.ToLower(typeOf.Field(i).Name) == columnType.Name() {
				switch valOf.Field(j).Kind().String() {
				case "int":
					f := new(int)
					fields = append(fields, f)
					f_h[typeOf.Field(i).Name] = f
				case "rune":
					f := new(rune)
					fields = append(fields, f)
					f_h[typeOf.Field(i).Name] = f
				case "int8":
					f := new(int8)
					fields = append(fields, f)
				case "int16":
					f := new(int16)
					fields = append(fields, f)
					f_h[typeOf.Field(i).Name] = f
				case "int32":
					f := new(int32)
					fields = append(fields, f)
					f_h[typeOf.Field(i).Name] = f
				case "int64":
					f := new(int8)
					fields = append(fields, f)
					f_h[typeOf.Field(i).Name] = f
				case "float32":
					f := new(float32)
					fields = append(fields, f)
					f_h[typeOf.Field(i).Name] = f
				case "float64":
					f := new(float64)
					fields = append(fields, f)
					f_h[typeOf.Field(i).Name] = f
				case "string":
					s := new(string)
					fields = append(fields, s)
					f_h[typeOf.Field(i).Name] = s
				default:
					return errors.New(fmt.Sprintf("Type %s don't supported!", valOf.Field(j).Kind()))
				}
				break
			}
		}
	}

	if rows.Next() {
		rows.Scan(fields...)
		for m := 0; m < typeOf.NumField(); m++ {
			if val, ok := f_h[typeOf.Field(m).Name]; ok {
				switch kind := valOf.Field(m).Kind().String(); kind {
				case "int":
					l, ok := val.(*int)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "rune":
					l, ok := val.(*rune)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int8":
					l, ok := val.(*int8)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int16":
					l, ok := val.(*int16)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int32":
					l, ok := val.(*int32)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int64":
					l, ok := val.(*int64)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "float32":
					l, ok := val.(*float32)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetFloat(float64(*l))
				case "float64":
					l, ok := val.(*float64)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetFloat(float64(*l))
				case "string":
					l, ok := val.(*string)
					if !ok {
						return errors.New("assert failed")
					}
					valOf.Field(m).SetString(*l)
				default:
					return errors.New(fmt.Sprintf("Type %s don't supported!", kind))
				}
			}
		}
	}
	return err
}