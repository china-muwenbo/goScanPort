package main

import (
	"./core"
	"net"
)
func main()  {
	testIP:=net.TCPAddr{}
	testIP.Port=21
	testIP.IP=net.ParseIP("123.207.215.205")
	core.LoginFtp(testIP)
}
