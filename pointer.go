package pointer

import "time"

// Bool returns a pointer to the specified bool value.
func Bool(v bool) *bool {
	return &v
}

// String returns a pointer to the specified string value.
func String(v string) *string {
	return &v
}

// Time returns a pointer to the specified time value.
func Time(v time.Time) *time.Time {
	return &v
}

// Duration returns a pointer to the specified duration value.
func Duration(v time.Duration) *time.Duration {
	return &v
}

// Int returns a pointer to the specified int value.
func Int(v int) *int {
	return &v
}

// Int8 returns a pointer to the specified int8 value.
func Int8(v int8) *int8 {
	return &v
}

// Int16 returns a pointer to the specified int16 value.
func Int16(v int16) *int16 {
	return &v
}

// Int32 returns a pointer to the specified int32 value.
func Int32(v int32) *int32 {
	return &v
}

// Int64 returns a pointer to the specified int64 value.
func Int64(v int64) *int64 {
	return &v
}

// Uint returns a pointer to the specified uint value.
func Uint(v uint) *uint {
	return &v
}

// Uint8 returns a pointer to the specified uint8 value.
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint16 returns a pointer to the specified uint16 value.
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint32 returns a pointer to the specified uint32 value.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint64 returns a pointer to the specified uint64 value.
func Uint64(v uint64) *uint64 {
	return &v
}

// Byte returns a pointer to the specified byte value.
func Byte(v byte) *byte {
	return &v
}

// Rune returns a pointer to the specified rune value.
func Rune(v rune) *rune {
	return &v
}

// Float32 returns a pointer to the specified float32 value.
func Float32(v float32) *float32 {
	return &v
}

// Float64 returns a pointer to the specified float64 value.
func Float64(v float64) *float64 {
	return &v
}

// Complex64 returns a pointer to the specified complex64 value.
func Complex64(v complex64) *complex64 {
	return &v
}

// Complex128 returns a pointer to the specified complex128 value.
func Complex128(v complex128) *complex128 {
	return &v
}
