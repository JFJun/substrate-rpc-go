package types


// AccountInfo contains information of an account
type AccountInfo struct {
	Nonce     U32
	Consumers U32
	Providers U32
	Data      struct {
		Free       U128
		Reserved   U128
		MiscFrozen U128
		FreeFrozen U128
	}
}

type AccountInfoOld struct {
	Nonce    U32
	Refcount U32
	Data     struct {
		Free       U128
		Reserved   U128
		MiscFrozen U128
		FreeFrozen U128
	}
}

type AccountInfoWithProviders struct {
	Nonce     U32
	Consumers U32
	Providers U32
	Data      struct {
		Free       U128
		Reserved   U128
		MiscFrozen U128
		FreeFrozen U128
	}
}

type AccountInfoWithTripleRefCount struct {
	Nonce       U32
	Consumers   U32
	Providers   U32
	Sufficients U32
	Data        struct {
		Free       U128
		Reserved   U128
		MiscFrozen U128
		FreeFrozen U128
	}
}
