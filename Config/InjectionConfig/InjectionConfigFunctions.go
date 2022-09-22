package InjectionConfig

import "reflect"

type InjectionList []reflect.Type

var injectionTypes InjectionTypes = InjectionTypes{
	Scope:     InjectionList{},
	Singleton: InjectionList{},
	Transient: InjectionList{},
}

func AddToInjection(datatype reflect.Type, injectType int) {
	if injectType == SINGLETON {
		injectionTypes.Singleton = append(injectionTypes.Singleton, datatype)
	} else if injectType == SCOPE {
		injectionTypes.Scope = append(injectionTypes.Scope, datatype)
	} else if injectType == TRANSIANT {
		injectionTypes.Transient = append(injectionTypes.Transient, datatype)
	}
}

func (list InjectionList) SearchInjectionType(datatype reflect.Type) reflect.Type {
	for _, dt := range list {
		if dt == datatype {
			return dt
		}
	}
	return nil
}

func PrepareObject[resultType any](datatype reflect.Type) resultType {
	if injectionTypes.Scope.SearchInjectionType(datatype) != nil {
		object := GetScopeObject(datatype)
		reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{})
		return object
	} else if injectionTypes.Singleton.SearchInjectionType(datatype) != nil {
		object := GetTransiantObject(datatype)
		reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{})
		return object
	} else if injectionTypes.Singleton.SearchInjectionType(datatype) != nil {
		object := GetSinglletonObject(datatype)
		reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{})
		return object
	}
	return nil
}

func GetScopeObject[resultType any](datatype reflect.Type) any {
	return reflect.New(datatype)
}
func GetTransiantObject[resultType any](datatype reflect.Type) any {
	return reflect.New(datatype)
}
func GetSinglletonObject[resultType any](datatype reflect.Type) any {
	return reflect.New(datatype)
}
