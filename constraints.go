// Package rules defines a set of useful constraints to be used
// with type parameters.
package rules

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Real is a constraint that permits any non-complex numeric type.
type Real interface {
	Float | Integer
}

// Complex is a constraint that permits any complex numeric type.
type Complex interface {
	~complex64 | ~complex128
}

// Number is a constraint that permits any numeric type.
type Number interface {
	Complex | Real
}

type Negable interface {
	Signed | Real | Complex
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
type Ordered interface {
	Real | ~string
}

// Raw is a constraint that permits numbers, strings, uintpointers and booleans
type Raw interface {
	Number | ~string | ~bool
}

// Lener is a constraint that permits any type passable
// to the builtin len function
type Lener[K comparable, T any] interface {
	~string | ~[]T | ~map[K]T
}

// aliases
type (
	C   Complex
	F   Float
	I   Integer
	Neg Negable
	Num Number
	R   Real
	S   Signed
	U   Unsigned
)
