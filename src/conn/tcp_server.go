package conn

import(
	"net"
	"fmt"
	"container/list"
)

type ITcpServer interface{
	Start(string, uint16) bool
	//Stop
}


type TcpServer struct{
	addr *net.TCPAddr
	listener *net.TCPListener
	conn_list *list.List	
}


func NewTcpServer(ip string, port uint16) *TcpServer{
	
	tcp_server := new(TcpServer)
	tcp_server.conn_list = list.New()

	tcp_server.Start(ip, port)
		
	return tcp_server
}

func (this *TcpServer) Start(ip string, port uint16) bool{

	if this.listener != nil{
		fmt.Println("TcpServer Listener exist")
		return false
	}

	addr, err := net.ResolveTCPAddr( "tcp4", fmt.Sprintf( "%s:%d", ip, port ) )
	
    if err != nil {
        fmt.Println( "ResolveTCPAddr failed. ", err.Error() );
        return false
    }

	this.addr = addr

	listener, err := net.ListenTCP("tcp", this.addr)

    if err != nil {
        fmt.Println( "TCPListener failed. ", err.Error() );
        return false
    }

	this.listener = listener

	defer listener.Close()

	for{
		
		conn, err := this.listener.Accept();

		if err != nil{
			fmt.Println("Accept close")
			break
		}

		go NewConnection(this, conn)

		this.conn_list.PushBack(conn)
	}
	return true
}

