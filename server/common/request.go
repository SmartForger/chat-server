package common

import (
	"bytes"
	"encoding/json"
	"fullstackdevs14/chat-server/lib"

	"github.com/gin-gonic/gin"
)

func GetRequestBody[T interface{}](c *gin.Context) (T, bool) {
	var data T

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	str := buf.String()

	if str == "" {
		return data, false
	}

	secret := c.GetString(lib.HEADER_NONCE_TOKEN)
	decrypted := lib.DecryptAES(str, secret)

	err := json.Unmarshal([]byte(decrypted), &data)

	if err != nil {
		return data, false
	}

	return data, true
}

func GetSocketMessage[T interface{}](msg string) (*T, bool) {
	var data T
	var payload SocketMessage

	err := json.Unmarshal([]byte(msg), &payload)
	if err != nil {
		return nil, false
	}

	priv := lib.CGet(lib.CK_PRIVATE)
	secret, err1 := lib.DecryptRSA(payload.S, priv)
	if err1 != nil {
		return nil, false
	}

	decrypted := lib.DecryptAES(payload.T, secret)
	if decrypted == "" {
		return nil, false
	}

	err2 := json.Unmarshal([]byte(decrypted), &data)
	if err2 != nil {
		return nil, false
	}

	return &data, true
}

func GetEncryptedData(data []byte, secret string) (string, bool) {
	encrypted := lib.EncryptAESByte(data, secret)
	if encrypted == "" {
		return "", false
	}

	json, err2 := json.Marshal(SocketMessage{
		T: encrypted,
	})
	if err2 != nil {
		return "", false
	}

	return string(json), true
}
