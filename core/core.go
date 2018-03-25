package core

import (
	"net"
	"fmt"
	"github.com/smallfish/ftp"
)
// 端口扫描 基本原理
// 如果与某个端口能建立扫描连接

func CheckPortCThread(ip net.IP, port int,c chan int  )  {
	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		if err:=recover();err!=nil{
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	CheckPort(ip,port)
	<-c
}

func  CheckPort(ip net.IP, port int) {
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if conn !=nil{
		fmt.Println(tcpAddr.IP,":",tcpAddr.Port,"连接成功","端口开放")
		conn.Close()
		LoginFtp(tcpAddr)
	}
	if err != nil {
	//	fmt.Println(tcpAddr.IP,":",tcpAddr.Port,"端口连接失败","端口关闭")
	//	fmt.Println(err)
	}
}

func LoginFtp(tcpAddr net.TCPAddr) (error) {
	ftp := new(ftp.FTP)
	//ftp.Debug=true
	ftp.Connect(tcpAddr.IP.String(), 21)
	if ftp==nil {
		return nil
	}
	ftp.Login("ftp", "123456")
	if ftp.Code == 230 {
		fmt.Println(ftp.Code)
		fmt.Println("扫描成功："+tcpAddr.IP.String()+"密码：123456")
		return nil
	}
	ftp.Close()
	return nil
}
