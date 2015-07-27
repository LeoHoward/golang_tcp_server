package handler

import(
	"fmt"
	//"pb"
	//"code.google.com/p/goprotobuf/proto"
)

func ParsePbMsg(opCode uint32, pbData []byte){

	fmt.Println(opCode)
	fmt.Println(pbData)

	//pbNew := &df_1001.Cs_10010001{
	//	OnlyId : proto.String("xiefan")
	//}

	//buf, err := proto.Marshal(pbNew)

	//fmt.Println(buf)
}

