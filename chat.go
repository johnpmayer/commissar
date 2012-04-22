package main

import (
	"fmt"
	//	"github.com/hoisie/web"
	"code.google.com/p/go.net/websocket"
//	"io"
	"web"
)

func echo(ws *websocket.Conn) {

	fmt.Println("Opened websocket")

	for {
		
		var msg string
		var err error
		
		err = websocket.Message.Receive(ws, &msg)
		check(err)
		
		fmt.Println(msg)
		
		err = websocket.Message.Send(ws, msg)
		check(err)
		
	}

	fmt.Println("Finished copying websocket")
}

func upgradeWebsocketHandler(wsHandler websocket.Handler) interface{} {
	return func(ctx *web.Context) {
		wsHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	}
}

func chatServer() interface{} {
	return upgradeWebsocketHandler(websocket.Handler(echo))
}
