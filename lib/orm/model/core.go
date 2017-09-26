package model

import (
	"reflect"
	"toy/lib/orm"
)

type Attributes map[string]interface{}

type Model interface {
	TbName() string
}

func Create(model Model) Model {
	fields := getFields(model)
	orm.CreateRecord(model.TbName(), fields...)
	return
}

func Update(model Model) bool {
	fields := getFields(model)
	orm.CreateRecord(model.TbName(), fields...)
	return true || false
}

func getFields(model Model) []orm.Field {
	var fields []orm.Field
	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		if field, ok := v.Field(i).Interface().(orm.Field); ok {
			fields = append(fields, field)
		}
	}
	return fields
}
