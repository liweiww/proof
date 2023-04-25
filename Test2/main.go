package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 文件1和文件2的哈希值
	h1 := sha256.Sum256([]byte("file1"))
	h2 := sha256.Sum256([]byte("file1"))

	// 生成随机数r作为证明的随机数
	r := make([]byte, 32)
	_, err := rand.Read(r)
	if err != nil {
		panic(err)
	}

	// 随机选择一个文件块的索引i
	i := 0

	// 计算文件1和文件2中第i个块的哈希值
	h1_i := sha256.Sum256([]byte(fmt.Sprintf("file1 block %d", i)))
	h2_i := sha256.Sum256([]byte(fmt.Sprintf("file2 block %d", i)))

	// 计算证据a、b、c
	var a, b, c [32]byte
	for j := 0; j < 32; j++ {
		a[j] = h1_i[j] ^ r[j]
		b[j] = h2_i[j] ^ r[j]
		c[j] = h1[j] ^ h2[j]
	}

	// 将a、b、c发送给验证者进行验证
	// 验证 a ^ b = c，即验证 a 和 b 的异或结果是否等于 c
	if verify(a, b, c) {
		fmt.Println("两个文件是相同的")
	} else {
		fmt.Println("两个文件不相同")
	}
}

// 零知识证明验证函数
func verify(a, b, c [32]byte) bool {
	for i := 0; i < 32; i++ {
		if a[i]^b[i] != c[i] {
			return false
		}
	}
	return true
}
