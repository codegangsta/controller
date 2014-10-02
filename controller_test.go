package controller

import (
	"reflect"
	"testing"
)

type TestController struct {
	Base
}

func (t *TestController) Index() error {
	return nil
}

func (t *TestController) BadAction() {
}

func (t *TestController) BadAction2() string {
	return ""
}

type NoController struct {
}

func (n *NoController) Foo() {
}

func TestValidateAction(t *testing.T) {
	var validateTests = []struct {
		action interface{}
		valid  bool
	}{
		{(*TestController).Index, true},
		{(*TestController).BadAction, false},
		{(*TestController).BadAction2, false},
		{(*NoController).Foo, false},
		{"bad", false},
	}

	for _, test := range validateTests {
		_, err := controllerType(reflect.ValueOf(test.action))
		if (err == nil) != test.valid {
			t.Errorf("Action: %v should be valid=%v but returned error %v", reflect.ValueOf(test.action), test.valid, err)
		}
	}
}
