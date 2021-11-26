package substrate_rpc

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/substrate-rpc-go/client"
	"github.com/JFJun/substrate-rpc-go/rpc"
	srgSource "github.com/JFJun/substrate-rpc-go/source"
	"github.com/itering/scale.go/source"
	"github.com/itering/scale.go/types"
)

type SubstrateAPI struct {
	RPC    *rpc.RPC
	Client client.Client
}
/*
url: 节点url
coinType： https://github.com/itering/scale.go/blob/master/network/ 下的json文件名字
*/
func NewSubstrateAPI(url,coinType string) (*SubstrateAPI, error) {
	cl, err := client.Connect(url)
	if err != nil {
		return nil, err
	}

	newRPC, err := rpc.NewRPC(cl)
	if err != nil {
		return nil, err
	}
	// 注册source
	types.RegCustomTypes(source.LoadTypeRegistry([]byte(srgSource.BaseType)))
	//注册network
	nwBytes,err:=getNetworkRegistryTypes(coinType)
	if err != nil {
		return nil, err
	}
	types.RegCustomTypes(source.LoadTypeRegistry(nwBytes))

	return &SubstrateAPI{
		RPC:    newRPC,
		Client: cl,
	}, nil
}

func getNetworkRegistryTypes(coinType string)([]byte,error){
	if coinType=="" {
		coinType = "polkadot"
	}
	var nwMap map[string]map[string]interface{}
	err := json.Unmarshal([]byte(srgSource.NetworkType),&nwMap)
	if err != nil {
		return nil, err
	}
	for key,value:=range nwMap{
		if key==coinType {
			return json.Marshal(value)
		}
	}
	return nil, fmt.Errorf("do not suppoort this cointype: %s",coinType)
}