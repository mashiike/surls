package main

import (
	"github.com/fujiwara/ridge"
	"github.com/mashiike/surls"
)

func main() {
	conf := surls.NewDefaultConfig()
	mux := surls.NewServeMux(conf)
	ridge.Run(":8080", "/", mux)
}
