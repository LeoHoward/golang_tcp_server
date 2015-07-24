package handler

import(
	"fmt"
)

func ParsePbMsg(opCode uint32, pbData []byte) (err error){

	fmt.Println(opCode)
	fmt.Println(pbData)

	return 
}

