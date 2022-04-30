package main

import (
	"crypto"
	"encoding/base64"
)

func CryptoHashString(data string) string {
	shaHashFunc := crypto.SHA384.New()
	shaHashFunc.Write([]byte(data))
	return base64.URLEncoding.EncodeToString(shaHashFunc.Sum(nil))
}
