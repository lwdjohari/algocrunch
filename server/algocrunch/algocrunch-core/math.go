package algocrunchcore

import (
	nc "nvm-gocore"
)

func Foo(a int, b int) nc.Option[int] {
	return nc.Some[int](a + b)
}
