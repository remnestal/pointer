package pointer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	var v bool = true // arbitrary
	assert.Equal(t, v, *Bool(v))
}

func TestString(t *testing.T) {
	var v string = "test" // arbitrary
	assert.Equal(t, v, *String(v))
}

func TestTime(t *testing.T) {
	var v time.Time = time.Date(2000, 1, 2, 3, 4, 5, 6, time.UTC) // arbitrary
	assert.Equal(t, v, *Time(v))
}

func TestDuration(t *testing.T) {
	var v time.Duration = time.Second // arbitrary
	assert.Equal(t, v, *Duration(v))
}

func TestInt(t *testing.T) {
	var v int = 1 // arbitrary
	assert.Equal(t, v, *Int(v))
}

func TestInt8(t *testing.T) {
	var v int8 = 2 // arbitrary
	assert.Equal(t, v, *Int8(v))
}

func TestInt16(t *testing.T) {
	var v int16 = 3 // arbitrary
	assert.Equal(t, v, *Int16(v))
}

func TestInt32(t *testing.T) {
	var v int32 = 4 // arbitrary
	assert.Equal(t, v, *Int32(v))
}

func TestInt64(t *testing.T) {
	var v int64 = 5 // arbitrary
	assert.Equal(t, v, *Int64(v))
}

func TestUint(t *testing.T) {
	var v uint = 6 // arbitrary
	assert.Equal(t, v, *Uint(v))
}

func TestUint8(t *testing.T) {
	var v uint8 = 7 // arbitrary
	assert.Equal(t, v, *Uint8(v))
}

func TestUint16(t *testing.T) {
	var v uint16 = 8 // arbitrary
	assert.Equal(t, v, *Uint16(v))
}

func TestUint32(t *testing.T) {
	var v uint32 = 9 // arbitrary
	assert.Equal(t, v, *Uint32(v))
}

func TestUint64(t *testing.T) {
	var v uint64 = 10 // arbitrary
	assert.Equal(t, v, *Uint64(v))
}

func TestByte(t *testing.T) {
	var v byte = 0x01 // arbitrary
	assert.Equal(t, v, *Byte(v))
}

func TestRune(t *testing.T) {
	var v rune = rune(11) // arbitrary
	assert.Equal(t, v, *Rune(v))
}

func TestFloat32(t *testing.T) {
	var v float32 = 0.1 // arbitrary
	assert.Equal(t, v, *Float32(v))
}

func TestFloat64(t *testing.T) {
	var v float64 = 0.2 // arbitrary
	assert.Equal(t, v, *Float64(v))
}

func TestComplex64(t *testing.T) {
	var v complex64 = 1 + 2i // arbitrary
	assert.Equal(t, v, *Complex64(v))
}

func TestComplex128(t *testing.T) {
	var v complex128 = 3 + 4i // arbitrary
	assert.Equal(t, v, *Complex128(v))
}
