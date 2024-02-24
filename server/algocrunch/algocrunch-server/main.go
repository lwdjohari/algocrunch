package main

import (
	_ "context"
	_ "log"
	_ "net"

	_ "google.golang.org/grpc"
)

func main() {
	ServiceInit()
}
