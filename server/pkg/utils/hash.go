package utils

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (saltStr string, ps string, err error) {
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return "", "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		3,       // 迭代次数
		64*1024, // 内存使用(64MB)
		4,       // 并行度
		32,      // 哈希长度
	)
	return base64.StdEncoding.EncodeToString(salt), base64.RawStdEncoding.EncodeToString(hash), err
}

func VerifyPassword(password string, salt string) (ps string, err error) {
	decodeString, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		decodeString,
		3,       // 迭代次数
		64*1024, // 内存使用(64MB)
		4,       // 并行度
		32,      // 哈希长度
	)

	return base64.RawStdEncoding.EncodeToString(hash), nil
}
