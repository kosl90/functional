package functional

import (
	"fmt"
)

type Maybe interface {
	IsNone() bool
	Data
	fmt.Stringer
	MonadInterface
	FunctorInterface
}

var _MaybeInterface = new(Maybe)

type JustT struct {
	v interface{}
}

func (*JustT) IsNone() bool {
	return false
}

func (m *JustT) Value() interface{} {
	return m.v
}

func (m *JustT) String() string {
	return fmt.Sprintf("Just %v", m.v)
}

func (m *JustT) Bind(fn func(interface{}) Monad) Monad {
	return CheckMonadBind(_MaybeInterface, fn(m.v))
}

func (m *JustT) Map(fn func(interface{}) interface{}) Functor {
	return Just(fn(m.v))
}

type NoneT struct { }

func (*NoneT) IsNone() bool {
	return true
}

func (*NoneT) Value() interface{} {
	panic("None has no value")
}

func (m *NoneT) Bind(func(interface{}) Monad) Monad {
	return m
}

func (m *NoneT) Map(func(interface{}) interface{}) Functor {
	return m
}

func (*NoneT) String() string {
	return "None"
}

func Just(v interface{}) Maybe {
//	return new(JustT, v)
	return &JustT{v}
}

func None() Maybe {
	return new(NoneT)
}
