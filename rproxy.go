package rproxy

import (
	"net"

	"github.com/soheilhy/cmux"
)

// Run start running reverse proxy
func Run(l net.Listener) {
	cl := cmux.New(l)
	registerHTTPServer(cl)
	registerGRPCServer(cl)
	cl.Serve()
}
