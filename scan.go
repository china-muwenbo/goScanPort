package main


import (
	 "strings"
	"strconv"
	"./core"
	"fmt"
	"net"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	maxchan:=make(chan  int ,512 )
	ip:=processIptest("123.207.5.204","123.207.215.222")
	fmt.Println(ip)
	for i,v:=range ip{
		maxchan<-i
	go 	core.CheckPortCThread(net.ParseIP(v), 21,maxchan)

	}
}
//生成IP地址列表
func processIptest(startIp,endIp string) []string{
	var ips = make([]string,0)
	for ;startIp != endIp;startIp = nextIptest(startIp){
		if startIp != ""{
			ips = append(ips,startIp)
		}
	}
	ips = append(ips,startIp)
	return ips
}
func nextIptest(ip string) string{
	ips := strings.Split(ip,".")
	var i int;
	for i = len(ips) - 1;i >= 0;i--{
		n,_ := strconv.Atoi(ips[i])
		if n >= 255{
			//进位
			ips[i] = "0"
		}else{
			//+1
			n++
			ips[i] = strconv.Itoa(n)
			break
		}
	}
	if i == -1{
		//全部IP段都进行了进位,说明此IP本身已超出范围
		return "";
	}
	ip = ""
	leng := len(ips)
	for i := 0;i < leng;i++{
		if i == leng -1{
			ip += ips[i]
		}else{
			ip += ips[i] + "."
		}
	}
	return ip
}
