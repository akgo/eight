# eight

## What is it
一个熟悉golang的小样例，websocket,http,reflect,map,go,channl,interface ,json 

## Requires
"github.com/gorilla/websocket"
"github.com/satori/go.uuid"
"gopkg.in/ini.v1"



## Basic Usage

### A websocket server 
```go
package main

import (

	"fmt"
	"net/http"
	"os"

	"eight/model"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"gopkg.in/ini.v1"
	"io"
)

var manager = model.ClientManager{
	Broadcase: make(chan []byte),
	Register: make(chan * model.Client),
	Unregister: make(chan *model.Client),
	Clients: make( map [*model.Client] bool),
}


func main () {
	fmt.Println("starting application")

	cfg, err := ini.Load("config/app.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	go manager.Start()

	//client
	http.HandleFunc("/", home)

	http.HandleFunc("/ws",wsPage)

	http.HandleFunc("/client", htmlServer)

	http.ListenAndServe(cfg.Section("server").Key("websocket_sercer_ip").String() + ":" + cfg.Section("server").Key("webscoket_port").String(),nil)

}

func wsPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &model.Client{Id: uuid.Must(uuid.NewV4()).String(), Socket: conn, Send: make(chan []byte)}

	manager.Register <- client

	go client.Read(&manager)
	go client.Write()

}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "home server!")
}

func htmlServer(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	wd = wd + "/client"
	fmt.Println(wd)

	http.StripPrefix("/client", http.FileServer(http.Dir(wd))).ServeHTTP(w, r)
}
```

