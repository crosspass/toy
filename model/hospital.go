package model

type Hospital struct {
	Id   int
	Name string
}

func (hospital Hospital) TbName() string {
	return "hospitals"
}
