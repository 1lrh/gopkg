package pwdx

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Salt generate 64-bytes salt
func Salt() string {
	salt := rand.String(44) + fmt.Sprintf("%020d", time.Now().Unix())
	return salt
}

func Encrypt(password, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Verify(password, passwordHash, salt string) bool {
	return Encrypt(password, salt) == passwordHash
}
