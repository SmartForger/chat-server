package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func marshalRSAPrivate(priv *rsa.PrivateKey) string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}))
}

func GenerateKey() (string, string, error) {
	reader := rand.Reader
	bitSize := 2048

	privateKey, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		return "", "", err
	}
	publicKey := &privateKey.PublicKey

	pubKeyStr := string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}))
	privKeyStr := string(pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}))

	return pubKeyStr, privKeyStr, nil
}

func EncryptRSA(msg, publicKey string) (string, error) {
	block, _ := pem.Decode([]byte(publicKey))
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		key,
		[]byte(msg),
		nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func DecryptRSA(data, priv string) (string, error) {
	data2, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode([]byte(priv))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, data2, nil)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func GenerateAESKey() string {
	key := make([]byte, 32)

	_, err := rand.Read(key)

	if err != nil {
		return ""
	}

	return base64Encode(key)
}

func base64Encode(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}

func base64Decode(key string) []byte {
	secretKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}

	return secretKey
}

func EncryptAESByte(bytes []byte, secretKey string) string {
	aes, err := aes.NewCipher(base64Decode(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, bytes, nil)

	return base64Encode(ciphertext)
}

func EncryptAES(plaintext string, secretKey string) string {
	return EncryptAESByte([]byte(plaintext), secretKey)
}

func DecryptAES(ciphertext string, secretKey string) string {
	aes, err := aes.NewCipher(base64Decode(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	ciphertextb := base64Decode(ciphertext)
	nonceSize := gcm.NonceSize()
	nonce, ciphertextb := ciphertextb[:nonceSize], ciphertextb[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertextb, nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
