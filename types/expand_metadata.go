package types

import (
	"errors"
	"fmt"
	"github.com/JFJun/substrate-rpc-go/utils"
	sType "github.com/itering/scale.go/types"
)

/*
func: 扩展metadata功能，使其快速实现
author： flynn
date: 2021/11/25
*/

type ExpandMetadata struct {
	meta *sType.MetadataStruct
}

func NewExpandMetadata(meta *sType.MetadataStruct) *ExpandMetadata {
	me := new(ExpandMetadata)
	me.meta = meta
	return me
}

func (em *ExpandMetadata) GetCallIndex(moduleName, fn string) (callIdx string, err error) {
	for idx, call := range em.meta.CallIndex {
		if call.Module.Name == moduleName {
			if call.Call.Name == fn {
				return idx, nil
			}
		}
	}
	return "", fmt.Errorf("do not find any callIdx by this module=%s,call name=%s ", moduleName, fn)
}

func (em *ExpandMetadata) FindNameByCallIndex(callIdx string) (moduleName, fn string, err error) {
	for idx, call := range em.meta.CallIndex {
		if idx == callIdx {
			return call.Module.Name, call.Call.Name, nil
		}
	}
	return "", "", fmt.Errorf("do not find this callIdx=%s ", callIdx)
}

//func (em *ExpandMetadata) GetConstants(modName, constantsName string) (constantsType string, constantsValue []byte, err error) {
//	return "", nil, fmt.Errorf("do not find constants by this mod name=%s,constants name=%s",modName,constantsName)
//}

func (em *ExpandMetadata) BalanceTransferCall(to string, amount uint64) (Call, error) {
	var (
		call Call
	)
	callIdx, err := em.GetCallIndex("Balances", "transfer")
	if err != nil {
		return call, err
	}
	recipientPubkey := utils.AddressToPublicKey(to)
	ma, err := NewMultiAddressFromHexAccountID(recipientPubkey)
	if err != nil {
		return call, err
	}
	return NewCall(callIdx, ma, NewUCompactFromUInt(amount))
}

func (em *ExpandMetadata) BalanceTransferKeepAliveCall(to string, amount uint64) (Call, error) {
	var (
		call Call
	)
	callIdx, err := em.GetCallIndex("Balances", "transfer_keep_alive")
	if err != nil {
		return call, err
	}
	recipientPubkey := utils.AddressToPublicKey(to)
	ma, err := NewMultiAddressFromHexAccountID(recipientPubkey)
	if err != nil {
		return call, err
	}
	return NewCall(callIdx, ma, NewUCompactFromUInt(amount))
}

func (em *ExpandMetadata) UtilityBatchTxCall(toAmount map[string]uint64, keepAlive bool) (Call, error) {
	var (
		call Call
		err  error
	)
	if len(toAmount) == 0 {
		return call, errors.New("toAmount is null")
	}
	var calls []Call
	for to, amount := range toAmount {
		var (
			btCall Call
		)
		if keepAlive {
			btCall, err = em.BalanceTransferKeepAliveCall(to, amount)
		} else {
			btCall, err = em.BalanceTransferCall(to, amount)
		}
		if err != nil {
			return call, err
		}
		calls = append(calls, btCall)
	}
	callIdx, err := em.GetCallIndex("Utility", "batch")
	if err != nil {
		return call, err
	}
	return NewCall(callIdx, calls)
}

/*
transfer with memo
*/
func (em *ExpandMetadata) UtilityBatchTxWithMemo(to, memo string, amount uint64) (Call, error) {
	var (
		call Call
	)
	btCall, err := em.BalanceTransferCall(to, amount)
	if err != nil {
		return call, err
	}
	smCallIdx, err := em.GetCallIndex("System", "remark")
	if err != nil {
		return call, err
	}
	smCall, err := NewCall(smCallIdx, memo)
	ubCallIdx, err := em.GetCallIndex("Utility", "batch")
	if err != nil {
		return call, err
	}
	return NewCall(ubCallIdx, btCall, smCall)
}
