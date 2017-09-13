package main

import (
	"encoding/json"
	"log"
	"net/http"
	"toy/config"
	"toy/lib/router"
)

func main() {
	config.Run()
	http.HandleFunc("/", h)
	log.Println(http.ListenAndServe(":8000", nil))
}

func h(w http.ResponseWriter, r *http.Request) {
	log.Printf("host: %s, url: %s, method: %s", r.RemoteAddr, r.URL.Path, r.Method)
	method, path := r.Method, r.URL.Path
	w.Header().Set("Content-Type", "application/json")

	var ret []byte
	var err error
	if router, ok := router.Routers[path]; ok {
		if path == "/" || path == "" {
			ret, err = router.RootHanlderFunc(r)
		} else {
			switch method {
			case "get":
				ret, err = router.Handler.Show(r)
			case "post":
				ret, err = router.Handler.Create(r)
			case "patch":
				ret, err = router.Handler.Update(r)
			case "put":
				ret, err = router.Handler.Update(r)
			case "delete":
				ret, err = router.Handler.Destroy(r)
			}
		}
		if err != nil {
			data := make(map[string]string)
			data["error"] = err.Error()
			bytes, err := json.Marshal(data)
			if err != nil {
				panic(err)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			w.Write(bytes)
			return
		} else {
			w.Write(ret)
			return
		}
	}
	data := make(map[string]string)
	data["error"] = "Resource don't found!"
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(bytes)
}
