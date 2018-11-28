package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Encrypter struct {

}

func (e *Encrypter) Encrypt(key []byte, data []byte) []byte{
	// AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encryptedData := make([]byte, aes.BlockSize+len(data))
	iv := encryptedData[:aes.BlockSize]
	// Preenche iv aleatoriamente
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	// Criptografa data em encryptedData
	stream.XORKeyStream(encryptedData[aes.BlockSize:], data)
	return encryptedData
}

func (e *Encrypter) Decrypt(key []byte, encData []byte) []byte{
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(encData) < aes.BlockSize {
		panic("Text is too short")
	}
	iv := encData[:aes.BlockSize]
	encData = encData[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encData, encData)
	return encData
}