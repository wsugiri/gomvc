package gomvc

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
	"strings"
)

var (
	defaultPort                             = "8080"
	defaultMasterPage                       = "views/share/template/master.html"
	HttpContext                             = ContextHandler{}
	configs                                 = make(map[string]interface{})
	driverDB                                = ""
	connInfo                                = ""
	Session           *sessions.CookieStore = nil
)

type RouteHandler struct {
	Path        string
	IController interface{}
}

type ContextHandler struct {
	routes []RouteHandler
}

func SetSession(secret string) *sessions.CookieStore {
	a := sessions.NewCookieStore([]byte(secret))
	return a
}

func GetSession(req *http.Request, sessionName string) (*sessions.Session, error) {
	ses, err := Session.Get(req, sessionName)
	if err != nil {
		return nil, err
	}
	return ses, nil
}

func Route(path string, ctl interface{}) {
	HttpContext.routes = append(HttpContext.routes, RouteHandler{path, ctl})
}

func RouteFolder(path string, folder string) {
	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(folder))))
}

func Run() {
	routes := HttpContext.routes

	for i := 0; i < len(routes); i++ {
		route := routes[i]
		path := strings.TrimSpace(route.Path)

		if path[:1] != "/" {
			path = "/" + path
		}

		if path[len(path)-1:len(path)] != "/" {
			path = path + "/"
		}

		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			rqpath := strings.TrimSpace(r.URL.Path)
			action := strings.ToLower(strings.Split(rqpath[len(path):], "/")[0])

			if action == "" {
				action = "index"
			}

			if action != "favicon.ico" {
				typeCont := reflect.TypeOf(route.IController)
				for i := 0; i < typeCont.NumMethod(); i++ {
					method := strings.ToLower(typeCont.Method(i).Name)
					if method == action {
						reflect.ValueOf(route.IController).Method(i).Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
						ctrlObj := reflect.ValueOf(route.IController).Elem().Field(0).Interface().(Controller)
						ctrlObj.RunAction(w, r)
						break
					}
				}
			} else {
				// fmt.Println("... favicon ...")
			}

		})
	}
	err := http.ListenAndServe(":"+defaultPort, nil)
	if err != nil {
		panic(err)
	}
}

func SetConfig(name string, value interface{}) {
	configs[name] = value

	switch name {
	case "port":
		defaultPort = value.(string)
	case "driverdb":
		driverDB = value.(string)
	case "conninfo":
		connInfo = value.(string)
	case "session":
		Session = value.(*sessions.CookieStore)
	case "masterpage":
		defaultMasterPage = value.(string)
	}
}

func GetConfig(name string) interface{} {
	return configs[name]
}

func GetFormValue(r *http.Request) map[string]interface{} {
	r.ParseForm()
	result := make(map[string]interface{})
	for k, v := range r.Form {
		result[k] = v[0]
	}
	return result

}

func Text() {
	text := " /demo test/a ska     "
	fmt.Println("...", reflect.TypeOf(HttpContext), strings.TrimSpace(text))
}
