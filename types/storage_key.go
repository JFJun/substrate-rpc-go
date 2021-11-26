package types

import (
	"errors"
	"fmt"
	"github.com/JFJun/substrate-rpc-go/utils"
	"github.com/JFJun/substrate-rpc-go/xxhash"
	scale "github.com/itering/scale.go"
	"strings"
)


/*
func: create storage key
author: flynn
date: 2021/11/25
*/

const (
	IsPlainType = iota
	IsMap
	IsDoubleMap
	IsNMap
)
type StorageKey []byte

// NewStorageKey creates a new StorageKey type
func NewStorageKey(b []byte) StorageKey {
	return b
}

func ( sk *StorageKey)Hex()string{
	return fmt.Sprintf("%#x", sk)
}

func CreateStorageKey(meta *scale.MetadataDecoder, module, fn string, args ...[]byte) (StorageKey, error) {
	var (
		hashMethods []string
		mapType  = -1
	)
	LOOP:
		for _,mod:=range meta.Metadata.Metadata.Modules{
			if mod.Name==module {
				for _,mt :=range mod.Storage{
					if mt.Name == fn {
						if mt.Type.PlainType !=nil {
							hashMethods = append(hashMethods,"Twox64Concat")
							mapType = IsPlainType
							break LOOP
						}
						if mt.Type.MapType != nil {
							hashMethods = append(hashMethods, mt.Type.MapType.Hasher)
							mapType = IsMap
							break LOOP
						}
						if mt.Type.DoubleMapType != nil {
							hashMethods = append(hashMethods,mt.Type.DoubleMapType.Hasher)
							hashMethods = append(hashMethods,mt.Type.DoubleMapType.Key2Hasher)
							mapType = IsDoubleMap
							break LOOP
						}
						if mt.Type.NMapType != nil {
							for _,h:=range mt.Type.NMapType.Hashers{
								hashMethods = append(hashMethods,h)
							}
							mapType = IsNMap
							break LOOP
						}
					}
				}
			}
		}
	if len(hashMethods) == 0 {
		return nil,fmt.Errorf("do not find any hash methods")
	}
	data,err:=createKey(module,fn,hashMethods,args,mapType)
	if err != nil {
		return nil,fmt.Errorf("create key error: %v",err)
	}
	if data == nil {
		return nil, errors.New("create storage key data is null")
	}
	return NewStorageKey(data), nil
}

func createKey(module,fn string, hasherMethods []string,args [][]byte,mapType int)([]byte, error){
	var (
		key []byte
		err error
	)
	if mapType!=IsPlainType {
		key,err = createNMapKey(hasherMethods,args)
		if err != nil {
			return nil, err
		}
	}
	return append(createPrefixedKey(module, fn), key...), nil
}



func createNMapKey(hasherMethods []string,args [][]byte)([]byte,error){
	if len(hasherMethods)!= len(args){
		return nil,fmt.Errorf("MapKey: len hasher=%d or len args=%d ,must be equal",len(hasherMethods),len(args))
	}
	var key []byte
	for i,method:=range hasherMethods{
		hasher,err:=utils.SelectHash(method)
		if err != nil {
			return nil, fmt.Errorf("NMapKey: select hasher error: %v",err)
		}
		_,err =hasher.Write(args[i])
		if err != nil {
			return nil,fmt.Errorf("NMapKey: hasher write error: %v",err)
		}
		subkey := hasher.Sum(nil)
		if strings.Contains(method, "Concat") {
			subkey = append(subkey, args[i]...)
		}
		key = append(key,subkey...)
	}
	return key,nil
}

func createPrefixedKey(module, fn string) []byte {
	return append(xxhash.New128([]byte(module)).Sum(nil), xxhash.New128([]byte(fn)).Sum(nil)...)
}