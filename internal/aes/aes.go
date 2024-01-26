package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

const (
	// nonceSize is the size of the nonce used for the AES-GCM encryption.
	// the value is taken from reverse-engineering the `marplogin` command.
	nonceSize = 16
)

var (
	// defaultKey is the AES defaultKey used for the AES-GCM encryption. The value is taken
	// from reverse-engineering the `marplogin` command. Based on the
	// defaultKey value, it's not really clear why encryption is used at all.
	defaultKey = []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
)

type AES struct {
	block cipher.Block
	aead  cipher.AEAD
	rand  randReader
}

type aesOptions struct {
	key  []byte
	rand randReader
}

type AESOption func(*aesOptions)

// WithRand allows overriding the rand.Reader used by the AES implementation.
func WithRand(rand randReader) AESOption {
	return func(aes *aesOptions) {
		aes.rand = rand
	}
}

// New returns a AES decrypter using the key and nonce size taken from
// reverse-engineering the `marplogin` command.
func New(opts ...AESOption) (*AES, error) {
	// set the default options
	options := &aesOptions{
		key:  defaultKey,
		rand: &cryptoRandReader{},
	}

	// apply the given options
	for _, opt := range opts {
		opt(options)
	}

	// create the AES-GCM cipher
	block, err := aes.NewCipher(options.key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCMWithNonceSize(block, nonceSize)
	if err != nil {
		return nil, err
	}

	return &AES{
		block: block,
		aead:  aesGCM,
		rand:  options.rand,
	}, nil
}

// Decrypt decrypts the given cipher text using the AES-GCM algorithm.
func (aes *AES) Decrypt(cipherText []byte) ([]byte, error) {
	nonce := cipherText[:nonceSize]
	cipherText = cipherText[nonceSize:]

	plainText, err := aes.aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// Encrypt encrypts the given plain text using the AES-GCM algorithm.
func (aes *AES) Encrypt(plainText []byte) ([]byte, error) {
	// generate a random nonce
	nonce := make([]byte, nonceSize)
	if n, err := aes.rand.Read(nonce); err != nil {
		return nil, err
	} else if n != nonceSize {
		return nil, fmt.Errorf("expected to read %d bytes, got %d", nonceSize, n)
	}

	cipherText := aes.aead.Seal(nil, nonce, plainText, nil)

	return append(nonce, cipherText...), nil
}
