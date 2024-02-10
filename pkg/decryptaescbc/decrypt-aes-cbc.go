package decryptaescbc

import (
	"errors"
	"project/pkg/lzstring"
	"project/pkg/pkcs7"

	"crypto/aes"
	"crypto/cipher"

	"crypto/sha256"
	"encoding/base64"
)

func AESCBCDecrypt(encrypted string, key string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(key))

	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("cipherText too short")
	}

	iv := hash[:aes.BlockSize]

	if len(cipherText)%aes.BlockSize != 0 {
		return "", errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	// cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	// data, err := lzstring.DecompressFromEncodedUriComponent(string(cipherText))
	data, err := lzstring.DecompressFromEncodedUriComponent(string(cipherText))
	if err != nil {
		return "", err
	}

	return data, nil
}
