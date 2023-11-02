package lib

import (
	"os"
)

func GetAdminSecret() string {
	return os.Getenv(ENV_ADMIN_SECRET) + "Zbri3TtBk="
}

func IsAdminToken(token, nonce string) bool {
	priv := CGet(CK_PRIVATE)

	decryptedPayload := DecryptAES(token, nonce)
	decryptedToken, err := DecryptRSA(decryptedPayload, priv)
	if err != nil {
		return false
	}

	return decryptedToken == GetAdminSecret()
}
