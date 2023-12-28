package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"sync"
)

const (
	// nonceSize is the size of the nonce used for the AES-GCM encryption.
	// the value is taken from reverse-engineering the `marplogin` command.
	nonceSize = 16
)

var (
	// key is the AES key used for the AES-GCM encryption. The value is taken
	// from reverse-engineering the `marplogin` command. Based on the
	// key value, it's not really clear why encryption is used at all.
	key = []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	// as the key is static, we can use a single AES decrypter for all
	// tickets
	aesSingle *AES
	lock      = &sync.Mutex{}
)

type AES struct {
	block *cipher.Block
	aead  *cipher.AEAD
}

// getAES returns a AES decrypter using the key and nonce size taken from
// reverse-engineering the `marplogin` command.
func getAES() (*AES, error) {
	if aesSingle != nil {
		return aesSingle, nil
	}

	lock.Lock()
	defer lock.Unlock()

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		return nil, err
	}

	return &AES{
		block: &block,
		aead:  &aesGCM,
	}, nil
}

// Decrypt decrypts the given cipher text using the AES-GCM algorithm.
func Decrypt(cipherText []byte) ([]byte, error) {
	aes, err := getAES()
	if err != nil {
		return nil, err
	}

	nonce := cipherText[:nonceSize]
	cipherText = cipherText[nonceSize:]

	plainText, err := (*aes.aead).Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// Encrypt encrypts the given plain text using the AES-GCM algorithm.
func Encrypt(plainText []byte) ([]byte, error) {
	aes, err := getAES()
	if err != nil {
		return nil, err
	}

	// generate a random nonce
	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	cipherText := (*aes.aead).Seal(nil, nonce, plainText, nil)

	return append(nonce, cipherText...), nil
}
