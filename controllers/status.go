package controllers

import (
	"log"
	"net/http"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"gensh.me/VirtualJudge/components/context/status"
)

type StatusController struct {
	BaseController
}

func (c *StatusController)Status() {
	c.Data["json"] = `{"id":0}`
	c.ServeJSON()
}

func (c *StatusController)Test() {
	status.Test()
	c.Data["json"] = `{"id":0}`
	c.ServeJSON()
}

func (c *StatusController)WebSocket() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	c.EnableRender = false
	status.Entry("0", ws) //todo :id
	defer status.Leave("")
	for {
		ws.WriteMessage(websocket.TextMessage, []byte("hello"))
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		log.Println(string(p))
		//publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
