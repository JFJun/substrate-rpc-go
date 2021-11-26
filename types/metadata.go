package types

import (
	"fmt"
	scalecodec "github.com/itering/scale.go"
	"github.com/itering/scale.go/utiles"
)

func DecodeMetadata(meta string) (_ *scalecodec.MetadataDecoder, err error) {

	// Process() 有可能传回panic
	defer func() {
		if re := recover(); re != nil {
			err = fmt.Errorf("process metadata panic: %v", re)
		}
	}()

	em := scalecodec.MetadataDecoder{}
	em.Init(utiles.HexToBytes(meta))
	if err := em.Process(); err != nil {
		return nil, fmt.Errorf("decode metadata error: %v", err)
	}

	return &em, nil
}
