package status

import (
	"github.com/gorilla/websocket"
	"container/list"
	"github.com/astaxie/beego"
	"encoding/json"
)

type Subscriber struct {
	Id   string
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	// Channel for new join users.
	subscribe = make(chan Subscriber, 10)
	// onStatusChange message here.
	messages = make(chan StatusData, 10)
	//the users that will be removed (disconnected usually)
	unSubscribe = make(chan string, 10)
	// all connected users.
	subscribers = list.New()
)

func webSocketListener() {
	for {
		select {
		case subs := <-subscribe:
		//entry,todo check weather it is the same connection
			subscribers.PushBack(subs)
			beego.Info("new webSocket connection")
		case msg := <-messages:
		//broadcast to all webSocket
			data, err := json.Marshal(msg)
			if err != nil {
				beego.Error("Fail to marshal event:", err)
				return
			}
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				// Immediately send event to WebSocket users.
				ws := sub.Value.(Subscriber).Conn
				if ws != nil {
					if ws.WriteMessage(websocket.TextMessage, data) != nil {
						// User disconnected.
						unSubscribe <- sub.Value.(Subscriber).Id
					}
				}
			}
		case unSubs := <-unSubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Id == unSubs {
					subscribers.Remove(sub)
					beego.Info("webSocket connection closed")
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
					}
					break
				}
			}

		}

	}
}

//join the web socket connection
func Entry(id string, conn *websocket.Conn) {
	subscribe <- Subscriber{Conn:conn, Id:id}
}

//remove web socket connection
func Leave(id string) {
	unSubscribe <- id
}
