package codec

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDeShortID(t *testing.T) {
	rand.New(rand.NewSource(time.Now().Unix()))
	// rand.Seed(time.Now().Unix())
	id := rand.Uint64()
	salt := rand.Uint64()

	g, _ := NewNameGenerator(AlphanumericSet)

	username := g.Encode(id, salt)
	fmt.Println(username)

	id2 := g.Decode(username, salt)

	fmt.Println(id, id2)

	if id != id2 {
		fmt.Println(salt)
	}
}
