package ub64c

import (
	"encoding/base64"

	"github.com/gofrs/uuid"
)

func b64EncodingWithString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func b64EncodingWithByte(by []byte) string {
	return base64.StdEncoding.EncodeToString(by)
}

func b64DecodingWithString(str string) (string, error) {
	by, err := base64.StdEncoding.DecodeString(str)
	return uuid.FromBytesOrNil(by).String(), err
}
