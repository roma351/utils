package utils

import (
	"fmt"
	"strconv"
)

func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func StringToIntNoErr(value string) int {
	n, _ := strconv.Atoi(value)
	return n
}

func StringToInt64(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 64)
}

func StringToInt64NoErr(value string) int64 {
	i64, _ := strconv.ParseInt(value, 10, 64)
	return i64
}

func StringToUint32(value string) (uint32, error) {
	ui64, err := strconv.ParseUint(value, 10, 32)
	return uint32(ui64), err
}

func StringToUint32NoErr(value string) uint32 {
	ui32, _ := strconv.ParseUint(value, 10, 32)
	return uint32(ui32)
}

func StringToUint64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

func StringToUint64NoErr(value string) uint64 {
	ui64, _ := strconv.ParseUint(value, 10, 64)
	return ui64
}

func IntToString(value int) string {
	return fmt.Sprintf("%d", value)
}

func Int32ToString(value int32) string {
	return fmt.Sprintf("%d", value)
}

func Uint32ToString(value uint32) string {
	return fmt.Sprintf("%d", value)
}

func Int64ToString(value int64) string {
	return fmt.Sprintf("%d", value)
}

func Uint64ToString(value uint64) string {
	return fmt.Sprintf("%d", value)
}

func StringToFloat32(value string) (float32, error) {
	f32, err := strconv.ParseFloat(value, 2)
	return float32(f32), err
}

func StringToFloat64(value string) (float64, error) {
	return strconv.ParseFloat(value, 2)
}

func StringToFloat64NoErr(value string) float64 {
	f64, _ := strconv.ParseFloat(value, 2)
	return f64
}

func StringToInt32NoErr(value string) int32 {
	ui32, _ := strconv.ParseInt(value, 10, 32)
	return int32(ui32)
}

func StringToInt32(value string) (int32, error) {
	ui32, err := strconv.ParseInt(value, 10, 32)
	return int32(ui32), err
}
