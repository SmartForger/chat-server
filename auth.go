package main

import (
	"os"
)

func GetAdminSecret() string {
	return os.Getenv(ENV_ADMIN_SECRET) + "Zbri3TtBk="
}

func IsAdminToken(token string) bool {
	priv := CGet(CK_PRIVATE)

	decryptedToken, err := DecryptRSA(token, priv)
	if err != nil {
		return false
	}

	return decryptedToken == GetAdminSecret()
}
