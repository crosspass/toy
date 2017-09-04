package config

import (
	"toy/handler/welcome"
	"toy/lib/router"
)

func Run() {
	router.Root(welcome.Index)
}
