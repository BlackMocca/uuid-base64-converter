package ub64c

import (
	"github.com/gofrs/uuid"
)

func isUUID(input string) (bool, error) {
	_, err := uuid.FromString(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func isBase64(input string) (bool, error) {
	_, err := b64DecodingWithString(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Convert(input string) (uid string, b64 string, err error) {
	_, uidErr := isUUID(input)
	_, b64Err := isBase64(input)

	if uidErr != nil && b64Err == nil {
		/* case input is base64 */
		decoding, decodeErr := b64DecodingWithString(input)

		b64 = input
		uid = decoding
		err = decodeErr
	} else if uidErr == nil && b64Err != nil {
		/* case input is uuid */
		b64 = b64EncodingWithByte(uuid.FromStringOrNil(input).Bytes())
		uid = input
	}

	return
}
