package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	mrand "math/rand"
	"strings"
	"time"
	"unsafe"

	"github.com/google/uuid"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var src = mrand.NewSource(time.Now().UnixNano())

func rand32ByteStringBytesMaskImprSrcUnsafe() string {
	const n = 32
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// GetAESDecrypted decrypts given text in AES 256 CBC
func GetAESDecrypted(encrypted, key, iv string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)

	return string(ciphertext), nil
}

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

// GetAESEncrypted encrypts given text in AES 256 CBC
func GetAESEncrypted(plaintext, key, iv string) (string, error) {
	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil
}

func GenerateKey(patientId uuid.UUID, iv string) (string, error) {
	randomString := rand32ByteStringBytesMaskImprSrcUnsafe()
	shortedId := strings.ReplaceAll(patientId.String(), "-", "")
	encrypted, err := GetAESEncrypted(randomString, shortedId, iv)
	if err != nil {
		return "", err
	}

	return encrypted[:32], nil
}
