package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"seb7887/goeth/internal/store"
	"seb7887/goeth/internal/token"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyEDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyEDSA)
	log.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in Wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	var (
		key [32]byte
		value [32]byte
	)

	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("tx sent")
	log.Println(tx.Hash().Hex())

	res, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(res[:]))

	tknAddress := common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
	tknInstance, err := token.NewToken(tknAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	deployer := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	bal, err := tknInstance.BalanceOf(&bind.CallOpts{}, deployer)
	if err != nil {
		log.Fatal(err)
	}

	name, err := tknInstance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := tknInstance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := tknInstance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(name)
	log.Println(symbol)
	log.Println(decimals)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	balance := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	log.Println(balance)
}