package gomvc

import (
	"golang.org/x/net/websocket"
	"net/http"
	"reflect"
	"strings"
)

func Socket(path string, ctl interface{}, checkOrigin bool) {
	if path[:1] != "/" {
		path = "/" + path
	}

	if path[len(path)-1:len(path)] != "/" {
		path = path + "/"
	}

	if checkOrigin {
		http.Handle(path, websocket.Handler(func(ws *websocket.Conn) {
			rqpath := strings.TrimSpace(ws.Request().URL.Path)
			action := strings.ToLower(strings.Split(rqpath[len(path):], "/")[0])
			typeCont := reflect.TypeOf(ctl)
			for i := 0; i < typeCont.NumMethod(); i++ {
				method := strings.ToLower(typeCont.Method(i).Name)
				if method == action {
					reflect.ValueOf(ctl).Method(i).Call([]reflect.Value{reflect.ValueOf(ws)})
					break
				}
			}
		}))
	} else {
		http.HandleFunc(path,
			func(w http.ResponseWriter, req *http.Request) {
				s := websocket.Server{Handler: websocket.Handler(func(ws *websocket.Conn) {
					rqpath := strings.TrimSpace(ws.Request().URL.Path)
					action := strings.ToLower(strings.Split(rqpath[len(path):], "/")[0])
					typeCont := reflect.TypeOf(ctl)
					for i := 0; i < typeCont.NumMethod(); i++ {
						method := strings.ToLower(typeCont.Method(i).Name)
						if method == action {
							reflect.ValueOf(ctl).Method(i).Call([]reflect.Value{reflect.ValueOf(ws)})
							break
						}
					}
				})}
				s.ServeHTTP(w, req)
			})
	}
}

// func (s *GomvcSocket) Close() error {
// 	var err error
// 	if err = s.websocket.Close(); err != nil {
// 		log.Println("Websocket could not be closed", err.Error())
// 		return err
// 	}
// 	return nil
// }

// func (s *GomvcSocket) SendString(message string) error {
// 	var err error
// 	if err = Message.Send(s.websocket, message); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *GomvcSocket) BroadcastString(message string) error {
// 	var err error
// 	ActiveClients[*s] = ""
// 	for a, _ := range ActiveClients {
// 		if err = Message.Send(a.websocket, message); err != nil {
// 			log.Println("Could not send message to ", a.clientIP, err.Error())
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (s *GomvcSocket) SendByte(file []byte) error {
// 	var err error
// 	if err = Message.Send(s.websocket, file); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *GomvcSocket) BroadcastByte(file []byte) error {
// 	var err error
// 	ActiveClients[*s] = ""
// 	for a, _ := range ActiveClients {
// 		if err = Message.Send(a.websocket, file); err != nil {
// 			log.Println("Could not send message to ", a.clientIP, err.Error())
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (s *GomvcSocket) Receive() (string, error) {
// 	var err error
// 	var clientMessage string
// 	if err = Message.Receive(s.websocket, &clientMessage); err != nil {
// 		Message.Send(s.websocket, fmt.Sprintf("ERROR: %s", err.Error()))
// 		return "", err
// 	}
// 	return clientMessage, nil
// }
