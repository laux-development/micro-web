package micro_web

import (
	"net/http"

	v "github.com/laux-development/micro-view"
)

type web struct{
	data interface{}
}

NewWeb(data interface{}) *web{
	return &web{data:data}
}

func (we *web) Home(w http.ResponseWriter, r *http.Request) {
	v.Home(w)
}

func (we *web) Login(w http.ResponseWriter, r *http.Request) {
	v.Login(w)
}

func (we *web) Profile(w http.ResponseWriter, r *http.Request) {
	v.Profile(w, we.data)
}
