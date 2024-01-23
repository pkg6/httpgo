package httpgo

import (
	"github.com/pkg6/httpgo/binding"
	"net/http"
)

func DefaultBind(req *http.Request, obj any) error {
	return Bind(req, obj, binding.Default(req.Method, req.Header.Get("Content-Type")))
}

func Bind(req *http.Request, obj any, binding binding.Binding) error {
	return binding.Bind(req, obj)
}
