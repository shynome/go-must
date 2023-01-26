package must

import "github.com/lainio/err2/try"

func To[F func() error, R func()](f F) R {
	return func() { try.To(f()) }
}
func To1[U any, F func() (U, error), R func() U](f F) R {
	return func() U { return try.To1(f()) }
}
func To2[U, V any, F func() (U, V, error), R func() (U, V)](f F) R {
	return func() (U, V) { return try.To2(f()) }
}
func To3[U, V, W any, F func() (U, V, W, error), R func() (U, V, W)](f F) R {
	return func() (U, V, W) { return try.To3(f()) }
}

func Tp[O any, F func(O) error, R func(O)](f F) R {
	return func(o O) { try.To(f(o)) }
}
func Tp1[O any, U any, F func(O) (U, error), R func(O) U](f F) R {
	return func(o O) U { return try.To1(f(o)) }
}
func Tp2[O any, U, V any, F func(O) (U, V, error), R func(O) (U, V)](f F) R {
	return func(o O) (U, V) { return try.To2(f(o)) }
}
func Tp3[O any, U, V, W any, F func(O) (U, V, W, error), R func(O) (U, V, W)](f F) R {
	return func(o O) (U, V, W) { return try.To3(f(o)) }
}

func Tpp[O, P any, F func(O, P) error, R func(O, P)](f F) R {
	return func(o O, p P) { try.To(f(o, p)) }
}
func Tpp1[O, P any, U any, F func(O, P) (U, error), R func(O, P) U](f F) R {
	return func(o O, p P) U { return try.To1(f(o, p)) }
}
func Tpp2[O, P any, U, V any, F func(O, P) (U, V, error), R func(O, P) (U, V)](f F) R {
	return func(o O, p P) (U, V) { return try.To2(f(o, p)) }
}
func Tpp3[O, P any, U, V, W any, F func(O, P) (U, V, W, error), R func(O, P) (U, V, W)](f F) R {
	return func(o O, p P) (U, V, W) { return try.To3(f(o, p)) }
}

func Tppp[O, P, Q any, F func(O, P, Q) error, R func(O, P, Q)](f F) R {
	return func(o O, p P, q Q) { try.To(f(o, p, q)) }
}
func Tppp1[O, P, Q any, U any, F func(O, P, Q) (U, error), R func(O, P, Q) U](f F) R {
	return func(o O, p P, q Q) U { return try.To1(f(o, p, q)) }
}
func Tppp2[O, P, Q any, U, V any, F func(O, P, Q) (U, V, error), R func(O, P, Q) (U, V)](f F) R {
	return func(o O, p P, q Q) (U, V) { return try.To2(f(o, p, q)) }
}
func Tppp3[O, P, Q any, U, V, W any, F func(O, P, Q) (U, V, W, error), R func(O, P, Q) (U, V, W)](f F) R {
	return func(o O, p P, q Q) (U, V, W) { return try.To3(f(o, p, q)) }
}
