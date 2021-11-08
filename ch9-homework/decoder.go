package main

import (
	"encoding/binary"
	"fmt"
)

type Protocol struct {
	PackageLen  uint32
	HeaderLen   uint16
	ProtocolVer uint16
	Operation   Operation
	SequenceID  uint32
	Body        string
}

type Operation int

const (
	OpHandshake Operation = iota
	OpHandshakeReply
	OpHeartbeat
	OpHeartbeatReply
	OpSendMsg
	OpSendMsgReply
	OpDisconnectReply
	OpAuth
	OpAuthReply
	OpRawBatch
)

func (d Operation) String() string {
	return [...]string{"handshake", "handshake reply", "heartbeat", "heartbeat reply",
		"send message", "send message reply", "connection disconnect reply", "auth connnect", "auth connect reply",
		"batch message for websocket"}[d]
}

// Decoder decodes GOIM protocol from a socket connection, per https://goim.io/docs/protocol.html
func Decoder(in []byte) (*Protocol, error) {
	if len(in) < 16 {
		return nil, fmt.Errorf("invalid data")
	}

	res := Protocol{}
	res.PackageLen = binary.BigEndian.Uint32(in[:4])
	res.HeaderLen = binary.BigEndian.Uint16(in[4:6])
	res.ProtocolVer = binary.BigEndian.Uint16(in[6:8])
	res.Operation = Operation(binary.BigEndian.Uint32(in[8:12]))
	res.SequenceID = binary.BigEndian.Uint32(in[12:16])

	if len(in) < (16 + int(res.HeaderLen) + int(res.PackageLen)) {
		return nil, fmt.Errorf("invalid data")
	}
	res.Body = string(in[16 : 16+int(res.HeaderLen)+int(res.PackageLen)])

	return &res, nil
}

func main() {
	res, _ := Decoder(append([]byte{0, 0, 0, 4, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 2}, []byte("body")...))
	fmt.Println(res)
}
