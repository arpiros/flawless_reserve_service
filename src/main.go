package main

import "reserve_service/sys"

func main() {
	println("Start reserve_service")

	sys.SystemInit()
	sys.StartEchoFramework()
}
