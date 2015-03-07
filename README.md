# gomvc
  gomvc is a simple Web Mvc libs for go

##Installation

installation
```go
go get github.com/wsugiri/gomvc
```

how to use
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


