package GoModbusTCP
import (
	"fmt"
	"encoding/binary"
	"errors"
)

//Input Coil
type InputCoil struct{
	Seq uint16
}

func(R *InputCoil)ConstructWriteMessage(unitAddr byte, regAddr uint16, value uint16)(b []byte, err error) {
	return  nil, errors.New("Input coils do not allow writes..")
}

func (R *InputCoil)ConstructReadMessage(unitAddr byte, regaddr uint16, length uint16) (b []byte, err error){
	addrSlice := make([]byte,2)
	lenSlice := make([]byte, 2)
	binary.BigEndian.PutUint16(addrSlice, regaddr)
	binary.BigEndian.PutUint16(lenSlice, length)
	ReadMsg := []byte{0x00,0x00,0x00,0x00,0x00,0x06, unitAddr, 0x02 ,addrSlice[0], addrSlice[1],lenSlice[0], lenSlice[1]}
	fmt.Println(ReadMsg)
	return ReadMsg, nil
}

func (R *InputCoil) Readmessage(data []byte)(interface {}, error){
	return nil,nil
}
