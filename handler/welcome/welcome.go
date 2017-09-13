package welcome

import (
	"encoding/json"
	"net/http"
)

// {
//  name: '珠海市人民医院',
//  departments: [
//   id: 1,
//   category: '儿科',
//   departments: [
//    id: 12,
//    name: '儿科门诊'
//   ]
//  ]
// }
type Hospital struct {
	Name string
}

type Handler struct{}

func Index(*http.Request) (data []byte, err error) {
	data, err = json.Marshal(Hospital{"H"})
	return
}

func (h Handler) Show() {
}

func (h Handler) Create() {
}

func (h Handler) Edit() {
}

func (h Handler) New() {
}

func Destroy() {
}

func Update() {
}
