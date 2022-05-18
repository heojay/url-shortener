package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/big"
)

func EncodeBase62(u uuid.UUID) string {
	var i big.Int
	i.SetBytes(u[:])
	return i.Text(62)
}

func DecodeBase62(encoded string) (uuid.UUID, error) {
	var i big.Int
	_, ok := i.SetString(encoded, 62)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("cannot parse base62: %q", encoded)
	}

	var u uuid.UUID
	copy(u[:], i.Bytes())
	return u, nil
}
