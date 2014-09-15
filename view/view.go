package view

import (
	"net/http"

	"github.com/codegangsta/ctrl"
	"gopkg.in/unrolled/render.v1"
)

var Renderer = render.New(render.Options{})

type ViewController struct {
	ctrl.Base
	View     map[string]interface{}
	renderer *render.Render
}

func (c *ViewController) Init(rw http.ResponseWriter, r *http.Request) {
	c.Base.Init(rw, r)
	c.renderer = Renderer
	c.View = make(map[string]interface{})
}

func (c *ViewController) HTML(code int, name string, opts ...render.HTMLOptions) {
	c.renderer.HTML(c.ResponseWriter, code, name, c.View, opts...)
}
