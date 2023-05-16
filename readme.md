constraints
---

Another tedious joy of third world living.
To be specific, I need a copy of the [exp/constraints][original] module, but little else of `exp`'s contents. Because I don't have reliable internet, and go forces you to manage dependencies online/adopt a language-specific environment to benefit from it's new `work` feature, it becomes necessary to hold on to extra copies of stuff. 
That said, it now contains a few other primitive constraints


```go
type C Complex
type F Float
type I Integer
type R Real
type S Signed
type U Unsigned

type Complex interface
type Float interface
type Integer interface
type Lener[K comparable, T any] interface // types that can be passed to builtin len function
type Negable interface // types that support the "-" operator
type Number interface
type Ordered interface
type Raw interface // Number | ~string | ~bool
type Real interface // Non-complex Number
type Signed interface
type Unsigned interface
```



[original]: https://github.com/golang/exp/tree/master/constraints
