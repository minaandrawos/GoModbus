////////////////////////
//factory for classes//
//////////////////////

package GoModbusTCP

import (
	"os"
	"errors"
	"fmt"
	"net"
	"strings"
)

const (
	DISCRETE_OUT_COIL = 0
	DISCRETE_IN_COIL = 1
	ANALOG_IN_REGISTER = 3
	ANALOG_HOLDING_REGISTER = 4
)

type RegisterFactory interface {
	ConstructWriteMessage(byte,uint16, uint16) ([]byte,error)
	ConstructReadMessage(byte, uint16, uint16)([]byte,error)
	Readmessage(data []byte)(interface {}, error)
}

func CreateRegister(Type uint) (RegisterFactory, error){
	switch Type{
	case DISCRETE_OUT_COIL:
		return new(OutputCoil),nil
	case DISCRETE_IN_COIL:
		return new(InputCoil),nil
	case ANALOG_IN_REGISTER:
		return new(IRMessage),nil
	case ANALOG_HOLDING_REGISTER:
		return new(HRMessage),nil
	}
	return nil,errors.New("ERROR: Invalid modbus register type..")
}

//Everything must be upper case
type ModbusRequest struct{
	RegType uint
	Raddr uint
	Uaddr uint
	Operation string
	Value uint // value for write operations or length for read operations
	Dest string
}

func (mr *	ModbusRequest)Handlerequest()(result interface{}, err error){
	Register,err := CreateRegister(mr.RegType)
	checkError(err)
	mr.Operation = strings.ToUpper(mr.Operation)
	if mr.Operation == "W" {
		result,err = performWrites(Register, mr)
	} else if mr.Operation == "R"{
		result, err = performReads(Register, mr)
	} else {
		//fmt.Println("Operation not defined!! It needs to be either R for read operations or W for write operations")
		return nil, errors.New("Operation not defined!! It needs to be either R for read operations or W for write operations..")
	}
	return result, err
}

func performReads(register RegisterFactory, mr *ModbusRequest )(interface{}, error){
	readMsg,err := register.ConstructReadMessage(byte(mr.Uaddr), uint16(mr.Raddr), uint16(mr.Value))
	checkError(err)
	rcvdData := handleMessage(mr.Dest, readMsg)
	return  register.Readmessage(rcvdData)
}

func performWrites(register RegisterFactory, mr *ModbusRequest )(interface{}, error){
	writeMsg,err := register.ConstructWriteMessage(byte(mr.Uaddr), uint16(mr.Raddr), uint16(mr.Value))
	if(err!=nil){
		return nil, err
	}
    rcvdData := handleMessage(mr.Dest, writeMsg)

	//If the received message is an echo of the sent write message then the write was successful
    return  (rcvdData[0]!=0x86 && rcvdData[0]!=0x90), nil
}

func handleMessage(dst string , data []byte) (response []byte){
	//sanity check
	if data == nil{
		return
	}
	conn, err := net.Dial("tcp", dst)
	checkError(err)
	defer conn.Close()
	fmt.Println("Send data Message \n", data)
	_, err = conn.Write(data)
	checkError(err)
	fmt.Println("Received data Message")
	response = make([]byte,512)
	n,err := conn.Read(response)
	fmt.Println(response[1:n])
	checkError(err)
	return
}

func checkError(err error){
	if(err!=nil){
		fmt.Println("Error occured", err)
		os.Exit(1)
	}
}
