package main

import (
	"fmt"
	"strconv"
)

// ParseUint64 parses s as an integer in decimal or hexadecimal syntax.
// Leading zeros are accepted. The empty string parses as zero.
func ParseUint64(s string) (uint64, bool) {
	if s == "" {
		return 0, true
	}
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		v, err := strconv.ParseUint(s[2:], 16, 64)
		return v, err == nil
	}
	v, err := strconv.ParseUint(s, 10, 64)
	return v, err == nil
}

const genesisBlock = "0x0a20e381741db5e128d572c459b41151dff44c713b6fa72d6107b69e630fe8ddebf912208c612a99487e277612d09f39e701b7a7b63014a0ef95ff3711922b36bb904ff21a4730450221008a482dc09db960269da3bfe9a92fb33f582b87d3b7af74548c6d75332c32575e022074be2f72776a85fa653cc29f8171228895a5a243dba0870554c6097c9a184f9b20cfd5b993062a5d57686f6576657220776f756c64206f7665727468726f7720746865206c696265727479206f662061206e6174696f6e206d75737420626567696e206279207375626475696e672074686520667265656e657373206f66207370656563683201003a86020a20170e50286de73bd7ff0574e638311e67913e91f76b869e107a5fe202aa74526712473045022100a474a389d079b9503464707626eb9096008a653e0ff77dbb9de465f51e1d180302200473fb01cad94ab3b6c73b54a38c38aa8c078ff7d377cde8a441fa8b800df9541a2103fab2023a5b2acb8855085004dc173f67d66df5591afdc3fbc3435880b9c6338b220100322a3078646439613337346538646365396436353630373365633135333538303330316237643263333835303a2a3078646439613337346538646365396436353630373365633135333538303330316237643263333835304213307832326231633863313232376130303030304a03307830520101"

func main(){
	s := fmt.Sprintf("%#x", uint64(100000))
	fmt.Println(s)
	c, _ := ParseUint64(s)
	fmt.Println(c)
}