package json

import (
	"encoding/json"
	"io"
)

func decodeWithUnmarshall(r io.Reader) (custom, error) {
	c := custom{}
	b, err := io.ReadAll(r)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
