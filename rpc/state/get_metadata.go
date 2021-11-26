// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"fmt"
	"github.com/JFJun/substrate-rpc-go/client"
	"github.com/JFJun/substrate-rpc-go/types"
	scalecodec "github.com/itering/scale.go"
	"github.com/itering/scale.go/utiles"
)

// GetMetadata returns the metadata at the given block
func (s *State) GetMetadata(blockHash types.Hash) (*scalecodec.MetadataDecoder, error) {
	return s.getMetadata(&blockHash)
}

// GetMetadataLatest returns the latest metadata
func (s *State) GetMetadataLatest() (*scalecodec.MetadataDecoder, error) {
	return s.getMetadata(nil)
}

func (s *State) getMetadata(blockHash *types.Hash) (md *scalecodec.MetadataDecoder,err error) {
	var (
		res string

	)
	err = client.CallWithBlockHash(s.client, &res, "state_getMetadata", blockHash)
	if err != nil {
		return nil, err
	}
	// Process() 有可能传回panic
	defer func() {
		if re := recover(); re != nil {
			err = fmt.Errorf("process metadata panic: %v",re)
		}
	}()

	m := scalecodec.MetadataDecoder{}
	m.Init(utiles.HexToBytes(res))
	if err := m.Process(); err != nil {
		return nil, fmt.Errorf("decode metadata error: %v",err)
	}
	return &m,nil
}
