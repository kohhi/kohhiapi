package main

import (
	"flag"
	"kohhiapi/router"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 5134, "Port number for kohhi api.")
	flag.Parse()
	router.Router(port)
}
