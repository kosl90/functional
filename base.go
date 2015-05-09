package functional

type UnitT struct{}

func Unit() *UnitT {
	return new(UnitT)
}

type Data interface {
	Value() interface{}
}
