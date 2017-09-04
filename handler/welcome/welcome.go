package welcome

import (
	"encoding/json"
	"net/http"
)

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
