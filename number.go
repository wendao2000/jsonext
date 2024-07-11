package jsonext

import (
	"encoding/json"
	"strconv"
)

type Number float64

func (n Number) IsWhole() bool {
	return float64(n) == float64(int64(n))
}

func (n Number) String() string {
	if n.IsWhole() {
		return strconv.FormatInt(int64(n), 10)
	}

	return strconv.FormatFloat(float64(n), 'f', -1, 64)
}

func (n Number) Float64() float64 {
	return float64(n)
}

func (n Number) Int64() int64 {
	return int64(n)
}

func (n Number) MarshalJSON() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *Number) UnmarshalJSON(b []byte) error {
	var s Stringable
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	num, err := s.ToNumber()
	if err != nil {
		return err
	}

	*n = num
	return nil
}
