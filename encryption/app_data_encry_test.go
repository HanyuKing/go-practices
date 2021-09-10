package encryption

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEnDecryptData(t *testing.T) {
	rawData := "Hello World"
	cipherData, _ := EncryptData(rawData, ClientDataPublicKey, AESKey, AESIV)
	fmt.Println(cipherData)

	newRawData, _ := DecryptData("wVHqmDvXyz/U9c6XGzPMNneBGKShYJqFrhXC6kv8Bh/MhdyPWW9NIwv2L5y/mGszdMVpPv0le7nuMOnRWqDj2t302ifLMiU8tNgKMKcZHeI9Q3dimmTyCtcEQBR7GCrMV+uiP8MAHIffhdqJ5wOp06MyFK3nf39g6ChKFk1U2pE=", ClientDataPrivateKey, AESKey)
	fmt.Println(newRawData)
}

func TestSimpleAES(t *testing.T) {
	//key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
	var aeskey = []byte(AESKey)
	pass := []byte("wanghanyu")
	xpass, err := AesEncrypt(pass, aeskey, []byte("@McQfTjWnZr4u7x!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey, []byte(AESIV))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}

func TestRSAEncryption(t *testing.T) {
	rawData := "Hello World"

	/****************** start server ******************/
	rsaData, err := RsaEncrypt([]byte(rawData), ClientDataPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	base64RSAData := base64.StdEncoding.EncodeToString(rsaData)
	/****************** end server ******************/

	/****************** start client ******************/
	rsaData, err = base64.StdEncoding.DecodeString(base64RSAData)
	if err != nil {
		t.Fatal(err)
	}
	data, err := RsaDecrypt(rsaData, ClientDataPublicKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("server rawData: %s", data))

	/****************** end client ******************/
}

func TestAppEncryption(t *testing.T) {
	rawData := "Hello World"

	/****************** start client ******************/
	randomKey := AESKey
	fmt.Println(fmt.Sprintf("client randomKey: %s", randomKey))
	fmt.Println(fmt.Sprintf("client rawData: %s", rawData))

	rsaRandomKey, err := RsaEncrypt([]byte(randomKey), ClientDataPublicKey)
	if err != nil {
		t.Fatal(err)
	}

	aesData, err := AesEncrypt([]byte(rawData), randomKey, AESIV)
	if err != nil {
		t.Fatal(err)
	}

	rsaRandomKeyPass64 := base64.StdEncoding.EncodeToString(rsaRandomKey)
	fmt.Println(fmt.Sprintf("client rsaRandomKeyPass64: %s", rsaRandomKeyPass64))

	rsaData, err := RsaEncrypt(aesData, ClientDataPublicKey)
	if err != nil {
		t.Fatal(err)
	}
	rsaDataPass64 := base64.StdEncoding.EncodeToString(rsaData)
	fmt.Println(fmt.Sprintf("client rsaDataPass64: %s", rsaDataPass64))

	/****************** end client ******************/

	/****************** start server ******************/
	rsaRandomKey, err = base64.StdEncoding.DecodeString(rsaRandomKeyPass64)
	if err != nil {
		t.Fatal(err)
	}
	randomKeyByte, err := RsaDecrypt(rsaRandomKey, ClientDataPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("server randomKeyByte: %s", randomKeyByte))

	rsaData, err = base64.StdEncoding.DecodeString(rsaDataPass64)
	if err != nil {
		t.Fatal(err)
	}
	aesData, err = RsaDecrypt(rsaData, ClientDataPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	rawDataBytes, err := AesDecrypt(aesData, randomKeyByte, AESIV)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("server rawData: %s", string(rawDataBytes)))

	/****************** end server ******************/
}
