package micro_web

import (
	"net/http"

	v "github.com/laux-development/micro_view"
)

type web struct {
	view v.ViewProvider
	data interface{}
}

func NewWeb(view v.ViewProvider, data interface{}) *web {
	return &web{view: view, data: data}
}

func (we *web) Home(w http.ResponseWriter, r *http.Request) {
	we.view.Home(w)
}

func (we *web) Login(w http.ResponseWriter, r *http.Request) {
	we.view.Login(w)
}

func (we *web) Profile(w http.ResponseWriter, r *http.Request) {
	err := we.view.Profile(w, we.data)

	if err != nil {
		http.Error(w, "Fatal Error", 500)
	}
}
