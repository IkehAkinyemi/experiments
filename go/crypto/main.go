package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func encryptMessage(key string, message string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(message))
	c.Encrypt(msgByte, []byte(message))
	return hex.EncodeToString(msgByte)
}

func decryptMessage(key string, message string) string {
	txt, _ := hex.DecodeString(message)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}
	msgByte := make([]byte, len(txt))
	c.Decrypt(msgByte, []byte(txt))

	msg := string(msgByte[:])
	return msg
}

func main() {

	plainText := "This is a secret"                  
	key := "this_must_be_of_32_byte_length!!"

	emsg := encryptMessage(key, plainText)
	dmesg := decryptMessage(key, emsg)

	fmt.Println("Encrypted Message: ", emsg)
	fmt.Println("Decrypted Message: ", dmesg)

}