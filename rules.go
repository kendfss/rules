// Package rules defines a set of useful constraints to be used
// with type parameters.
package rules

type Word[C Char] interface {
	~[]C
}

// Char is a constraint that permits any character type.
type Char interface {
	// ~int32 | ~uint8 //| ~string
	// ~rune | ~byte //| ~string
	~rune | ~byte
}

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// SingedNumber is a constraint that permits any signed number type
type SignedNumber interface {
	Signed | Float | Complex
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

// OrderedNumber is a constraint that permits any non-complex numeric type.
type OrderedNumber interface {
	Float | Integer
}

// Complex is a constraint that permits any complex numeric type.
type Complex interface {
	~complex64 | ~complex128
}

// Number is a constraint that permits any numeric type.
type Number interface {
	Complex | OrderedNumber
}

// OrderedNegable is a constraint that permits any signed number-type that is also ordered
type OrderedNegable interface {
	Signed | Float
}

// Negable is a constraint that permits any signed number-type
type Negable interface {
	OrderedNegable | Complex
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
type Ordered interface {
	OrderedNumber | ~string
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

// // Risk is a constraint that permits a value of any type
// // or a pointer to it
// type Risk[T any] interface {
// 	T | *T
// }

// utility aliases
type (
	C     = Complex
	F     = Float
	I     = Integer
	Int   = Integer
	Neg   = Negable
	Num   = Number
	R     = OrderedNumber
	Real  = OrderedNumber
	S     = Signed
	U     = Unsigned
	Uint  = Unsigned
	Adder = Ordered // supports the "+" operator
)
