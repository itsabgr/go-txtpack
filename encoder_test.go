package txtpack

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/itsabgr/go-handy"
	"io"
	"testing"
)

func randBytes(len int) []byte {
	buf := make([]byte, len)
	_, err := io.ReadFull(rand.Reader, buf)
	handy.Throw(err)
	return buf
}
func randPairs(n int) map[string][]byte {
	pairs := make(map[string][]byte, 0)
	for range handy.N(uint(n)) {
		pairs[hex.EncodeToString(randBytes(10))] = []byte(hex.EncodeToString(randBytes(10)))
	}
	return pairs
}
func equalMaps(a, b map[string][]byte) bool {
	for k, v := range a {
		bv, ok := b[k]
		if !ok {
			return false
		}
		if !bytes.Equal(bv, v) {
			return false
		}
	}
	return true
}
func TestEncoder(t *testing.T) {
	for range handy.N(10) {
		pairs1 := randPairs(10)
		pairs2 := randPairs(10)
		buf := new(bytes.Buffer)
		err := Marshal(buf, pairs1)
		if err != nil {
			t.Fatal(err)
		}
		err = Marshal(buf, pairs2)
		if err != nil {
			t.Fatal(err)
		}
		pairs3, err := Unmarshal(bufio.NewReader(buf))
		if err != nil {
			t.Fatal(err)
		}
		if !equalMaps(pairs3, pairs1) {
			t.Fatal(fmt.Errorf("pairs are not equal"))
		}
		pairs4, err := Unmarshal(bufio.NewReader(buf))
		if err != nil {
			t.Fatal(err)
		}
		if !equalMaps(pairs4, pairs2) {
			t.Fatal(fmt.Errorf("pairs are not equal"))
		}
	}
}
