package pwdx

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"testing"
)

func TestEncrypt(t *testing.T) {

}

func TestSaltLength(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		salt := Salt()
		if len(salt) != 64 {
			t.Fatal(salt, len(salt))
		}
	}
}

func TestDecrypt(t *testing.T) {
	salt := Salt()
	password := "123fasdljfkasdghlasdhg;lasdkhga;lsdjgklasdghasdg;jlksadhgl;kdshg;asgsad;ghasl;dghkkklsdahglksadhglstuewpt"
	passwordHash := Encrypt(password, salt)
	fmt.Println(len(passwordHash))

	ok := Verify(password, passwordHash, salt)

	fmt.Println("salt: ", salt)
	fmt.Println("passwordHash", passwordHash)
	fmt.Println(ok)
}

func TestHashLength(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		password := rand.String(32)
		salt := Salt()
		hash := Encrypt(password, salt)
		if len(hash) != 64 {
			t.Fatal(hash, len(hash))
		}
	}
}
