package main

import (
	"flawless_reserve_service/flawless_reserve_service/sys"
)

func main() {
	println("Start flawless_reserve_service")

	sys.SystemInit()
	sys.StartEchoFramework()
}
