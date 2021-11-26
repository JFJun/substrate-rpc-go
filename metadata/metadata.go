package metadata

import (
	"errors"
	"fmt"
	"github.com/JFJun/substrate-rpc-go/types"
	"github.com/JFJun/substrate-rpc-go/utils"
	scalecodec "github.com/itering/scale.go"

)

/*
func: 扩展metadata功能，使其快速实现
author： flynn
date: 2021/11/25
*/


type ExpandMetadata struct {
	meta *scalecodec.MetadataDecoder

}

func NewExpandMetadata(meta *scalecodec.MetadataDecoder) *ExpandMetadata {
	me := new(ExpandMetadata)
	me.meta = meta
	return me
}

func (em *ExpandMetadata)GetCallIndex(moduleName, fn string) (callIdx string, err error){
	for idx,call :=range em.meta.Metadata.CallIndex{
		if call.Module.Name==moduleName {
			if call.Call.Name==fn {
				return idx,nil
			}
		}
	}
	return "", fmt.Errorf("do not find any callIdx by this module=%s,call name=%s ",moduleName,fn)
}

func (em *ExpandMetadata)FindNameByCallIndex(callIdx string) (moduleName, fn string, err error) {
	for idx,call :=range em.meta.Metadata.CallIndex{
		if idx==callIdx {
			return call.Module.Name,call.Call.Name,nil
		}
	}
	return "", "", fmt.Errorf("do not find this callIdx=%s ",callIdx)
}

//func (em *ExpandMetadata) GetConstants(modName, constantsName string) (constantsType string, constantsValue []byte, err error) {
//	return "", nil, fmt.Errorf("do not find constants by this mod name=%s,constants name=%s",modName,constantsName)
//}


func (em *ExpandMetadata) BalanceTransferCall(to string, amount uint64) (types.Call, error) {
	var (
		call types.Call
	)
	callIdx,err:=em.GetCallIndex("Balances", "transfer")
	if err != nil {
		return call, err
	}
	recipientPubkey := utils.AddressToPublicKey(to)
	ma,err:=types.NewMultiAddressFromHexAccountID(recipientPubkey)
	if err != nil {
		return call,err
	}
	return types.NewCall(callIdx,ma,types.NewUCompactFromUInt(amount))
}

func (em *ExpandMetadata) BalanceTransferKeepAliveCall(to string, amount uint64) (types.Call, error) {
	var (
		call types.Call
	)
	callIdx,err:=em.GetCallIndex("Balances", "transfer_keep_alive")
	if err != nil {
		return call, err
	}
	recipientPubkey := utils.AddressToPublicKey(to)
	ma,err:=types.NewMultiAddressFromHexAccountID(recipientPubkey)
	if err != nil {
		return call,err
	}
	return types.NewCall(callIdx,ma,types.NewUCompactFromUInt(amount))
}


func (em *ExpandMetadata) UtilityBatchTxCall(toAmount map[string]uint64, keepAlive bool) (types.Call, error) {
	var (
		call types.Call
		err  error
	)
	if len(toAmount) == 0 {
		return call, errors.New("toAmount is null")
	}
	var calls []types.Call
	for to, amount := range toAmount {
		var (
			btCall types.Call
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
	return types.NewCall(callIdx, calls)
}

/*
transfer with memo
*/
func (em *ExpandMetadata) UtilityBatchTxWithMemo(to, memo string, amount uint64) (types.Call, error) {
	var (
		call types.Call
	)
	btCall, err := em.BalanceTransferCall(to, amount)
	if err != nil {
		return call, err
	}
	smCallIdx, err := em.GetCallIndex("System", "remark")
	if err != nil {
		return call, err
	}
	smCall, err := types.NewCall(smCallIdx, memo)
	ubCallIdx, err := em.GetCallIndex("Utility", "batch")
	if err != nil {
		return call, err
	}
	return types.NewCall(ubCallIdx, btCall, smCall)
}
