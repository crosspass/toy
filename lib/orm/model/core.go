package model

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"toy/lib/orm"
)

type Attributes map[string]interface{}

type Model interface {
	TbName() string
}

func Find(model *Model, fields ...Field) error {
  orm.FindRecord()
  if err == nil {
    return model
  }
}

func Create(model Model) (Model, error) {
	fields := getFields(model)
	rows, err := orm.CreateRecord(model.TbName(), fields...)
	if err != nil {
		return nil, err
	}
	setFields(model, rows)
	return model, err
}

func Update(model Model) bool {
	fields := getFields(model)
  err := orm.UpdateRecord(model.TbName(), getId(model),  fields...)
  return err == nil
}

func getId(model Model) orm.Field {
	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name == "Id" {
			return orm.Field{"id", v.Field(i).Interface()}
		}
	}
  return orm.Field{}
}


func getFields(model Model) []orm.Field {
	var fields []orm.Field
	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Name == "Id" {
			continue
		}
		fields = append(fields, orm.Field{t.Field(i).Name, v.Field(i).Interface()})
	}
	return fields
}

// Fetch data from database and set to model
func setFields(model Model, rows *sql.Rows) {
	var err error
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
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
					panic(fmt.Sprintf("Type %s don't supported!", valOf.Field(j).Kind()))
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
						panic("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "rune":
					l, ok := val.(*rune)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int8":
					l, ok := val.(*int8)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int16":
					l, ok := val.(*int16)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int32":
					l, ok := val.(*int32)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "int64":
					l, ok := val.(*int64)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetInt(int64(*l))
				case "float32":
					l, ok := val.(*float32)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetFloat(float64(*l))
				case "float64":
					l, ok := val.(*float64)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetFloat(float64(*l))
				case "string":
					l, ok := val.(*string)
					if !ok {
						panic("assert failed")
					}
					valOf.Field(m).SetString(*l)
				default:
					panic(fmt.Sprintf("Type %s don't supported!", kind))
				}
			}
		}
	}
}
