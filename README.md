rules
---

Another tedious joy of third world living.
To be specific, I need a copy of the [exp/constraints][original] module, but little else of `exp`'s contents. Because I don't have reliable internet, and go forces you to manage dependencies online/adopt a language-specific environment to benefit from it's new `work` feature, it becomes necessary to hold on to extra copies of stuff. 
That said, it now contains a few other primitive constraints


```go

package rules // import "github.com/kendfss/rules"

Package rules defines a set of useful constraints to be used with type
parameters.

TYPES

type Adder = Ordered // supports the "+" operator
    // utility aliases

type C = Complex
    // utility aliases

type Char interface {
	~rune | ~byte
}
    // Char is a constraint that permits any character type.

type Complex interface {
	~complex64 | ~complex128
}
    // Complex is a constraint that permits any complex numeric type.

type F = Float
    // utility aliases

type Float interface {
	~float32 | ~float64
}
    // Float is a constraint that permits any floating-point type.

type I = Integer
    // utility aliases

type Int = Integer
    // utility aliases

type Integer interface {
	Signed | Unsigned
}
    // Integer is a constraint that permits any integer type.

type Lener[K comparable, T any] interface {
	~string | ~[]T | ~map[K]T
}
    // Lener is a constraint that permits any type passable to the builtin len
    // function

type Neg = Negable
    // utility aliases

type Negable interface {
	OrderedNegable | Complex
}
    // Negable is a constraint that permits any signed number-type

type Num = Number
    // utility aliases

type Number interface {
	Complex | OrderedNumber
}
    // Number is a constraint that permits any numeric type.

type Ordered interface {
	OrderedNumber | ~string
}
    // Ordered is a constraint that permits any ordered type: any type that
    // supports the operators < <= >= >.

type OrderedNegable interface {
	Signed | Float
}
    // OrderedNegable is a constraint that permits any signed number-type that is
    // also ordered

type OrderedNumber interface {
	Float | Integer
}
    // OrderedNumber is a constraint that permits any non-complex numeric type.

type R = OrderedNumber
    // utility aliases

type Raw interface {
	Number | ~string | ~bool
}
    // Raw is a constraint that permits numbers, strings, uintpointers and booleans

type Real = OrderedNumber
    // utility aliases

type S = Signed
    // utility aliases

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
    // Signed is a constraint that permits any signed integer type.

type SignedNumber interface {
	Signed | Float | Complex
}
    // SingedNumber is a constraint that permits any signed number type

type U = Unsigned
    // utility aliases

type Uint = Unsigned
    // utility aliases

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
    // Unsigned is a constraint that permits any unsigned integer type.

type Word[C Char] interface {
	~[]C
}

```



[original]: https://github.com/golang/exp/tree/master/constraints
