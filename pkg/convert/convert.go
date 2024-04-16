package convert

import "strconv"

func ConvertStringUint(strValue string) uint64 {
	Value, err := strconv.ParseUint(strValue, 10, 64)
	if err != nil {
		return 0
	}
	return Value
}

func ConvertStringInt(strValue string) int32 {
	Value, err := strconv.ParseInt(strValue, 10, 32)
	if err != nil {
		return 0
	}
	return int32(Value)

}
