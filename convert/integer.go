package convert

import (
	"fmt"
	"strconv"
)

func MustToInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return x
}

func MustToUint32(str string) uint32 {
	v, _ := strconv.ParseUint(str, 10, 64)
	return uint32(v)
}

func MustToI64(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return x
}

func MustToU64(s string) uint64 {
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return x
}

func ToString(i interface{}) (string, bool) {
	switch iv := i.(type) {
	case int:
		return strconv.Itoa(iv), true
	case int8:
		return strconv.Itoa(int(iv)), true
	case int16:
		return strconv.Itoa(int(iv)), true
	case int32:
		return strconv.Itoa(int(iv)), true
	case int64:
		return strconv.FormatInt(iv, 10), true
	case string:
		return iv, true
	default:
		return "", false
	}
}

func MustToString(i interface{}) string {
	switch iv := i.(type) {
	case int:
		return strconv.Itoa(iv)
	case int8:
		return strconv.Itoa(int(iv))
	case int16:
		return strconv.Itoa(int(iv))
	case int32:
		return strconv.Itoa(int(iv))
	case int64:
		return strconv.FormatInt(iv, 10)
	case string:
		return iv
	default:
		return fmt.Sprintf("%+v", i)
	}
}
