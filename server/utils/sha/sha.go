package sha

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var (
	iv  = []byte("we3tg4j4ekjhabns")
	key = []byte("wertghjdlkjhabnswertghjdlkjhabns")
)

// EnCode Encode
func EnCode(str string) string {
	c, _ := aes.NewCipher([]byte(key))
	strNew := []byte(str)

	cfb := cipher.NewCFBEncrypter(c, iv)
	ciphertext := make([]byte, len(strNew))
	cfb.XORKeyStream(ciphertext, strNew)
	// fmt.Printf("%s=>%x\n", strNew, ciphertext)
	return hex.EncodeToString(ciphertext)
}

// AesDecryptCFB DeCode
func AesDecryptCFB(str string) (decrypted string) {
	block, _ := aes.NewCipher([]byte(key))
	encrypted, _ := hex.DecodeString(str)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return fmt.Sprintf("%s", encrypted)
}
