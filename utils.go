package txtpack

import "io"

func writeAll(writer io.Writer, str []byte) error {
	written := 0
	for written < len(str) {
		n, err := writer.Write(str[written:])
		written += n
		if err != nil {
			return err
		}
	}
	return nil
}
