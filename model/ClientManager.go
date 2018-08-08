package model

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"reflect"
)

type Client struct {
	Id string
	Socket *websocket.Conn
	Send chan []byte
}

type ClientManager struct {
	Clients map[*Client]bool
	Broadcase chan []byte
	Register  chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content string `json:"content,omitempty"`
}

func (manager * ClientManager) Start () {
	for {

		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage,_ := json.Marshal(&Message{Content: "/A new sockets has connected"})
			manager.Send(jsonMessage,conn)
		case conn := <-manager.Unregister:
			if _,ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients,conn)
				jsonMessage,_ := json.Marshal(&Message{Content:"/A sockets has disconnected"})
				manager.Send(jsonMessage,conn)
			}
		case message := <-manager.Broadcase:
			//解析广播 是群播，还是自己给自己

			var m Message
			json.Unmarshal([]byte(message), &m)

			if modelManage.isBind(m.Content) {
				modelManage.Exec(m.Content,[]reflect.Value{})
			}
		
			for conn := range manager.Clients {
				select {
				case conn.Send <- message :
				default:
					close(conn.Send)
					delete(manager.Clients,conn)
				}
			}

		}

	}

}

func (manager *ClientManager) Send (message []byte,ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}

}

func (c *Client) Read (manage *ClientManager) {
	defer func () {
		manage.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_,message ,err := c.Socket.ReadMessage()
		if err != nil {
			manage.Unregister <- c
			c.Socket.Close()
			break
		}

		jsonMessage , _:= json.Marshal(&Message{Sender:c.Id,Content:string(message)})
		manage.Broadcase <- jsonMessage
	}

}

func (c *Client) Write () {
	defer func () {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}