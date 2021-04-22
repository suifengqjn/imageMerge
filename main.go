package main

import (
	"flag"
	"myProject/imageMerge/GUI"
)

var (
	addr    = flag.String("addr", "127.0.0.1:4000", "address to start the server on")
	appName = flag.String("appName", "main", "")
)

func main() {

	flag.Parse()
	GUI.StartServer(*appName, *addr)
}