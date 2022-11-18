package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// metamask 签名方式
func main() {

	//获取私钥对象
	privateKey, err := crypto.HexToECDSA("31504efef8d79d63c68922c3a142f2e7206c4f48e5b1601cee6259053b3d4cf1")
	//privateKey, err := crypto.HexToECDSA("a1a75712f9d1da8b2bfea81190f946e33cacc5a4d77db268f450914e1c9bb446")

	if err != nil {
		log.Fatal(err)
	}
	//随机数
	data := []byte("b5e4b1477dc4973fa09d1777be736dd9")
	signHash := signHash(data)
	//使用随机数进行加密。
	signature, err := crypto.Sign(signHash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	signature[64] += 27
	encode := hexutil.Encode(signature)
	fmt.Println("BigDemo:", encode)

	//对Sign进行解密
	sig, err := hexutil.Decode(encode)
	msg := accounts.TextHash([]byte("b5e4b1477dc4973fa09d1777be736dd9"))
	sig[crypto.RecoveryIDOffset] -= 27
	recovered, err := crypto.SigToPub(msg, sig)
	pubAddr := crypto.PubkeyToAddress(*recovered).String()
	fmt.Println(pubAddr)
}

func signHash(data []byte) common.Hash {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg))
}
