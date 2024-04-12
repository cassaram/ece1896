package mcp23s17

type mcp23s17Register byte

const (
	IODIRA   mcp23s17Register = 0x00
	IODIRB   mcp23s17Register = 0x01
	IPOLA    mcp23s17Register = 0x02
	IPOLB    mcp23s17Register = 0x03
	GPINTENA mcp23s17Register = 0x04
	GPINTENB mcp23s17Register = 0x05
	DEFVALA  mcp23s17Register = 0x06
	DEFVALB  mcp23s17Register = 0x07
	INTCONA  mcp23s17Register = 0x08
	INTCONB  mcp23s17Register = 0x09
	IOCON    mcp23s17Register = 0x0A // Second address to IOCON on 0x0B
	GPPUA    mcp23s17Register = 0x0C
	GPPUB    mcp23s17Register = 0x0D
	INTFA    mcp23s17Register = 0x0E
	INTFB    mcp23s17Register = 0x0F
	INTCAPA  mcp23s17Register = 0x10
	INTCAPB  mcp23s17Register = 0x11
	GPIOA    mcp23s17Register = 0x12
	GPIOB    mcp23s17Register = 0x13
	OLATA    mcp23s17Register = 0x14
	OLATB    mcp23s17Register = 0x15
)
