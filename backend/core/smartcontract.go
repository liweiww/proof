package core

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func Smartcontract(dataName string, identificationCode string, author string, projectName string, dataVolume string, version string, dataTraceability string, creationOrganization string, dataFileHash string) string {

	//合约地址
	contractAddr := common.HexToAddress("0xaadF5F54a499b49d601DFFc9f4dD2CfE9f2df0Fa")

	input, err := abiObj.Pack("createExperimentVoucher", dataName, identificationCode, author, projectName, dataVolume, version, dataTraceability, creationOrganization, dataFileHash)
	if err != nil {
		log.Fatal(err)
	}
	//账户地址
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress("0x3BA662560C6e652B4E0c1BB76F2B63d09972131C"))
	if err != nil {
		log.Fatal(err)
	}

	// 构造交易对象
	tx := types.NewTransaction(
		nonce, // 从您的钱包中获取 nonce
		contractAddr,
		big.NewInt(0),           // 发送 0 ETH
		2000000,                 // 从 gasPrice 和 gasLimit 估算得到
		big.NewInt(20000000000), // 从您的钱包中获取 gasPrice
		input,
	)

	// 签名交易
	privateKey, err := crypto.HexToECDSA("e8dc3271ba7ea344e4338899afbe05bdac22db3074febcfb50a3b970967d2d7b")
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// 等待交易被打包
	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("transaction failed")
	} else if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("transaction successful")
	}
	if len(receipt.Logs) == 0 {
		log.Fatal("no logs found in receipt")
	}
	var logs []string
	for _, log := range receipt.Logs {
		logs = append(logs, fmt.Sprintf("%+v", log))
	}
	logsStr := strings.Join(logs, "\n")
	return logsStr

}

func CreateExperimentVoucher(dataName string, identificationCode string, author string, projectName string, dataVolume string, version string, dataTraceability string, creationOrganization string, dataFileHash string) {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("e8dc3271ba7ea344e4338899afbe05bdac22db3074febcfb50a3b970967d2d7b")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
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
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x29fE5A13483975D80A85CDA92876256E671491Ee")
	instance, err := NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//key := [32]byte{}
	//value := [32]byte{}
	//copy(key[:], []byte("foo"))
	//copy(value[:], []byte("bar"))

	tx, err := instance.CreateExperimentVoucher(auth, dataName, identificationCode, author, projectName, dataVolume, version, dataTraceability, creationOrganization, "123")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction hash:", tx.Hash().Hex())

	// Wait for the transaction to be mined
	ctx := context.Background()
	tx2, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction hash:", tx2.TxHash.Hex())

	for _, event := range tx2.Logs {
		fmt.Println(event)
	}
	//jsonBytes, err := json.Marshal(tx)
	//if err != nil {
	//	fmt.Println("Failed to convert struct to JSON:", err)
	//	return
	//}
	//
	//jsonStr := string(jsonBytes)
	//fmt.Println("结构体", jsonStr)
	//	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

}

type result struct {
	DataName             string
	IdentificationCode   string
	Author               string
	ProjectName          string
	DataVolume           string
	Version              string
	DataTraceability     string
	CreationOrganization string
	DataFileHash         string
	Timestamp            string
}

func GetExperiment(hexStr string) result {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xaadF5F54a499b49d601DFFc9f4dD2CfE9f2df0Fa")
	instance, err := NewContract(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// 将十六进制字符串转换为字节数组
	//hexStr := "3075ae7ea5af4597356cf2df096029b1cbd0e59598cca2232fa6eecdca4353b2"
	byteArray, err := hex.DecodeString(hexStr[2:])
	if err != nil {
		fmt.Println("Failed to decode hex string:", err)
	}

	// 将字节数组复制到 [32]byte 数组
	var byteArray32 [32]byte
	copy(byteArray32[:], byteArray)

	// 输出 [32]byte 数组元素
	fmt.Println("地址", byteArray32)
	//experimentId := [32]byte{0xcf, 0x3b, 0x9c, 0xfe, 0x71, 0x8f, 0x65, 0xa2, 0xb2, 0x38, 0x8e, 0x8b, 0x8e, 0xaa, 0x4c, 0x9b, 0xe8, 0x30, 0x1d, 0xed, 0x11, 0xe0, 0x4b, 0x1e, 0xe7, 0x72, 0xb7, 0x23, 0xa1, 0x9f, 0x93, 0x85}
	version, err := instance.GetExperiment(nil, byteArray32)
	if err != nil {
		log.Fatal(err)
	}

	t := time.Unix(version.Timestamp.Int64(), 0)
	// 将时间格式化为字符串
	str := t.Format("2006-01-02 15:04:05")
	fmt.Println("结构", version) // "1.0"
	return result{
		DataFileHash:         version.DataFileHash,
		DataName:             version.DataName,
		DataTraceability:     version.DataTraceability,
		DataVolume:           version.DataVolume,
		IdentificationCode:   version.IdentificationCode,
		Author:               version.Author,
		Timestamp:            str,
		Version:              version.Version,
		ProjectName:          version.ProjectName,
		CreationOrganization: version.CreationOrganization,
	}
}

func QueryLog() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xaadF5F54a499b49d601DFFc9f4dD2CfE9f2df0Fa")
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	fmt.Println("日志", logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(ContractABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		event := struct {
			experimentId    [32]byte
			transactionHash [32]byte
		}{}
		//err := contractAbi.Unpack(event, eventName, log.Data)
		err := contractAbi.UnpackIntoInterface(&event, "ExperimentCreated", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.experimentId[:]))    // foo
		fmt.Println(string(event.transactionHash[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	//eventSignature := []byte("ExperimentCreated(bytes32,bytes32)")
	//hash := crypto.Keccak256Hash(eventSignature)
	//fmt.Println(hash.Hex())
}
