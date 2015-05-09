package functional_test

import (
	. "github.com/kosl90/functional"
	. "github.com/onsi/gomega"
	"testing"
)

func TestFPMaybeMonad(t *testing.T) {
	RegisterTestingT(t)

	Expect(func() {
		Just(1).Bind(func(interface{}) Monad {
			return Left(1)
		})
	}).To(Panic())

	v := Just(1).Bind(func(a interface{}) Monad {
		return Just(a.(int) + 1)
	}).Value()
	Expect(v).To(Equal(2))

	nv := None().Bind(func(a interface{}) Monad {
		return Just(1)
	}).Bind(func(a interface{}) Monad {
		return Just(a.(int) + 10)
	}).(Maybe)
	Expect(nv.IsNone()).To(BeTrue())
}

func TestFPMaybeString(t *testing.T) {
	if None().String() != "None" {
		t.Fail()
	}

	if Just(1).String() != "Just 1" {
		t.Fail()
	}
}

func TestFPMaybeFunctor(t *testing.T) {
	RegisterTestingT(t)

	v := Just(1).Map(func(a interface{}) interface{} {
		return a.(int) + 1
	}).Value()
	Expect(v).To(Equal(2))

	nv := None().Map(func(interface{}) interface{} {
		return Just(1)
	}).(Maybe)
	Expect(nv.IsNone()).To(BeTrue())
}

func newMaybeMonad(v interface{}) Monad {
	return Just(v)
}

func TestFPMaybeMonadLaws(t *testing.T) {
	RegisterTestingT(t)

	Ω(CheckMonadLawLeftIdentity(1, newMaybeMonad, func(v interface{}) Monad { return Just(v.(int) + 1) }, func(m1 Monad, m2 Monad) bool {
		return m1.Value() == m2.Value()
	})).To(BeTrue())

	Expect(CheckMonadLawLeftIdentity(1, newMaybeMonad, func(interface{}) Monad { return None() }, func(m1 Monad, m2 Monad) bool {
		return m1.(Maybe).IsNone() && m2.(Maybe).IsNone()
	})).To(BeTrue())

	Expect(CheckMonadLawRightIdentity(Just(1), newMaybeMonad, func(m1 Monad, m2 Monad) bool {
		return m1.(Maybe).Value() == m2.(Maybe).Value()
	})).To(BeTrue())

	Expect(CheckMonadLawAssociativity(Just(1), func(v interface{}) Monad {
		return Just(v)
	}, func(v interface{}) Monad {
		return Just(v.(int) * 2)
	}, func(m1 Monad, m2 Monad) bool {
		return m1.(Maybe).Value() == m2.(Maybe).Value()
	})).To(BeTrue())

	Expect(CheckMonadLawAssociativity(Just(1), func(v interface{}) Monad {
		return Just(v)
	}, func(v interface{}) Monad {
		return None()
	}, func(m1 Monad, m2 Monad) bool {
		return m1.(Maybe).IsNone() && m2.(Maybe).IsNone()
	})).To(BeTrue())
}
