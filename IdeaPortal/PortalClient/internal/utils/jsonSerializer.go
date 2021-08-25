package utils

import (
	"encoding/json"
	"io"
)

func ToJSON(w io.Writer, body []byte) error {
	var v interface{}
	err := json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	e := json.NewEncoder(w)
	return e.Encode(v)
}

func FromJSON(r io.Reader, body []byte) error {
	var v interface{}
	err := json.Unmarshal(body, &v)
	if err != nil {
		return err
	}
	e := json.NewDecoder(r)
	return e.Decode(v)
}
