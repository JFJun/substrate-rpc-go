package utils

import (
	"encoding/hex"
	"errors"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/substrate-rpc-go/xxhash"
	"golang.org/x/crypto/blake2b"
	"hash"
)

func SelectHash(method string) (hash.Hash, error) {
	switch method {
	case "Twox128":
		return xxhash.New128(nil), nil
	case "Blake2_256":
		return blake2b.New256(nil)
	case "Blake2_128":
		return blake2b.New(16, nil)
	case "Blake2_128Concat":
		return blake2b.New(16, nil)
	case "Twox64Concat":
		return xxhash.New64(nil), nil
	case "Identity":
		return nil, nil
	default:
		return nil, errors.New("unknown hash method")

	}
}

func AddressToPublicKey(address string) string {
	if address == "" {
		return ""
	}
	pub, err := ss58.DecodeToPub(address)

	if err != nil {
		return ""
	}
	if len(pub) != 32 {
		return ""
	}
	pubHex := hex.EncodeToString(pub)
	return pubHex
}