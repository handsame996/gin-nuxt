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
	return saltStr, base64.RawStdEncoding.EncodeToString(hash), err
}
