package substrate_rpc

import (
	"encoding/json"
	"fmt"
	"testing"
)

const url  = "wss://rpc.polkadot.io"

var substrate,_ = NewSubstrateAPI(url,"polkadot")

func Test_RpcChainGetBlock(t *testing.T){
	block,err:=substrate.RPC.Chain.GetBlockLatest()
	if err != nil {
		t.Fatal(err)
	}
	d,_:=json.Marshal(block)
	fmt.Println(string(d))
}

func Test_Rpc_StateGetMetadata(t *testing.T){
	md,err:=substrate.RPC.State.GetMetadataLatest()
	if err != nil {
		t.Fatal(err)
	}
	//d,_:=json.Marshal(md)
	//fmt.Println(string(d))

	d,_:=json.Marshal(md.Metadata.CallIndex)
	fmt.Println(string(d))
}



