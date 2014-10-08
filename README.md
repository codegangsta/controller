# controller
Package controller is a lightweight and composable controller implementation
for net/http

Sometimes plain net/http handlers are not enough, and you want to have logic
that is resource/concept specific, and data that is request specific.

This is where controllers come into play. Controllers are structs that
implement a specific interface related to lifecycle management. A Controller
can also contain an arbitrary amount of methods that can be used as handlers to
incoming requests. This package makes it easy to automatically construct a new
Controller instance and invoke a specified method on that controller for every
request.

## Example

``` go
package main

import (
  "net/http"

  "github.com/codegangsta/controller"
)

type MyController struct {
  controller.Base
}

func (c *MyController) Index() error {
  c.ResponseWriter.Write([]byte("Hello World"))
  return nil
}

func main() {
  http.Handle("/", controller.Action((*MyController).Index))
  http.ListenAndServe(":3000", nil)
}

```
