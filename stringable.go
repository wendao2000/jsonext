package jsonext

import "strconv"

type Stringable string

func (s Stringable) String() string {
	return string(s)
}

func (s Stringable) Float64() (float64, error) {
	n, err := s.ToNumber()
	if err != nil {
		return float64(0), err
	}
	return n.Float64(), nil
}

func (s Stringable) Int64() (int64, error) {
	n, err := s.ToNumber()
	if err != nil {
		return int64(0), err
	}
	return n.Int64(), nil
}

func (s Stringable) ToNumber() (Number, error) {
	if len(s) == 0 {
		return Number(0), nil
	}
	val, err := strconv.ParseFloat(s.String(), 64)
	if err != nil {
		return Number(0), ErrNotNumber
	}
	return Number(val), nil
}

func (s Stringable) MustToNumber() Number {
	num, err := s.ToNumber()
	if err != nil {
		panic(err)
	}
	return num
}

func (s Stringable) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.String() + "\""), nil
}

func (s *Stringable) UnmarshalJSON(b []byte) error {
	if b == nil {
		*s = Stringable("")
		return nil
	}

	str := string(b)
	if str[0] == '"' && str[len(str)-1] == '"' {
		*s = Stringable(str[1 : len(str)-1])
		return nil
	}

	*s = Stringable(b)
	return nil
}
