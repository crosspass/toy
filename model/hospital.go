package model

type Hospital struct {
	Id   int
	Name string
}

func (stu Student) TbName() string {
	return "hospital"
}
