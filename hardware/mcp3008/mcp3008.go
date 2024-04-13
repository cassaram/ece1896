package mcp3008

import (
	"periph.io/x/conn/v3/spi"
)

type MCP3008 struct {
	conn spi.Conn
}

func NewMCP3008(conn *spi.Conn) *MCP3008 {
	d := MCP3008{
		conn: *conn,
	}
	return &d
}

func (d *MCP3008) ReadPort(port uint8) (float64, error) {
	cmdPacket := 0x80 | ((port & 0x7) << 4)
	txPackets := []byte{0x01, cmdPacket, 0x00}
	rxPackets := make([]byte, len(txPackets))
	err := d.conn.Tx(txPackets, rxPackets)
	if err != nil {
		return 0, err
	}
	val16 := uint16(rxPackets[1]&0x07)<<8 | uint16(rxPackets[2])
	val := float64(val16) / 1024
	return val, nil
}
