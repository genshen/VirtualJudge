package auth

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func SHA1(data string) string {
	h := sha1.New()
	io.WriteString(h, data)
	return hex.EncodeToString(h.Sum(nil))
}
