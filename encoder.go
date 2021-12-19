package txtpack

import (
	"bytes"
	"io"
	"strings"
)

const NL = "\n"
const COLON = ":"
const SEP = COLON + " "
const NLChar = '\n'

type Pair struct {
	key, val string
}

func P(key, val string) Pair {
	return Pair{key, val}
}
func (p Pair) Key() string {
	return p.key
}
func (p Pair) Value() string {
	return p.val
}

func (p Pair) HasValue() bool {
	return p.Value() != ""
}

type Pairs []Pair

func MapToPairs(m map[string]string) Pairs {
	pairs := make(Pairs, 0, len(m))
	for k, v := range m {
		pairs = pairs.Append(P(k, v))
	}
	return pairs
}
func NewPairs(pairs ...Pair) Pairs {
	return append(make(Pairs, 0, len(pairs)), pairs...)
}
func DecodePair(str string) Pair {
	nlIndex := strings.Index(str, NL)
	if nlIndex < 0 {
		nlIndex = len(str)
	}
	sepIndex := strings.Index(str, SEP)
	if sepIndex < 0 {
		return P(str[0:nlIndex], "")
	}
	return P(str[0:sepIndex], str[sepIndex+2:nlIndex])
}
func DecodePack(src interface {
	ReadBytes(delim byte) ([]byte, error)
}) (Pairs, error) {
	pairs := NewPairs()
	for {
		line, err := src.ReadBytes(NLChar)
		if err != nil {
			return pairs, err
		}
		if line == nil || len(line) == 0 || string(line) == "\n" {
			break
		}
		pairs = pairs.Append(DecodePair(string(line)))
	}
	return pairs, nil
}
func (pairs *Pairs) EncodeTo(dst io.StringWriter) {
	for _, pair := range *pairs {
		if !pair.HasValue() {
			continue
		}
		dst.WriteString(pair.Key())
		dst.WriteString(SEP)
		dst.WriteString(pair.Value())
		dst.WriteString(NL)
	}
	dst.WriteString("\n")
}
func (pairs *Pairs) Encode() string {
	dst := new(bytes.Buffer)
	pairs.EncodeTo(dst)
	return dst.String()
}
func (pairs Pairs) Get(key string) string {
	if pairs.Count() == 0 {
		return ""
	}
	for _, pair := range pairs {
		if pair.Key() == key {
			return pair.Value()
		}
	}
	return ""
}
func (pairs Pairs) Append(p Pair) Pairs {
	return append(pairs, p)
}
func (pairs Pairs) Count() int {
	if pairs == nil {
		return 0
	}
	return len(pairs)
}
func (pairs Pairs) Equal(another Pairs) bool {
	if pairs.Count() != another.Count() {
		return false
	}
	for i := range pairs {
		if pairs[i].key != another[i].key ||
			pairs[i].val != another[i].val {
			return false
		}
	}
	return true
}
func (pairs Pairs) Clone() Pairs {
	clone := make(Pairs, len(pairs))
	copy(clone, pairs)
	return clone
}
func (pairs Pairs) Prepend(p Pair) Pairs {
	return append(Pairs{p}, pairs...)
}
