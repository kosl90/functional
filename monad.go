package functional

import (
	"fmt"
	"reflect"
)

// MonadInterface is the interface declaration for Monad, which is used on Monad implementation.
type MonadInterface interface {
	// Bind pipes the value of a Monad to a function which returns another Monad.
	// the returned Monad should be the same Monad type to which Bind belongs.
	// using CheckMonadBind to make a runtime checking.
	Bind(func(interface{}) Monad) Monad
}

// Monad used on Monad's API is a wrapper of MonadInterface for convenient.
type Monad interface {
	// Data is used to get value for convenient.
	Data
	MonadInterface
}

// checkBindMonad will check the Monad type of argType and returnType.
// If the types are different, panics, otherwise do nothing.
func checkBindMonad(argType reflect.Type, returnType reflect.Type) {
	if !returnType.Implements(argType) {
		panic(fmt.Sprintf("the return value should be the same Monad as invoker: expected: %v, actually: %v", argType, returnType))
	}
}

// CheckMonadBind will check whether two Monads are the same Monad Type(the different value types are ok),
// which used at Bind implementation.
// @valueInterface is the interface used to implement Monad, like Maybe interface.
// NB: Be careful, the correct return type won't be gotten if the function is not executed(just Monad is recognized).
func CheckMonadBind(valueInterface interface{}, returnValue Monad) Monad {
	checkBindMonad(reflect.TypeOf(valueInterface).Elem(), reflect.TypeOf(returnValue))
	return returnValue
}

// https://wiki.haskell.org/Monad_laws is a little help for understanding monad laws.

// CheckMonadLawLeftIdentity checks the left law of monad.a.k.a. return x >>= f == f x
// @newMonad equals to the return in Haskell, because return cannot be implemented in golang.
// because the @fn might not return the "right" monad, @equals accepts Monad as parameters
func CheckMonadLawLeftIdentity(v interface{}, newMonad func(interface{}) Monad, fn func(interface{}) Monad, equals func(Monad, Monad) bool) bool {
	return equals(newMonad(v).Bind(fn), fn(v))
}

// CheckMonadLawRightIdentity checks the right identity of monad. a.k.a. m >>= return == m
// @newMonad equals to the return in Haskell, because return cannot be implemented in golang.
func CheckMonadLawRightIdentity(m Monad, newMonad func(interface{}) Monad, equals func(Monad, Monad) bool) bool {
	return equals(m, m.Bind(newMonad))
}

// CheckMonadLawAssociativity checks the associativity law of monad. a.k.a. m >>= f1 >>= f2 == m >>= (\x -> f1 x >>= f2)
func CheckMonadLawAssociativity(m Monad, f1 func(interface{}) Monad, f2 func(interface{}) Monad, equals func(Monad, Monad) bool) bool {
	return equals(m.Bind(f1).Bind(f2), m.Bind(func(v interface{}) Monad {
		return f1(v).Bind(f2)
	}))
}
