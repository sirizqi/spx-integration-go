package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const checkSignFormat = "%d_%d_%d_%s"

func GenerateCheckSign(appId uint64, secret string, timestamp, random int64, payload []byte) (string, error) {
	original := fmt.Sprintf(checkSignFormat, appId, timestamp, random, payload)
	m := hmac.New(sha256.New, []byte(secret))
	if _, err := m.Write([]byte(original)); err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}

// Verify for webhook
func VerifyCheckSign(appId uint64, secret string, timestamp, random int64, payload []byte, provided string) bool {
	want, err := GenerateCheckSign(appId, secret, timestamp, random, payload)
	if err != nil {
		return false
	}
	return hmac.Equal([]byte(want), []byte(provided))
}
