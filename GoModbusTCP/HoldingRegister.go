package GoModbusTCP

import (
	"errors"
	"fmt"
	"encoding/binary"
	"bytes"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Holding Register
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type HRMessage struct{
	Seq uint16
}

//Write a single value
//TODO support multiple writes
func(R *HRMessage)ConstructWriteMessage(unitAddr byte, regAddr uint16 , value uint16)(b []byte, err error){
	addrSlice := make([]byte,2)
	valSlice := make([]byte, 2)
	binary.BigEndian.PutUint16(addrSlice, regAddr)
	binary.BigEndian.PutUint16(valSlice, value)
	WriteMsg := []byte{0x00,0x00,0x00,0x00,0x00,0x06, unitAddr, 0x06,addrSlice[0], addrSlice[1],valSlice[0], valSlice[1]}
	return WriteMsg, nil
}

func (R *HRMessage)ConstructReadMessage(unitAddr byte, regAddr uint16, length uint16) (b []byte, err error){
	addrSlice := make([]byte,2)
	lenSlice := make([]byte, 2)
	binary.BigEndian.PutUint16(addrSlice, regAddr)
	binary.BigEndian.PutUint16(lenSlice, length)
	ReadMsg := []byte{0x00,0x01,0x00,0x00,0x00,0x06, unitAddr, 0x03 ,addrSlice[0], addrSlice[1],lenSlice[0], lenSlice[1]}
	return ReadMsg, nil
}

func (R *HRMessage) Readmessage(data []byte)(interface {}, error){
    fmt.Println("Processing response message... ")
	//if the length of the byte array is less than 6, then the message is invalid
	if len(data) < 6{
		return nil, errors.New("Invalid response message size")
	}
	//Could panic
	length := binary.BigEndian.Uint16([]byte{data[4],data[5]})
	fmt.Println("Message size: " , length)
	fmt.Println("Slave ID: ", data[6])
	if data[7] == 0x03{
		fmt.Println("Read response confirmed, proceeding with the read operation... ")
	} else {
		return nil, errors.New("Invalid function type in reponse message ")
	}
	nValues := int(data[8]/2)
	fmt.Println("Number of values available: ", nValues)
	values := make([]uint16, nValues)
    valuesbuf := bytes.NewReader(data[9:])
	err := binary.Read(valuesbuf, binary.BigEndian, &values)
	return values,err
}
