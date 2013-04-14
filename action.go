package    main

import    (
        "encoding/json"
	"code.google.com/p/go.net/websocket"
)

type    gameAction    struct    {
        u,    v    int
}

func    actionServer(globe    *Geodesic)    interface{}    {

	broadcast    :=    make(chan    string)
	register    :=    make(chan    (chan    string))
	unregister    :=    make(chan    (chan    string))
	listeners    :=    make(map[(chan    string)]bool)

	go    func()    {
		
		for    {
			
			select    {
				
			case    rcv    :=    <-    register:
				
				listeners[rcv]    =    true
				
			case    rcv    :=    <-    unregister:
				
				delete(listeners,    rcv)
				
			case    msg    :=    <-broadcast:
				
				//    TODO:    switch    on    adding    a    new    listener
			        
                                var    action    gameAction

                                err    :=    json.Unmarshal([]byte(msg),    action)
                                if    err    !=    nil    {
                                        continue
                                }

                                node    :=    globe.U_Array[action.u][action.v]
                                space    :=    node.Space
                                space.PlayerID    =    (space.PlayerID    +    1)    %    3

                                rsp,    err    :=    json.Marshal(node)
                                check(err)

				for    l    :=    range    listeners        {	

					l    <-    string(rsp)
					
				}
				
			}
			
		}

	}()
	
	actionHandler    :=    func(ws    *websocket.Conn)    {
		
		rcv    :=    make(chan    string)
		
		register    <-    rcv
		
		defer    func()    {    unregister    <-    rcv    }()
		
		go    func()    {
			
			for    {
				
				msg    :=    <-    rcv
				err    :=    websocket.Message.Send(ws,    msg)
				check(err)
				
			}
			
		}()
		
		for    {
			
			var    msg    string
			err    :=    websocket.Message.Receive(ws,    &msg)
			check(err)
			
			broadcast    <-    msg
			
		}
		
	}
	
	return    upgradeWebsocketHandler(websocket.Handler(actionHandler))

}
