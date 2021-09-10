package encryption

import "encoding/base64"

func EncryptData(originData string, publicKey string, aesKey, aesIV []byte) (string, error) {

	aesData, err := AesEncrypt([]byte(originData), aesKey, aesIV)
	if err != nil {
		return "", err
	}

	rsaData, err := RsaEncrypt(aesData, publicKey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(rsaData), nil
}

func DecryptData(cipherText string, privateKey string, aesKey []byte) (string, error) {

	rsaData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", nil
	}
	aesData, err := RsaDecrypt(rsaData, privateKey)
	if err != nil {
		return "", nil
	}
	rawDataBytes, err := AesDecrypt(aesData, aesKey, AESIV)
	if err != nil {
		return "", nil
	}
	return string(rawDataBytes), nil
}
