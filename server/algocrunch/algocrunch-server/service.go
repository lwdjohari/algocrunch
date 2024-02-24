package main

import (
	alc "algocrunch-core"
	"fmt"
)

func ServiceInit() {
	fmt.Print("Hello from Foo: ")
	fmt.Println(alc.Foo(1, 2))
}
