package main

import (
	"os"
	"fmt"
	"GoMdbus/GoModbusTCP"
	"flag"
)

func main() {
	fmt.Println("Modbus Client Started")
	regType := flag.Uint("t", 4 ,"Register type")
	raddr := flag.Uint("r", 1 ,"Starting address")
	uaddr := flag.Uint("u", 2 ,"Unit address")
	operation := flag.String("o","W", "Read or Write operation")
	value := flag.Uint("v", 3 ,"Value to Write for write requests")
	dest :=  flag.String("dst","127.0.0.1:502", "Destination address")
	flag.Parse()
	requestHandler := GoModbusTCP.ModbusRequest{
			*regType,
			*raddr,
			*uaddr,
			*operation,
			*value,
			*dest,
		}
		result, err := requestHandler.Handlerequest()
		checkError(err)
		if(result != nil) {
			fmt.Println("Operation results: ", result)
		}
}



func checkError(err error) {
	if (err != nil) {
		fmt.Println("Error occured", err)
		os.Exit(1)
	}
}
