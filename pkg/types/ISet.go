package types

type ISet interface {
	Add(interface{})
	IsPresent(interface{}) bool
	ToArray() []interface{}
}
