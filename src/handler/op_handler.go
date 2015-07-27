package handler

import(
	"fmt"
	//"pb"
	"encoding/binary"
	"code.google.com/p/goprotobuf/proto"
)

const(
	MSG_HEADER_LEN = 4
	OP_CODE_LEN = 4
	HEADER_OPCODE_LEN = 8
)

func ParsePbMsg(opCode uint32, pbData []byte){

	//fmt.Println(opCode)
	//pbNew := &df_1001.Cs_10010001{
	//	OnlyId: proto.String("xiefan"),
	//}
	//data, _ := EncodePbMsg(10010001, pbNew)
	//fmt.Println(data)
}

func EncodePbMsg(opCode uint32, pb proto.Message) (packBuf []byte, err error){
	//包头(4字节) + opcode(4字节) + pbdata 
	//包头长度 = 4 + len(pbdata)
	pbBuf, err := proto.Marshal(pb)
	fmt.Println(pbBuf)
	if(err != nil){
		err = fmt.Errorf( "encode message to falied. %s", err.Error());
		return	
	}
	//pb长度
	pbBufSize := len(pbBuf)
	//(opcode + pb)长度
	packSize := uint32(OP_CODE_LEN + pbBufSize)
	packBuf = make([]byte, MSG_HEADER_LEN + packSize)

	binary.BigEndian.PutUint32(packBuf[:MSG_HEADER_LEN], packSize)
	binary.BigEndian.PutUint32(packBuf[MSG_HEADER_LEN:MSG_HEADER_LEN+OP_CODE_LEN], opCode)
	copy(packBuf[HEADER_OPCODE_LEN:], pbBuf)
	return
}


