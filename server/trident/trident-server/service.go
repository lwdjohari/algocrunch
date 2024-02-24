package main

import (
	"fmt"
	tc "trident-core"
)

func ServiceInit() {
	fmt.Print("Hello from Foo: ")
	fmt.Println(tc.Foo(1, 2))
}
