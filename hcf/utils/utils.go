package utils

import (
	"github.com/df-mc/dragonfly/server"
)

var srv *server.Server

func Server() *server.Server {
	return srv
}

func SetServer(srvC *server.Server) {
	srv = srvC
}