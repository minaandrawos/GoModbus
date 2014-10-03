GoModbus
========

A TCP Modbus client in Go

This is a simple TCP Modbus client written in Go for testing and protyping purposes. The code is in an alpha state and is provided as-is. Please keep my name in the files as the author if reused.

What is Modbus?
Modbus is a popular industrial device protocol used to communicate with numerous devices worldwide. A good primer on the protocol can be found at: http://www.simplymodbus.ca/FAQ.htm#Modbus 

Code structure:
 1. ModbusClient folder is where the main package is, the entry point for the project
 2. GoModbusTCP folder is where the Modbus TCP implementation resides. There is a file for each Modbus device type and      a file that acts as a factory of device structs based on the user requests. This could evolve to become an API in      the future

 Notes:
  1. CRC not supported
  2. Code can be tested with Modbus simulators
  3. I may write a tutorial in the future covering the design details when I get a chance.
  4. Intrested for more material? check my website at www.minaandrawos.com
