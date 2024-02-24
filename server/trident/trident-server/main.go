package main

import (
	_ "context"
	"fmt"
	_ "log"
	_ "net"
	tc "trident-core"

	_ "google.golang.org/grpc"
)

func main() {
	fmt.Print("Hello from Foo: ")
	fmt.Println(tc.Foo(1, 2))
}
