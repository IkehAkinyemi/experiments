package udp

const (
	DatagramSize = 516 // the maximum supported datagram size
	BlockSize = DatagramSize - 4 // the DatagramSize mins a 4-byte header
)

type OpCode uint16

const (
	OpRRQ OpCode = iota + 1
	_ // no WRQ support
	OpData
	OpAck
	OpErr
)

type ErrCode uint16

const (
	ErrUnknown ErrCode = iota
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrInknownID
	ErrFileExists
	ErrNoUser
)

// curl localhost:2019/config/apps/http/servers/test \ -X POST -H "Content-Type: application/json" \
// -d '{
//       "listen": ["localhost:2030"],
//       "routes": [{
//         "handle": [{
//           "handler": "static_response",
//           "body": "Welcome to my temporary test server."
// }] }]
// }'