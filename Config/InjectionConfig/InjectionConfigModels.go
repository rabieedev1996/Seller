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

type InjectedObject struct {
	Object     any
	ObjectType reflect.Type
}

type ScopeInjectedObject struct {
	Objects []InjectedObject
	Guid    string
}
