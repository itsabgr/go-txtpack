package txtpack

import (
	"bytes"
	"io"
)

func Unmarshal(src interface {
	ReadBytes(delim byte) ([]byte, error)
}) (map[string][]byte, error) {
	pairs := make(map[string][]byte)
	for {
		line, err := src.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return pairs, nil
			}
			return pairs, err
		}
		if line == nil || len(line) == 0 {
			continue
		}
		if bytes.Equal(line, []byte("\n")) {
			return pairs, nil
		}
		deli := bytes.IndexByte(line, ':')
		if deli == -1 {
			pairs[string(line)] = nil
			continue
		}
		pairs[string(line[:deli])] = line[deli+2 : len(line)-1]
	}
}
func Marshal(buff io.Writer, kvs map[string][]byte) error {
	for k, v := range kvs {
		err := writeAll(buff, []byte(k))
		if err != nil {
			return err
		}
		err = writeAll(buff, []byte(": "))
		if err != nil {
			return err
		}
		err = writeAll(buff, v)
		if err != nil {
			return err
		}
		err = writeAll(buff, []byte("\n"))
		if err != nil {
			return err
		}
	}
	buff.Write([]byte("\n"))
	return nil
}
