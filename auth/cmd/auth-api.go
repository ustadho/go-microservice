package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ustadho/go-microservice/auth/models"
	"github.com/ustadho/go-microservice/auth/routers"
)

var (
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)

func main() {
	flag.Parse()
	flags := models.NewFlags(*ip, *port)

	fmt.Println("Starting ip")

	logger := log.New(os.Stdout, "auth", 1)
	route := routers.NewRoute(logger, flags)
	engine := route.RegisterRoutes()
	url, _ := flags.GetApplicationUrl()
	engine.Run(*url)
}
