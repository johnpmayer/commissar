package    main

import    (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"io"
	"os"
	"strconv"
)

func    check(err    error)    {
	if    err    !=    nil    {
		panic(err.Error())
	}
}

func    gamePageHandler(ctx    *web.Context)    {
	gamePage,    err    :=    os.Open("index.html")
	check(err)
	ctx.ContentType("html")
	_,    err    =    io.Copy(ctx,    gamePage)
	check(err)
}

func    scriptHandler(ctx    *web.Context,    path    string)    {
	file,    err    :=    os.Open(path)
	check(err)
	ctx.ContentType("js")
	_,    err    =    io.Copy(ctx,    file)
	check(err)
}

func    pngHandler(ctx    *web.Context,    path    string)    {
	file,    err    :=    os.Open(path)
	check(err)
	ctx.ContentType("png")
	_,    err    =    io.Copy(ctx,    file)
	check(err)
}

func    vertexShaderHandler(ctx    *web.Context,    path    string)    {
	file,    err    :=    os.Open(path)
	check(err)
	ctx.ContentType("x-shader/x-vertex")
	_,    err    =    io.Copy(ctx,    file)
	check(err)
}

func    fragmentShaderHandler(ctx    *web.Context,    path    string)    {
	file,    err    :=    os.Open(path)
	check(err)
	ctx.ContentType("x-shader/x-fragment")
	_,    err    =    io.Copy(ctx,    file)
	check(err)
}

func    main()    {

	globe    :=    MakeGeodesic(1,    3)

	globeHandler    :=    func(ctx    *web.Context)    {

		obj,    err    :=    json.Marshal(globe)
		check(err)
		ctx.ContentType("json")
		_,    err    =    ctx.Write(obj)
		check(err)

	}

	clickHandler    :=    func(ctx    *web.Context)    {
		
		u,    err    :=    strconv.Atoi(ctx.Params["u"])
		check(err)
		
		v,    err    :=    strconv.Atoi(ctx.Params["v"])
		check(err)
		
		fmt.Println(u,    v)

		node    :=    globe.U_Array[u][v]

		fmt.Println(node)

		var    result    []byte

		space    :=    node.Space

		if    node.Space    ==    nil    {
			space    =    &BoardSpace{PlayerID:    0,    Armies:    0}
			node.Space    =    space
		}

                space.PlayerID    =    (space.PlayerID    +    1)    %    3

		result,    err    =    json.Marshal(space)
		check(err)
		ctx.ContentType("json")
		_,    err    =    ctx.Write([]byte(result))
		check(err)

	}
	
	authHandler    :=    func(ctx    *web.Context)    {
		
		value,    hasCookie    :=    ctx.GetSecureCookie("user")
		
		fmt.Println(hasCookie,    value)
		
	}
	
	server    :=    web.NewServer()
	
	server.Config.CookieSecret    =    "todo~commissar165412399"
	
	//    Static    routers
	server.Get("/",    gamePageHandler)
	server.Get("/(images/.*[.]png)",    pngHandler)
	server.Get("/(scripts/.*[.]js)",    scriptHandler)
	server.Get("/(shaders/.*[.]vert)",    vertexShaderHandler)
	server.Get("/(shaders/.*[.]frag)",    fragmentShaderHandler)
	
	//    Serve    globe    terrain
	server.Get("/globe",    globeHandler)
	
	//    What    do    do    when    we    clicked    on    a    board    space
	server.Post("/click",    clickHandler)
	server.Post("/auth",    authHandler)
	
	server.Get("/echo",    chatServer())
        server.Get("/action",    actionServer(globe))
	
	server.Run(":8080")
		
}

