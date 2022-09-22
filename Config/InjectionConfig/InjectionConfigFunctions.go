package InjectionConfig

import "reflect"

type InjectionList []reflect.Type

var injectionTypes InjectionTypes = InjectionTypes{
	Scope:     InjectionList{},
	Singleton: InjectionList{},
	Transient: InjectionList{},
}

var SingletonObjects = []InjectedObject{}
var ScopeObjects = []ScopeInjectedObject{}

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

func PrepareObject[resultType any](datatype reflect.Type, guid string) resultType {
	if injectionTypes.Scope.SearchInjectionType(datatype) != nil {
		object := GetScopeObject(datatype, guid)
		return object
	} else if injectionTypes.Singleton.SearchInjectionType(datatype) != nil {
		object := GetTransiantObject(datatype)
		return object
	} else if injectionTypes.Singleton.SearchInjectionType(datatype) != nil {
		object := GetSinglletonObject(datatype)
		return object
	}
	return nil
}

func GetScopeObject[resultType any](datatype reflect.Type, guid string) any {
	for _, Scope := range ScopeObjects {
		if Scope.Guid == guid {
			for _, Object := range Scope.Objects {
				if Object.ObjectType == datatype {
					return Object
				}
			}
			object := reflect.New(datatype)
			reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{})
			Scope.Objects = append(Scope.Objects, InjectedObject{
				ObjectType: datatype,
				Object:     object,
			})
			return object
		}
	}
	object := reflect.New(datatype)
	reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{})
	ScopeObjects = append(ScopeObjects, ScopeInjectedObject{
		Guid: guid,
		Objects: []InjectedObject{
			{
				ObjectType: datatype,
				Object:     object,
			},
		},
	})
	return object
}
func GetTransiantObject[resultType any](datatype reflect.Type) any {
	object := reflect.New(datatype)
	reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{})
	return object
}
func GetSinglletonObject[resultType any](datatype reflect.Type) any {
	for _, item := range SingletonObjects {
		if item.ObjectType == datatype {
			return item.Object
		}
	}
	object := reflect.New(datatype)
	reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{})
	return object
}
