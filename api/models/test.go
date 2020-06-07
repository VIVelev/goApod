package models

type Test struct {
	Info string
}

func (T Test) GetInfo() string {
	return T.Info
}
