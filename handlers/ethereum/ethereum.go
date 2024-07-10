package ethereum

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var (
	rpcURL             string
	privateKeyStr      string
	contractAddressStr string
	contractABI        string
)

var client *ethclient.Client
var contractInstance *bind.BoundContract
var contractAddress common.Address
var parsedABI abi.ABI

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpcURL = os.Getenv("RPC_URL")
	privateKeyStr = os.Getenv("PRIVATE_KEY")
	contractAddressStr = os.Getenv("CONTRACT_ADDRESS")

	abiFile, err := ioutil.ReadFile("./abi/FileRegistryABI.json")
	if err != nil {
		log.Fatal("Failed to read ABI file: ", err)
	}
	contractABI = string(abiFile)

	client, err = ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum client: ", err)
	}

	parsedABI, err = abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal("Failed to parse ABI: ", err)
	}

	contractAddress = common.HexToAddress(contractAddressStr)
	contractInstance = bind.NewBoundContract(contractAddress, parsedABI, client, client, client)
}

func TransactWithContract(function string, params ...interface{}) error {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return err
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	_, err = contractInstance.Transact(auth, function, params...)
	return err
}

func CallContractFunction(function string, params ...interface{}) (string, error) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	var results []interface{}
	err := contractInstance.Call(callOpts, &results, function, params...)
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no result returned from contract function call")
	}

	resultString, ok := results[0].(string)
	if !ok {
		return "", fmt.Errorf("result type assertion to string failed")
	}

	return resultString, nil
}
