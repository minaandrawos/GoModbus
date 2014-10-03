package GoModbusTCP

import (
	"fmt"
	"encoding/binary"
)

//Output Coil
type OutputCoil struct{
	Seq uint16
}

//TODO support multiple writes
//Single coil only
func(R *OutputCoil)ConstructWriteMessage(unitAddr byte, regAddr uint16, value uint16)(b []byte, err error) {
	addrSlice := make([]byte,2)
	valSlice := make([]byte, 2)
	binary.BigEndian.PutUint16(addrSlice, regAddr)
	binary.BigEndian.PutUint16(valSlice, value)
	writeMsg := []byte{0x00,0x00,0x00,0x00,0x00,0x06, unitAddr, 0x05 ,addrSlice[0], addrSlice[1],valSlice[0], valSlice[1]}
	fmt.Println(writeMsg)
	return  writeMsg, nil
}

func (R *OutputCoil)ConstructReadMessage(unitAddr byte, regAddr uint16, length uint16) (b []byte, err error){
	addrSlice := make([]byte,2)
	lenSlice := make([]byte, 2)
	binary.BigEndian.PutUint16(addrSlice, regAddr)
	binary.BigEndian.PutUint16(lenSlice, length)
	ReadMsg := []byte{0x00,0x00,0x00,0x00,0x00,0x06, unitAddr, 0x01 ,addrSlice[0], addrSlice[1],lenSlice[0], lenSlice[1]}
	fmt.Println(ReadMsg)
	return ReadMsg, nil
}

func (R *OutputCoil) Readmessage(data []byte)(interface {}, error){
	return nil,nil
}
