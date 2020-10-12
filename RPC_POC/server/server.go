package main

import (
	"RPC_POC/common"
	"io"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	//Create a `*College` object.
	mit := common.NewCollege()

	//Register `mit` object  with `rpc.DefaultServer`
	//Register publishes the receiver's methods in the DefaultServer.
	rpc.Register(mit)

	// 	HandleHTTP registers an HTTP handler for RPC messages to DefaultServer on
	// DefaultRPCPath and a debugging handler on DefaultDebugPath. It is still
	// necessary to invoke http.Serve(), typically in a go statement.
	/*Register an HTTP handler for RPC communication on `http.DefaultServerMux`(default)
	regsiters a handler on the `rpc.DefaultRPCPath` endpoint  to respond to RPC messages
	registers a hander*/
	rpc.HandleHTTP()

	//Sample test ennpoint.
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC Server Live")
	})
	log.Println("Server Started:")
	//Listen and serve default HTTP server.
	http.ListenAndServe(":9000", nil)
}
