package ctrl

import (
	"errors"
	"net/http"
	"reflect"
)

var (
	interfaceType = interfaceOf((*Controller)(nil))
)

type Controller interface {
	Init(http.ResponseWriter, *http.Request)
	Destroy()
}

type Base struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func (b *Base) Init(rw http.ResponseWriter, r *http.Request) {
	b.Request, b.ResponseWriter = r, rw
}

func (b *Base) Destroy() {
}

func Action(action interface{}) http.Handler {
	val := reflect.ValueOf(action)
	t, err := controllerType(val)
	if err != nil {
		panic(err)
	}

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		v := reflect.New(t)
		c := v.Interface().(Controller)
		c.Init(rw, r)
		val.Call([]reflect.Value{v})
		c.Destroy()
	})
}

func controllerType(action reflect.Value) (reflect.Type, error) {
	t := action.Type()

	if t.Kind() != reflect.Func {
		return t, errors.New("Action is not a function")
	}

	if t.NumIn() != 1 {
		return t, errors.New("Wrong Number of Arguments in action")
	}

	t = t.In(0)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if !reflect.PtrTo(t).Implements(interfaceType) {
		return t, errors.New("Controller does not implement ctrl.Controller interface")
	}

	return t, nil
}

func interfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}
