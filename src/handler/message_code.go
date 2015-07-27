package handler

import(
	"encoding/binary"
)

const(
	MSG_HEADER_LEN = 4
	OP_CODE_LEN = 4

	reqLogin 	=	10010001
	respLogin	=	10010002
)

func Wrap(opCode uint32, pbData []byte) (data []byte, err error){
	pbDataLen := uint32(len(pbData))
	packLen := MSG_HEADER_LEN + OP_CODE_LEN + pbDataLen
	buf := make([]byte, packLen)
	binary.BigEndian.PutUint32(buf[:MSG_HEADER_LEN], packLen)
	binary.BigEndian.PutUint32(buf[MSG_HEADER_LEN:MSG_HEADER_LEN+OP_CODE_LEN], opCode)
	copy(buf[MSG_HEADER_LEN+OP_CODE_LEN:], pbData)
	return
}


