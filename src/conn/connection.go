package conn

import(
	"net"
	"fmt"
	"bytes"
	"encoding/binary"

	"handler"
)

const(
	MAX_RECV_QUEUE_SIZE = 1024
    MAX_SEND_QUEUE_SIZE = 1024
    MAX_RECV_BUFF_SIZE  = 65536
	MSG_HEADER_LEN = 4
	OP_CODE_LEN = 4
	CACHE_SIZE = 1024
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

	cache := make([]byte, CACHE_SIZE)

	buf := bytes.NewBuffer(make([]byte, 0, MAX_RECV_BUFF_SIZE))

	var contentLen uint32

	for{

		size, err := this.conn.Read(cache)

        if err != nil {
            fmt.Printf("Read error, %v\n", err.Error())
            break
        }

		//把cache读取的内容写到buf
		buf.Write(cache[:size])

		for{

			//本次缓冲区数据包正好读完，重置内容长度
            if buf.Len() == 0 {
                contentLen = 0
                break
            }

            // 开始读取一个新的数据包
            if contentLen == 0 {
                // 判断缓冲区剩余数据是否足够读取一个包长
                if buf.Len() < MSG_HEADER_LEN {
                    break
                }
                packByteSize := make([]byte, MSG_HEADER_LEN)
                _, err = buf.Read(packByteSize)
                contentLen = binary.BigEndian.Uint32(packByteSize)
            }

			//判断缓冲区剩余数据是否足够读取一个完整的包
			//true -> 继续读取(contentLen - buf.Len())长度的字节数据
            if int(contentLen) > buf.Len() || contentLen == 0 {
                break
            }

            data := make([]byte, contentLen)
			//data为完整数据包
            _, err = buf.Read(data)

			opCode := binary.BigEndian.Uint32(data[:OP_CODE_LEN])

			pbData := data[OP_CODE_LEN:]

            //fmt.Println("opcode:", opCode)
			//fmt.Println("pb:", pbData)

            contentLen = 0

			handler.ParsePbMsg(opCode, pbData)

		}
	}
}

func (this *Connection) onSend(){

}

