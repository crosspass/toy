package model

import (
	"testing"
	"toy/lib/orm"
)

type Hospital struct {
	Id   int
	Name string
	Age  int
}

func (stu Student) TbName() string {
	return "hospital"
}

