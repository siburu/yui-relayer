package main

import (
	"log"

	tendermint "github.com/hyperledger-labs/yui-relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
	mock "github.com/hyperledger-labs/yui-relayer/provers/mock/module"
	hd "github.com/hyperledger-labs/yui-relayer/signers/hd/module"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		hd.Module{},
		mock.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
