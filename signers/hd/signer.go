package hd

import "github.com/hyperledger-labs/yui-relayer/core"

type Signer struct {
	mnemonic string
	path     *Path
}

var _ core.Signer = (*Signer)(nil)

func NewSigner(mnemonic string, path *Path) *Signer {
	return &Signer{
		mnemonic: mnemonic,
		path:     path,
	}
}

func (s *Signer) Sign(msg []byte) ([]byte, error) {
	panic("not implemented yet")
}
