package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	// 文件 A 和文件 B 的内容
	fileA := []byte("File Content")
	fileB := []byte("File Content")

	// 文件 A 和文件 B 的哈希值
	hashA := sha256.Sum256(fileA)
	hashB := sha256.Sum256(fileB)

	// 生成随机数 r
	r, err := rand.Int(rand.Reader, new(big.Int).SetInt64(100))
	if err != nil {
		fmt.Println("生成随机数失败：", err)
		return
	}

	// 计算挑战值 c = H(r || hashA || hashB)
	h := sha256.New()
	h.Write(r.Bytes())
	h.Write(hashA[:])
	h.Write(hashB[:])
	c := new(big.Int).SetBytes(h.Sum(nil))

	// 计算响应值 sA = r + c * privateKeyA，sB = r + c * privateKeyB
	privateKeyA := new(big.Int).SetBytes(hashA[:])
	sA := new(big.Int).Mul(c, privateKeyA)
	sA.Add(sA, r)

	privateKeyB := new(big.Int).SetBytes(hashB[:])
	sB := new(big.Int).Mul(c, privateKeyB)
	sB.Add(sB, r)

	// 发送挑战值 c 和响应值 sA, sB 给验证方
	fmt.Println("发送给验证方的挑战值 c:", c)
	fmt.Println("发送给验证方的响应值 sA:", sA)
	fmt.Println("发送给验证方的响应值 sB:", sB)

	// 验证方验证零知识证明
	// 计算哈希值哈希A' 和哈希值哈希B'
	hashAPrime := sha256.Sum256(append(r.Bytes(), hashA[:]...))
	hashBPrime := sha256.Sum256(append(r.Bytes(), hashB[:]...))

	// 计算挑战值 c' = H(r || hashA' || hashB')
	hPrime := sha256.New()
	hPrime.Write(r.Bytes())
	hPrime.Write(hashAPrime[:])
	hPrime.Write(hashBPrime[:])
	cPrime := new(big.Int).SetBytes(hPrime.Sum(nil))

	if c.Cmp(cPrime) == 0 && sA.Cmp(sB) == 0 {
		fmt.Println("验证通过，文件 A 和文件 B 是相同的")
	} else {
		fmt.Println("验证失败，文件 A 和文件 B 不相同")
	}
}
