.PHONY: all
all: compile-contracts abigen

.PHONY: compile-contracts
compile-contracts:
	solc --bin --abi -o bin --overwrite contracts/test.sol

.PHONY: abigen
abigen:
	abigen --abi bin/TestContract.abi --bin bin/TestContract.bin --pkg bindings --out bindings/bindings.go