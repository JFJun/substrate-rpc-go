package types

import (
	scalecodec "github.com/itering/scale.go"
	"sync"
)

// SerDeOptions are serialise and deserialize options for types
type SerDeOptions struct {
	// NoPalletIndices enable this to work with substrate chains that do not have indices pallet in runtime
	NoPalletIndices bool
}

var defaultOptions = SerDeOptions{}
var mu sync.RWMutex

// SetSerDeOptions overrides default serialise and deserialize options
func SetSerDeOptions(so SerDeOptions) {
	defer mu.Unlock()
	mu.Lock()
	defaultOptions = so
}

// SerDeOptionsFromMetadata returns Serialise and deserialize options from metadata
func SerDeOptionsFromMetadata(meta *scalecodec.MetadataDecoder) SerDeOptions {
	var opts SerDeOptions
	noPalletIndices:=false
	for _,mod:=range meta.Metadata.Metadata.Modules{
		if mod.Name=="Indices" {
			noPalletIndices = true
		}
	}
	if !noPalletIndices {
		opts.NoPalletIndices = true
	}

	return opts
}
