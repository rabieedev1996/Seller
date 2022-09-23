package InjectionConfig

import "reflect"

const (
	TRANSIANT = iota
	SCOPE
	SINGLETON
)

type InjectionTypes struct {
	Singleton InjectionList
	Scope     InjectionList
	Transient InjectionList
}
type InjectionTypeItem struct {
	InterfaceType reflect.Type
	ObjectType    reflect.Type
}
type InjectedObject struct {
	Object       any
	IntefaceType reflect.Type
}

type ScopeInjectedObject struct {
	Objects []InjectedObject
	Guid    string
}
