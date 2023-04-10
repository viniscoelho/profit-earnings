package types

import "strconv"

type RoundedFloat64 float64

// RoundedFloat64 is a modified version of a float64 which,
// whenever it gets marshaled, it rounds a float64 into
// 2 point precision to facilitate the standard output
func (r RoundedFloat64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(r), 'f', 2, 64)), nil
}
