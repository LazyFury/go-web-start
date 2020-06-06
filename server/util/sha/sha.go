package sha

import (
	"crypto/aes"
	"crypto/cipher"
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
	return fmt.Sprintf("%x", ciphertext)
}

// DeCode DeCode
func DeCode(str string) string {
	c, _ := aes.NewCipher([]byte(key))
	strNew := []byte(str)
	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, iv)
	plaintextCopy := make([]byte, len(strNew))
	cfbdec.XORKeyStream(plaintextCopy, strNew)
	// fmt.Printf("%x=>%s\n", strNew, plaintextCopy)
	return fmt.Sprintf("%s", plaintextCopy)
}
