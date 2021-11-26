package state

import (
	"fmt"
	"github.com/JFJun/substrate-rpc-go/client"
	"github.com/JFJun/substrate-rpc-go/types"
	scalecodec "github.com/itering/scale.go"
	"github.com/itering/scale.go/utiles"
)

func (s *State) GetMetadataExpand(blockHash types.Hash) (*types.ExpandMetadata, error) {
	return s.getMetadataExpand(&blockHash)
}

func (s *State) GetMetadataLatestExpand() (*types.ExpandMetadata, error) {
	return s.getMetadataExpand(nil)
}

func (s *State) getMetadataExpand(blockHash *types.Hash) (md *types.ExpandMetadata, err error) {
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
			err = fmt.Errorf("process metadata panic: %v", re)
		}
	}()

	m := scalecodec.MetadataDecoder{}
	m.Init(utiles.HexToBytes(res))
	if err := m.Process(); err != nil {
		return nil, fmt.Errorf("decode metadata error: %v", err)
	}
	em := types.NewExpandMetadata(&m.Metadata)
	return em, nil
}
