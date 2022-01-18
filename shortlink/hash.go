package shortlink

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"example.com/shortlink-kata/db"
)

const defaultHashLength = 7

var isAlreadyCreated = db.Exists

func generateHash(s string) (string, error) {
	hash := sha1.Sum([]byte(s)) //outputs into hex format, but should be sufficient for our purposes

	var potentialHash string
	maxAdditionalChars := len(hash) - defaultHashLength
	for i := 0; i <= maxAdditionalChars; i++ {
		potentialHash = hex.EncodeToString(hash[:defaultHashLength+i])

		if !isAlreadyCreated(potentialHash) {
			break
		}
		if i == maxAdditionalChars {
			return "", fmt.Errorf("could not generate an unused hash for address %v", s)
		}
	}

	return potentialHash, nil
}
