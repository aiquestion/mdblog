package models

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"os"
)

func ValidateSignature(sig, payload string) bool {
	key := []byte(os.Getenv("GITHUB_SIG"))
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(payload))
	payloadSignature := "sha1=" + hex.EncodeToString(mac.Sum(nil))
	return subtle.ConstantTimeCompare([]byte(sig), []byte(payloadSignature)) == 1
}
