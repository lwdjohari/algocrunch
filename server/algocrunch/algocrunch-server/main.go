package main

import (
	tc "algocrunch-core"
	_ "context"
	"fmt"
	_ "log"
	_ "net"

	_ "google.golang.org/grpc"
)

func main() {
	fmt.Print("Hello from Foo: ")
	fmt.Println(tc.Foo(1, 2))
}
