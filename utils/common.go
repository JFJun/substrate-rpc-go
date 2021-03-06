package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JFJun/go-substrate-crypto/ss58"
	"github.com/JFJun/substrate-rpc-go/xxhash"
	"golang.org/x/crypto/blake2b"
	"hash"
	"strings"
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

func Remove0X(hexData string) string {
	if strings.HasPrefix(hexData, "0x") {
		return hexData[2:]
	}
	return hexData
}

func ZeroBytes(data []byte) {
	for i, _ := range data {
		data[i] = 0
	}
}

func RemoveHex0x(hexStr string) string {
	if strings.HasPrefix(hexStr, "0x") {
		return hexStr[2:]
	}
	return hexStr
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsNumberString(str string) bool {
	for _, a := range str {
		if a > 57 || a < 48 {
			return false
		}
	}
	return true
}

func CheckStructData(object interface{}) {
	d, _ := json.Marshal(object)
	fmt.Println(string(d))
}
