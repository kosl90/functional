package functional

type FunctorInterface interface {
	Map(func(interface{}) interface{}) Functor
}

type Functor interface {
	Data
	FunctorInterface
}
