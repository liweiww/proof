package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"strings"
)

var (
	client *ethclient.Client
	abiObj abi.ABI
)

func init() {
	// 连接以太坊客户端
	client, _ = ethclient.Dial("http://localhost:8545") // 这里使用 Ganache 的默认地址和端口
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 智能合约地址和 ABI 文件路径
	abiPath := "D:\\Postgraduate\\assetCloud\\BlockChain\\jiaohu\\backend\\core\\experimentvocher.abi"

	// 解析 ABI 文件
	abiBytes, err := ioutil.ReadFile(abiPath)
	if err != nil {
		log.Fatal(err)
	}
	abiObj, err = abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatal(err)
	}
}
