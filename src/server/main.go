package main

import(
	"conn"
)

func main(){
	
	conn.NewTcpServer("0.0.0.0", 55555)
}
