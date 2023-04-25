package core

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	CreateExperimentVoucher("a", "1", "1", "1", "1", "1", "1", "2", "2")
}

func TestGetExperiment(t *testing.T) {
	QueryLog()
}

func TestN(t *testing.T) {
	logsStr := Smartcontract("a", "a", "a", "a", "a", "a", "a", "a", "a")
	logs := strings.Split(logsStr, " ")
	experimentId := logs[2]
	transactionId := logs[3]
	//txHash := logs[6]
	fmt.Println("experimentId:", experimentId)
	fmt.Println("transactionId:", transactionId)
}
