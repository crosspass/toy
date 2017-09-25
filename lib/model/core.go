package model

import (
  "toy/lib/orm"
  "reflect"
  "fmt"
)

type Attributes [string]interface{}

type Model interface {
  TbName() string
}

func Create(model Model) {
  var fields []Field
  v := reflect.ValueOf(model)
  for i := 0; i < v.NumField(); i++ {
    if field, ok := v.Field(i).(Field); ok {
      append(fields, field)
    }
  }
  orm.CreateRecord(model.TbName(), fields....)
}
