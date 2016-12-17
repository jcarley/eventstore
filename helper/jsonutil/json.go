package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Encodes/Marshals the given object into JSON
func EncodeJSON(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := EncodeJSONToWriter(&buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func EncodeJSONToString(v interface{}) (string, error) {
	data, err := EncodeJSON(v)
	return string(data), err
}

func EncodeJSONToWriter(writer io.Writer, v interface{}) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(v)
}

// Decodes/Unmarshals the given JSON into a desired object
func DecodeJSON(data []byte, out interface{}) error {
	if data == nil {
		return fmt.Errorf("'data' being decoded is nil")
	}
	if out == nil {
		return fmt.Errorf("output parameter 'out' is nil")
	}

	return DecodeJSONFromReader(bytes.NewReader(data), out)
}

// Decodes/Unmarshals the given io.Reader pointing to a JSON, into a desired object
func DecodeJSONFromReader(r io.Reader, out interface{}) error {
	if r == nil {
		return fmt.Errorf("'io.Reader' being decoded is nil")
	}
	if out == nil {
		return fmt.Errorf("output parameter 'out' is nil")
	}

	dec := json.NewDecoder(r)

	// While decoding JSON values, intepret the integer values as `json.Number`s instead of `float64`.
	dec.UseNumber()

	// Since 'out' is an interface representing a pointer, pass it to the decoder without an '&'
	return dec.Decode(out)
}
