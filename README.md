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
```

