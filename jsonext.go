package jsonext

import "encoding/json"

// MarshalString returns the string equivalent of the JSON encoding of v.
func MarshalString(v any) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

// MustMarhalString acts like MarshalString, but will panic if json.Marshal
// returns error.
func MustMarshalString(v any) string {
	if b, err := json.Marshal(v); err != nil {
		panic(err)
	} else {
		return string(b)
	}
}

// Sptring sprint string, without returning any error.
func Sptring(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
