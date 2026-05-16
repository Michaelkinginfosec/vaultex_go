package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

//function to generate api key

func GenerateAPIKey() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("pk_live_%s", hex.EncodeToString(bytes)), nil
}

// functin to generate secret key
func GenerateAPISecretKey() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("sk_live_%s", hex.EncodeToString(bytes)), nil
}
