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

	chatMux := make(chan string)
	listeners := make([](chan string), 0, 10)

	go func() {

		for {

			msg := <-chatMux

			// TODO: switch on adding a new listener

			for i := 0; i < len(listeners); i += 1 {

				listeners[i] <- msg

			}

		}

	}()
	
	broadcast := func(ws *websocket.Conn) {
		
		rcv := make(chan string)
		
		listeners = append(listeners, rcv)
		
		go func() {
			
			for {
				
				msg := <- rcv
				err := websocket.Message.Send(ws, msg)
				check(err)
				
			}
			
		}()
		
		for {
			
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			check(err)
			
			chatMux <- msg
			
		}
		
	}
	
	return upgradeWebsocketHandler(websocket.Handler(broadcast))
}
