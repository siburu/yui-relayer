package hd

import (
	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ core.SignerConfig = (*SignerConfig)(nil)

func (c *SignerConfig) Build() (core.Signer, error) {
	return NewSigner(c.Mnemonic, c.Path), nil
}

func (c *SignerConfig) Validate() error {
	return nil
}
