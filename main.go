package main

import (
	"context"
	"log"
	"math/big"

	"github.com/bonedaddy/bdsm/testenv"
	"github.com/bonedaddy/geth-pending-state-test/bindings"
)

func main() {
	ctx := context.TODO()
	testenv, err := testenv.NewBlockchain(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, tx, bindings, err := bindings.DeployBindings(testenv.Auth, testenv)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := testenv.DoWaitDeployed(tx, "test contract deployment")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = bindings.SetValid(testenv.Auth, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	if err := testenv.DoWaitMined(tx); err != nil {
		log.Fatal(err)
	}

	// should be 1
	supply, err := bindings.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}
	if supply.Int64() != 1 {
		log.Fatal("supply should be 1")
	}
	if err := testenv.SendETH(addr, big.NewInt(100000)); err != nil {
		log.Fatal(err)
	}
	// should be 2
	supply, err = bindings.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}
	if supply.Int64() != 2 {
		log.Fatal("error: supply should be 2")
	}
}
