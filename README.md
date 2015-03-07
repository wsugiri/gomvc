# gomvc
  gomvc is a simple Web Mvc libs for go

##Installation

installation
```go
go get github.com/wsugiri/gomvc
```

how to use main
```go
package main

import (
	"github.com/wsugiri/gomvc"
	"myapps/controllers"
)

func init() {
	gomvc.SetConfig("port", "9090")
}

func main() {
	gomvc.RouteFolder("/", "www")
	gomvc.Route("layout", &controllers.LayoutController{})
	gomvc.Route("process", &controllers.ProcessController{})

	gomvc.Run()
}
```

how to use controllers
```go
package controllers

import (
	"github.com/wsugiri/gomvc"
	"net/http"
)

type DemoController struct {
	gomvc.Controller
}

func (c *DemoController) Text(rw http.ResponseWriter, req *http.Request) {
	c.ServeText("<h1>Demo</h1> serve <b>text</b>")
}

func (c *DemoController) Html(rw http.ResponseWriter, req *http.Request) {
	c.ServeHtml("<h1>Demo</h1> serve <b>html</b>")
}

func (c *DemoController) View(rw http.ResponseWriter, req *http.Request) {
	c.ServeView("view.html", nil)
}

func (c *DemoController) Json(rw http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["firstName"] = "Peter"
	data["lastName"] = "Parker"
	data["alias"] = "Spiderman"
	c.ServeJson(data)
}

func (c *DemoController) Template(rw http.ResponseWriter, req *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Nexigo"
	c.ServeTpl([]string{"views/app/home.html", "views/header.tpl", "views/footer.tpl"}, data)
}
```



