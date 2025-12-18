package lib

import (
	"encoding/json"
	"io"
)

func ReadBody[T any](body io.Reader) (T, error) {
	var result T
	data, err := io.ReadAll(body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}
