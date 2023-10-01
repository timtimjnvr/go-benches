package json

import (
	"bytes"
	"encoding/json"
	"io"
)

func decoderWithSize(r io.Reader, size int) (custom, error) {
	b := make([]byte, size)
	_, err := r.Read(b)
	if err != nil {
		return custom{}, err
	}

	var (
		dec = json.NewDecoder(bytes.NewReader(b))
		c   = custom{}
	)

	err = dec.Decode(&c)
	if err != nil {
		return c, err
	}

	return c, nil
}
