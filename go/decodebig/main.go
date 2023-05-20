package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

const genesisBlock = "0x0a20e381741db5e128d572c459b41151dff44c713b6fa72d6107b69e630fe8ddebf912208c612a99487e277612d09f39e701b7a7b63014a0ef95ff3711922b36bb904ff21a4730450221008a482dc09db960269da3bfe9a92fb33f582b87d3b7af74548c6d75332c32575e022074be2f72776a85fa653cc29f8171228895a5a243dba0870554c6097c9a184f9b20cfd5b993062a5d57686f6576657220776f756c64206f7665727468726f7720746865206c696265727479206f662061206e6174696f6e206d75737420626567696e206279207375626475696e672074686520667265656e657373206f66207370656563683201003a86020a20170e50286de73bd7ff0574e638311e67913e91f76b869e107a5fe202aa74526712473045022100a474a389d079b9503464707626eb9096008a653e0ff77dbb9de465f51e1d180302200473fb01cad94ab3b6c73b54a38c38aa8c078ff7d377cde8a441fa8b800df9541a2103fab2023a5b2acb8855085004dc173f67d66df5591afdc3fbc3435880b9c6338b220100322a3078646439613337346538646365396436353630373365633135333538303330316237643263333835303a2a3078646439613337346538646365396436353630373365633135333538303330316237643263333835304213307832326231633863313232376130303030304a03307830520101"

const badNibble = ^uint64(0)

const uintBits = 32 << (uint64(^uint(0)) >> 63)

// Errors
var (
	ErrEmptyString   = &decError{"empty hex string"}
	ErrSyntax        = &decError{"invalid hex string"}
	ErrMissingPrefix = &decError{"hex string without 0x prefix"}
	ErrOddLength     = &decError{"hex string of odd length"}
	ErrEmptyNumber   = &decError{"hex string \"0x\""}
	ErrLeadingZero   = &decError{"hex number with leading zero digits"}
	ErrUint64Range   = &decError{"hex number > 64 bits"}
	ErrUintRange     = &decError{fmt.Sprintf("hex number > %d bits", uintBits)}
	ErrBig256Range   = &decError{"hex number > 256 bits"}
)

type decError struct{ msg string }

func (err decError) Error() string { return err.msg }

var bigWordNibbles int

func init() {
	// This is a weird way to compute the number of nibbles required for big.Word.
	// The usual way would be to use constant arithmetic but go vet can't handle that.
	b, _ := new(big.Int).SetString("FFFFFFFFFF", 16)
	switch len(b.Bits()) {
	case 1:
		bigWordNibbles = 16
	case 2:
		bigWordNibbles = 8
	default:
		panic("weird big.Word size")
	}
}

// EncodeNoPrefix encodes b as a hex string without 0x prefix.
func EncodeNoPrefix(b []byte) string {
	return hex.EncodeToString(b)
}

// Encode encodes b as a hex string with 0x prefix.
func Encode(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

// Decode decodes a hex string with 0x prefix.
func Decode(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("input is empty")
	}
	if !Has0xPrefix(input) {
		return nil, errors.New("hex prefix is missing")
	}
	b, err := hex.DecodeString(input[2:])
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex string: %w", err)
	}

	return b, err
}

// DecodeNoPrefix decodes a hex string without 0x prefix.
func DecodeNoPrefix(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("input is empty")
	}
	b, err := hex.DecodeString(input)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex string: %w", err)
	}

	return b, err
}

// Has0xPrefix
func Has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

// DecodeUint64 decodes a hex string with 0x prefix as a quantity.
func DecodeUint64(input string) (uint64, error) {
	raw, err := checkNumber(input)
	if err != nil {
		return 0, err
	}
	dec, err := strconv.ParseUint(raw, 16, 64)
	if err != nil {
		return 0, err
	}
	return dec, nil
}

// EncodeUint64 encodes i as a hex string with 0x prefix.
func EncodeUint64(i uint64) string {
	enc := make([]byte, 2, 10)
	copy(enc, "0x")
	return string(strconv.AppendUint(enc, i, 16))
}

// EncodeInt64 encodes i as a hex string with 0x prefix.
func EncodeInt64(i int64) string {
	enc := make([]byte, 2, 10)
	copy(enc, "0x")
	return string(strconv.AppendInt(enc, i, 16))
}

// EncodeBig encodes bigint as a hex string with 0x prefix.
// The sign of the integer is ignored.
func EncodeBig(bigint *big.Int) string {
	nbits := bigint.BitLen()
	if nbits == 0 {
		return "0x0"
	}
	return fmt.Sprintf("%#x", bigint)
}

// DecodeBigFromBytesToUint64 decodes a byte array to uint64.
func DecodeBigFromBytesToUint64(data []byte) uint64 {
	zeroBig := big.NewInt(0)
	return zeroBig.SetBytes(data).Uint64()
}

// EncodeUint64ToBytes encodes a uint64 number to bytes.
func EncodeUint64ToBytes(number uint64) []byte {
	if number == 0 {
		return []byte{0}
	}

	return big.NewInt(0).SetUint64(number).Bytes()
}

// EncodeUint64ToBytes encodes a uint64 number to bytes.
func EncodeUint64BytesToHexString(number []byte) string {
	return EncodeBig(big.NewInt(0).SetBytes(number))
}

// DecodeBig decodes a hex string with 0x prefix as a quantity.
// Numbers larger than 256 bits are not accepted.
func DecodeBig(input string) (*big.Int, error) {
	raw, err := checkNumber(input)
	if err != nil {
		return nil, err
	}
	if len(raw) > 64 {
		return nil, ErrBig256Range
	}
	words := make([]big.Word, len(raw)/bigWordNibbles+1)
	end := len(raw)
	for i := range words {
		start := end - bigWordNibbles
		if start < 0 {
			start = 0
		}
		for ri := start; ri < end; ri++ {
			nib := decodeNibble(raw[ri])
			if nib == badNibble {
				return nil, ErrSyntax
			}
			words[i] *= 16
			words[i] += big.Word(nib)
		}
		end = start
	}
	dec := new(big.Int).SetBits(words)
	return dec, nil
}

// ExtractHex hexadecimal value out of a string.
func ExtractHex(s string) string {
	r := regexp.MustCompile(`0x([A-Fa-f0-9]{6,})`)
	matches := r.FindStringSubmatch(s)
	if len(matches) > 1 {
		return "0x" + matches[1]
	}

	return ""
}

func checkNumber(input string) (raw string, err error) {
	if len(input) == 0 {
		return "", ErrEmptyString
	}
	if !Has0xPrefix(input) {
		return "", ErrMissingPrefix
	}
	input = input[2:]
	if len(input) == 0 {
		return "", ErrEmptyNumber
	}
	if len(input) > 1 && input[0] == '0' {
		return "", ErrLeadingZero
	}
	return input, nil
}

func decodeNibble(in byte) uint64 {
	switch {
	case in >= '0' && in <= '9':
		return uint64(in - '0')
	case in >= 'A' && in <= 'F':
		return uint64(in - 'A' + 10)
	case in >= 'a' && in <= 'f':
		return uint64(in - 'a' + 10)
	default:
		return badNibble
	}
}

func main() {
	b, _ := Decode(genesisBlock)
	fmt.Println(string(b))
}