/////////////////////////////
//Author: Mina Andrawos//////
////////////////////////////
	
package GoModbusTCP

import (
	"fmt"
	"encoding/binary"
	"errors"
)

//Input Register
type IRMessage struct{
	Seq uint16
}

func(R *IRMessage)ConstructWriteMessage(unitAddr byte, regAddr uint16, value uint16)(b []byte, err error) {
	return  nil, errors.New("Input registers do not allow writes..")
}

func (R *IRMessage)ConstructReadMessage(unitAddr byte, regAddr uint16, length uint16) (b []byte, err error){
	addrSlice := make([]byte,2)
	lenSlice := make([]byte, 2)
	binary.BigEndian.PutUint16(addrSlice, regAddr)
	binary.BigEndian.PutUint16(lenSlice, length)
	ReadMsg := []byte{0x00,0x00,0x00,0x00,0x00,0x06, unitAddr, 0x03 ,addrSlice[0], addrSlice[1],lenSlice[0], lenSlice[1]}
	fmt.Println(ReadMsg)
	return ReadMsg, nil
}

func (R *IRMessage) Readmessage(data []byte)(interface {}, error){
	return nil,nil
}
