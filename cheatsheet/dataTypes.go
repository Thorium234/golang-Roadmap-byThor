// Boolean
var b bool = true

// String
var s string = "Hello"

// Integers
var i int = 42
var i8 int8 = -128
var i16 int16 = -32768
var i32 int32 = -2147483648
var i64 int64 = -9223372036854775808

// Unsigned Integers
var u uint = 42
var u8 uint8 = 255       // alias: byte
var u16 uint16 = 65535
var u32 uint32 = 4294967295
var u64 uint64 = 18446744073709551615
var up uintptr           // pointer integer type

// Floating Point
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

// Complex Numbers
var c64 complex64 = complex(1, 2)    // float32 real+imag
var c128 complex128 = complex(5, -7) // float64 real+imag

// Rune (alias for int32, Unicode code point)
var r rune = '世'

// Byte (alias for uint8)
var by byte = 'A'
