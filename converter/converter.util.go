package converter

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func AsPointer[T bool | int32 | int64 | float32 | float64 | string](scalar T) *T {
	pointer := new(T)
	*pointer = scalar
	return pointer
}

func AsScalar[T ~bool | ~int32 | ~int64 | ~float32 | ~float64 | ~string](pointer *T) T {
	if pointer == nil {
		var zero T
		return zero
	} else {
		return *pointer
	}
}

func StringToInt32(s string, defaultValue int32) int32 {
	n, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		return defaultValue
	} else {
		return int32(n)
	}
}

func StringToInt64(s string, defaultValue int64) int64 {
	n, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return defaultValue
	} else {
		return n
	}
}

func StringToBool(s string, defaultValue bool) bool {
	n, err := strconv.ParseBool(s)

	if err != nil {
		return defaultValue
	} else {
		return n
	}
}

func StringToDouble(s string, defaultValue float64) float64 {
	n, err := strconv.ParseFloat(s, 64)

	if err != nil {
		return defaultValue
	} else {
		return n
	}
}

func StringToCommonSeparatedList(s string) []string {
	if s == "" {
		return []string{}
	}

	return strings.Split(s, ",")
}

func Cast(from, to interface{}) error {
	if from == nil {
		return errors.New("nil value was given. cannot cast")
	}

	switch from.(type) {
	case string:
		return json.Unmarshal([]byte(from.(string)), to)
	default:
		if js, err := json.Marshal(from); err != nil {
			return err
		} else {
			return json.Unmarshal(js, to)
		}
	}
}
