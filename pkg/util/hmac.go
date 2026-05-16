package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ComputeHMAC(timestamp, method, path, body, secret string) string {
	payload := fmt.Sprintf("%s%s%s%s", timestamp, method, path, body)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}

func VerifyHMAC(timestamp, method, path, body, secret, receivedSignature string) bool {
	expected := ComputeHMAC(timestamp, method, path, body, secret)
	expectedBytes, err := hex.DecodeString(expected)
	if err != nil {
		return false
	}
	receivedBytes, err := hex.DecodeString(receivedSignature)
	if err != nil {
		return false
	}
	return hmac.Equal(expectedBytes, receivedBytes)
}
