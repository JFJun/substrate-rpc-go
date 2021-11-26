package types

import (
	"encoding/json"
	"fmt"
	scalecode "github.com/itering/scale.go"
	"github.com/itering/scale.go/types"
	"github.com/itering/scale.go/utiles"
)

func DecodeExtrinsic(meta *scalecode.MetadataDecoder, extrinsic string) (data []byte, err error) {
	defer func() {
		if errS := recover(); errS != nil {
			err = fmt.Errorf("decode extrinsic have panic: %v", errS)
		}
	}()
	e := scalecode.ExtrinsicDecoder{}
	option := types.ScaleDecoderOption{Metadata: &meta.Metadata}
	e.Init(types.ScaleBytes{Data: utiles.HexToBytes(extrinsic)}, &option)
	e.Process()
	data, err = json.Marshal(e.Value)
	return
}
