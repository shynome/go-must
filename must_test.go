package must_test

import (
	"testing"

	"github.com/shynome/go-must"
)

func TestMust(t *testing.T) {
	f00 := func() (err error) { return }
	f10 := func(a string) (err error) { return }
	f11 := func(a string) (x int, err error) { return }
	f12 := func(a string) (x, y int, err error) { return }
	f13 := func(a string) (x, y, z int, err error) { return }
	f20 := func(a, b string) (err error) { return }
	f21 := func(a, b string) (x int, err error) { return }
	f22 := func(a, b string) (x, y int, err error) { return }
	f23 := func(a, b string) (x, y, z int, err error) { return }
	f30 := func(a, b, c string) (err error) { return }
	f31 := func(a, b, c string) (x int, err error) { return }
	f32 := func(a, b, c string) (x, y int, err error) { return }
	f33 := func(a, b, c string) (x, y, z int, err error) { return }

	must.To(f00)

	must.Tp(f10)
	must.Tp1(f11)
	must.Tp2(f12)
	must.Tp3(f13)

	must.Tpp(f20)
	must.Tpp1(f21)
	must.Tpp2(f22)
	must.Tpp3(f23)

	must.Tppp(f30)
	must.Tppp1(f31)
	must.Tppp2(f32)
	must.Tppp3(f33)
}
