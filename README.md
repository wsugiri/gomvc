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
)

func main() {
	gomvc.SetConfig("port", "9090")
	gomvc.Run()
}
```


