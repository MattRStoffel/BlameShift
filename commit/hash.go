package commit

import (
	"crypto/sha256"
)

func hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}
