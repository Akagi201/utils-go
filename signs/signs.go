package signs

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func GetMd5Sign(secret string, params string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(params))

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GetSha1Sign(params string) (string, error) {
	sha := sha1.New()
	_, err := sha.Write([]byte(params))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha.Sum(nil)), nil
}

func GetSha256Sign(secret string, params string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func GetSha512Sign(secret string, params string) (string, error) {
	mac := hmac.New(sha512.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func GetHmacSha1Sign(secret string, params string) (string, error) {
	mac := hmac.New(sha1.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func GetHmacMd5Sign(secret string, params string) (string, error) {
	mac := hmac.New(md5.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func GetHmacSha384Sign(secret string, params string) (string, error) {
	mac := hmac.New(sha512.New384, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", nil
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

func GetHmacSha256B64Sign(secret string, params string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(params))
	if err != nil {
		return "", err
	}
	signByte := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signByte), nil
}

func GetHmacSha512B64Sign(hmacKey string, hmacData string) string {
	hmh := hmac.New(sha512.New, []byte(hmacKey))
	hmh.Write([]byte(hmacData))

	hex_data := hex.EncodeToString(hmh.Sum(nil))
	hash_hmac_bytes := []byte(hex_data)
	hmh.Reset()

	return base64.StdEncoding.EncodeToString(hash_hmac_bytes)
}
