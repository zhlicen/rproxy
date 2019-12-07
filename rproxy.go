package rproxy

import (
	"log"
	"net"

	"github.com/soheilhy/cmux"
)

// Run start running reverse proxy
func Run(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	cl := cmux.New(l)
	registerHTTPServer(cl)
	registerGRPCServer(cl)
	cl.Serve()
}
