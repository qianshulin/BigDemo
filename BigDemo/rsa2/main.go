package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type ERC20 struct {
	Type        string `bson:"type" json:"-"`
	Name        string `bson:"name" json:"name"`
	TxHash      string `bson:"txhash" json:"txhash"`
	From        string `bson:"fromaddress" json:"fromaddress"`
	To          string `bson:"toaddress" json:"toaddress"`
	Amount      string `bson:"price" json:"price"`
	BlockNumber int64  `bson:"blocknumber" json:"blocknumber"`
	Time        int64  `bson:"time" json:"time"`
	UserAddr    string `bson:"addr" json:"addr"`
	GoldRate    string `bson:"goldrate"json:"goldrate"`
}

var pubKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDCF5neFxaxrWDDUbdtAX12LXjC\nX+uDiQherMeRLE5amZWCMwyeU7SxlPyOFXYMQPzCIHdxC1ZQnRl8zEYXtP5DVyn6\nEXcGTb/NPUvf2VoVzZ1ndHwUT48U0vnXQLa2jqmEVDoglhJU16e75h+thYSBZywZ\n9DbR2KsK64SJ0KVvgwIDAQAB\n-----END PUBLIC KEY-----\n"
var prvKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICWwIBAAKBgQDCF5neFxaxrWDDUbdtAX12LXjCX+uDiQherMeRLE5amZWCMwye\nU7SxlPyOFXYMQPzCIHdxC1ZQnRl8zEYXtP5DVyn6EXcGTb/NPUvf2VoVzZ1ndHwU\nT48U0vnXQLa2jqmEVDoglhJU16e75h+thYSBZywZ9DbR2KsK64SJ0KVvgwIDAQAB\nAoGAGuER2gPwjKKqJ+KFOH9gVKFve9u0zf6IPjRHXv93ymxCnEldkf+dooozf6nR\nM1k8p9W/NNSRG5DmGrZfbqEOJR9geQ0fbH8H9J4epuAZXzXAxURhBCOe1Cxoc6um\nvjEGDB0CH4jz79Ce+x3goU9WFjGrQwvZxxLXHMcJGS+tS6ECQQDv/m0L2Gjkq4zj\nxoRbASpM8QFhwyaEY4nLI3p8t0Fu01yaQfbmQ36YqfzCVuBcZuPuEzfvw6PnTpd8\nNCyvEa5/AkEAzwl3M6OIlNesmmp6GrdC2ph00h8/H13ebgoItUM2oDmFn5452MdB\nhuULnGbHhTKrlUMsDj4RuBjhJypyxY4E/QJATxzsDY+VZGw40Y1LbESCEUMRVYzj\nUZkOf2x+oEY4x7PdPuyn21g8j07aB6Zj55HzaTPkqKExkrqRrVv1rvCBOQJAQhW8\nZdBCTXhWUOfB1/s9LLdTOqrcpC97S1aKlwlAS35w55VC064ufVdpEGBCOYMF/9v0\nGQu47jGK1MyY2/RMJQJAUdeUK9Z/PweC8csUJ2pAtgMb8ioLDtOAp+5nZaxHpZmx\nC9a3etfFSTuERgCt0ezDw1FfCL1Sa7ZXuXreQw3bZw==\n-----END RSA PRIVATE KEY-----\n"

func main() {
	erc := ERC20{
		Type:        "20",
		Name:        "ETH",
		TxHash:      "Hash",
		BlockNumber: 11,
		To:          "to",
		From:        "From",
		Amount:      "1",
		Time:        time.Now().Unix(),
		UserAddr:    "测试",
		GoldRate:    "100",
	}
	data, _ := json.Marshal(erc)
	block, err := RsaEncryptBlock(data, []byte(pubKey))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hex.EncodeToString(block))
	//解密
	decryptBlock, err := RsaDecryptBlock(block, []byte(prvKey))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decryptBlock)

}

//RsaEncryptBlock 公钥加密-分段
func RsaEncryptBlock(src, publicKeyByte []byte) (bytesEncrypt []byte, err error) {
	block, _ := pem.Decode(publicKeyByte)
	if block == nil {
		return nil, errors.New("获取公钥失败")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	keySize, srcSize := publicKey.(*rsa.PublicKey).Size(), len(src)
	pub := publicKey.(*rsa.PublicKey)
	//logs.Debug("密钥长度：", keySize, "\t明文长度：\t", srcSize)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, pub, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt = buffer.Bytes()
	return
}

//RsaDecryptBlock 私钥解密
func RsaDecryptBlock(src, privateKeyByte []byte) (bytesDecrypt []byte, err error) {
	block, _ := pem.Decode(privateKeyByte)
	if block == nil {
		return nil, errors.New("获取私钥失败")
	}
	//ParsePKCS8PrivateKey
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	keySize, srcSize := privateKey.Size(), len(src)
	//logs.Debug("密钥长度：", keySize, "\t密文长度：\t", srcSize)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt = buffer.Bytes()
	return
}
