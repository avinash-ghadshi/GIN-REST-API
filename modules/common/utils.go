package common

import (
	"crypto/sha1"
	"encoding/base64"
)

func GetSHA1(bv []byte) string {
	hasher := sha1.New()
	hasher.Write(bv)
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
