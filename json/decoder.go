package json

import (
	"encoding/json"
	"io"
)

type custom struct {
	Field string `json:"field"`
}

func decodeWithDecoder(r io.Reader) (custom, error) {
	dec := json.NewDecoder(r)
	c := custom{}
	err := dec.Decode(&c)
	if err != nil {
		return c, err
	}

	return c, nil
}
