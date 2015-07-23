package conn

import(
	"net"
	"fmt"
)

const(
	MAX_RECV_QUEUE_SIZE = 1024
    MAX_SEND_QUEUE_SIZE = 1024
    MAX_RECV_BUFF_SIZE  = 65536
)

type RecvQueue chan []byte
type SendQueue chan []byte

type Connection struct{
	tcp_server ITcpServer
	conn net.Conn
	recv RecvQueue
	send SendQueue
}

func NewConnection(tcp_server ITcpServer, conn net.Conn) *Connection{
	
	task := new(Connection)
	task.tcp_server = tcp_server
	task.conn = conn
	
	task.recv = make(RecvQueue, MAX_RECV_QUEUE_SIZE)
	task.send = make(SendQueue, MAX_SEND_QUEUE_SIZE)

	go task.onRecv()
	go task.onSend()
	
	return task
}


func (this *Connection) onRecv(){
	
	buffer := make([]byte, MAX_RECV_BUFF_SIZE)

	for{
	
		len, _ := this.conn.Read(buffer[0:])
		
		packet := buffer[:len]
		fmt.Println("onRecv:", buffer[:len])	

		this.send <- packet
	}
}

func (this *Connection) onSend(){

	for{
		
		msg, _ := <- this.send
		
		fmt.Println("onSend:", msg)
	}
}

