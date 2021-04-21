package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func GenerateKey() string {
	// Generate a random 32 byte key for AES-256.
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println(err.Error())
		return ""
	}
	// Encode key in bytes to string and keep as secret, put in a vault.
	key := hex.EncodeToString(bytes)
	return key
}

func Encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	// Since the key is in string, we need to convert decode it to bytes.
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	// Create a new Cipher Block from the key.
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return ""	}

	// Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode.
	// https://golang.org/pkg/crypto/cipher/#NewGCM.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
		return ""	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err.Error())
		return ""	}

	// Encrypt the data using aesGCM.Seal.
	// Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func Decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	// Create a new Cipher Block from the key.
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Create a new GCM.
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	// Get the nonce size.
	nonceSize := aesGCM.NonceSize()

	// Extract the nonce from the encrypted data.
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// Decrypt the data.
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return fmt.Sprintf("%s", plaintext)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
