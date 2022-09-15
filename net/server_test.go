package net

import (
	"testing"
)

func TestServer(t *testing.T) {
	_ = Serve("8080")
}

func TestServer_StartAndServer(t *testing.T) {
	s := &Server{
		addr: "8080",
	}
	_ = s.StartAndServer()
}
