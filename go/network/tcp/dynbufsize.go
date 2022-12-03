package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	BinaryType uint8 = iota + 1
	StringType

	MaxPayloadSize = 10 << 20 // 10MB
)

var ErrMaxPayloadSize = errors.New("maximum payload size exceeded")

type Payload interface {
	fmt.Stringer
	io.ReaderFrom
	io.WriterTo
	Bytes() []byte
}

type Binary []byte

func (b Binary) Bytes() []byte { return b }

func (b Binary) String() string { return string(b) }

func (b Binary) WriteTo(w io.Writer) (n int64, err error) {
	err = binary.Write(w, binary.BigEndian, BinaryType)
	if err != nil {
		return 0, err
	}

	n = 1

	err = binary.Write(w, binary.BigEndian, uint64(len(b)))
	if err != nil {
		return n, err
	}

	n += 4
	o, err := w.Write(b)
	return n + int64(o), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ) // 1-byte type
	if err != nil {
		return 0, err
	}
	var n int64 = 1
	if typ != BinaryType {
		return n, errors.New("invalid Binary")
	}
	var size uint32
	err = binary.Read(r, binary.BigEndian, &size) // 4-byte size
	if err != nil {
		return n, err
	}
	n += 4
	if size > MaxPayloadSize {
		return n, ErrMaxPayloadSize
	}

	*m = make([]byte, size)
	o, err := r.Read(*m) // payload
	return n + int64(o), err
}

type String string

func (m String) Bytes() []byte  { return []byte(m) }
func (m String) String() string { return string(m) }

func (m String) WriteTo(w io.Writer) (n int64, err error) {
	err = binary.Write(w, binary.BigEndian, StringType) // 1-byte type
	if err != nil {
		return 0, err
	}

	n = 1

	err = binary.Write(w, binary.BigEndian, uint32(len(m))) // 4-byte size
	if err != nil {
		return n, err
	}
	n += 4

	o, err := w.Write([]byte(m)) // payload
	return n + int64(o), err
}

func (m *String) ReadFrom(r io.Reader) (int64, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ) // 1-byte type
	if err != nil {
		return 0, err
	}

	var n int64 = 1

	if typ != StringType {
		return n, errors.New("invalid String")
	}

	var size uint32
	err = binary.Read(r, binary.BigEndian, &size) // 4-byte size
	if err != nil {
		return n, err
	}
	n += 4

	buf := make([]byte, size)
	o, err := r.Read(buf) // payload
	if err != nil {
		return n, err
	}
	*m = String(buf)

	return n + int64(o), nil
}

func decode(r io.Reader) (Payload, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Here1: %+v\n", typ)

	var payload Payload
	switch typ {
	case BinaryType:
		payload = new(Binary)	
	case StringType:
		payload = new(String)
		fmt.Printf("Here2.2: %+v\n", payload)
	default:
		return nil, errors.New("unknown type")
	}

	_, err = payload.ReadFrom(
		io.MultiReader(bytes.NewReader([]byte{typ}), r))
	if err != nil {
		return nil, err
	}

	return payload, nil
}
