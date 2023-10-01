package json

import (
	"encoding/json"
	"io"
)

func unmarshallWithSize(r io.Reader, size int) (custom, error) {
	var (
		c = custom{}
		b = make([]byte, size)
	)

	_, err := r.Read(b)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(b, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
