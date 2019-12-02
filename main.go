package main

import (
	"log"

	"github.com/docker/go-plugins-helpers/authorization"
)

const (
	pluginUnixSocket = "/run/docker/plugins/docker-authz.sock"
	rootGid          = 0
)

func main() {
	p, err := NewPlugin()

	if err != nil {
		log.Fatal(err)
	}

	h := authorization.NewHandler(p)
	if err := h.ServeUnix(pluginUnixSocket, rootGid); err != nil {
		log.Fatal(err)
	}
}
