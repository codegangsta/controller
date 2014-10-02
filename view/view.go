package view

import (
	"net/http"

	"github.com/codegangsta/controller"
	"gopkg.in/unrolled/render.v1"
)

var Renderer = render.New(render.Options{})

type ViewController struct {
	controller.Base
	View     map[string]interface{}
	renderer *render.Render
}

func (c *ViewController) Init(rw http.ResponseWriter, r *http.Request) error {
	c.renderer = Renderer
	c.View = make(map[string]interface{})
	return c.Base.Init(rw, r)
}

func (c *ViewController) HTML(code int, name string, opts ...render.HTMLOptions) {
	c.renderer.HTML(c.ResponseWriter, code, name, c.View, opts...)
}
