package gomvc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	// "reflect"
	"strings"
)

type Controller struct {
	MsgType   string
	Html      string
	Text      string
	Templates []string
	Data      interface{}
}

func (c *Controller) New(mtype string) {
	c.MsgType = mtype
	c.Html = ""
	c.Text = ""
	c.Templates = make([]string, 0)
}

func (c *Controller) ServeHtml(html string) {
	c.New("html")
	c.Html = html
}

func (c *Controller) ServeText(text string) {
	c.New("text")
	c.Text = text
}

func (c *Controller) ServeTpl(tmpls []string, data interface{}) {
	c.New("tpl")
	c.Templates = append(c.Templates, defaultMasterPage)
	c.Templates = append(c.Templates, tmpls...)
	c.Data = data
}

func (c *Controller) ServeView(view string, data interface{}) {
	c.New("view")
	c.Templates = []string{"views/" + view}
	c.Data = data
}

func (c *Controller) ServeJson(data interface{}) {
	c.New("json")
	c.Data = data
}

func (c *Controller) Redirect(url string) {
	c.New("redirect")
	c.Text = url
}

func (c *Controller) RunAction(w http.ResponseWriter, r *http.Request) {
	switch c.MsgType {
	case "html":
		fmt.Fprintf(w, c.Html)
	case "text":
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(strings.TrimSpace(c.Text)))
	case "tpl":
		t := template.Must(template.ParseFiles(c.Templates...))
		t.Execute(w, c.Data)
	case "view":
		t := template.Must(template.ParseFiles(c.Templates...))
		t.Execute(w, c.Data)
	case "json":
		js, err := json.Marshal(c.Data)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	case "redirect":
		http.Redirect(w, r, c.Text, 301)
	default:
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("X-Content-Type-Options", "nosniff")
		fmt.Fprintf(w, r.URL.Path)
	}
}

func Test() {
	t, _ := template.ParseFiles("...")
	fmt.Println(t)
}
