package InjectionConfig

import (
	"Seller/Seller.Application/Contract/Presistence"
	"Seller/Seller.Domain/Entities"
	"Seller/Seller.Infrastructure/DBRepositories"
	"reflect"
)

type InjectionList []InjectionTypeItem

var injectionTypes InjectionTypes = InjectionTypes{
	Scope:     InjectionList{},
	Singleton: InjectionList{},
	Transient: InjectionList{},
}

var SingletonObjects = []InjectedObject{}
var ScopeObjects = []ScopeInjectedObject{}

func InjectionConfig() {

	var genericReppository Presistence.IGenericRepository[Entities.User]
	var userRepository Presistence.IUserRepository
	AddToInjection(reflect.TypeOf(genericReppository), reflect.TypeOf(DBRepositories.GenericRepository[Entities.User]{}), TRANSIANT)
	AddToInjection(reflect.TypeOf(userRepository), reflect.TypeOf(DBRepositories.UserRepository{}), TRANSIANT)
}

func AddToInjection(interfaceType reflect.Type, objectType reflect.Type, injectType int) {
	if injectType == SINGLETON {
		injectionTypes.Singleton = append(injectionTypes.Singleton, InjectionTypeItem{
			InterfaceType: interfaceType,
			ObjectType:    objectType,
		})
	} else if injectType == SCOPE {
		injectionTypes.Scope = append(injectionTypes.Scope, InjectionTypeItem{
			InterfaceType: interfaceType,
			ObjectType:    objectType,
		})
	} else if injectType == TRANSIANT {
		injectionTypes.Transient = append(injectionTypes.Transient, InjectionTypeItem{
			InterfaceType: interfaceType,
			ObjectType:    objectType,
		})
	}
}

func (list InjectionList) SearchInjectionType(datatype reflect.Type) *InjectionTypeItem {
	for _, dt := range list {
		if dt.InterfaceType == datatype {
			return &dt
		}
	}
	return nil
}

func PrepareObject[resultType any](datatype reflect.Type, guid string) *resultType {

	injectionType := injectionTypes.Scope.SearchInjectionType(datatype)
	if injectionType != nil {
		object := GetScopeObject[resultType](injectionType.InterfaceType, injectionType.ObjectType, guid)
		return &object
	}

	injectionType = injectionTypes.Transient.SearchInjectionType(datatype)
	if injectionType != nil {
		object := GetTransiantObject[resultType](injectionType.InterfaceType, injectionType.ObjectType)
		return &object
	}

	injectionType = injectionTypes.Singleton.SearchInjectionType(datatype)
	if injectionType != nil {
		object := GetSinglletonObject[resultType](injectionType.InterfaceType, injectionType.ObjectType)
		return &object
	}
	return nil
}

func GetScopeObject[resultType any](interfaceType reflect.Type, objectType reflect.Type, guid string) resultType {
	for _, Scope := range ScopeObjects {
		if Scope.Guid == guid {
			for _, Object := range Scope.Objects {
				if Object.IntefaceType == interfaceType {
					return reflect.ValueOf(Object.Object).Interface().(resultType)
				}
			}
			object := reflect.New(objectType)
			reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{reflect.ValueOf(guid)})
			Scope.Objects = append(Scope.Objects, InjectedObject{
				IntefaceType: interfaceType,
				Object:       object,
			})
			return object.Interface().(resultType)
		}
	}
	object := reflect.New(objectType)
	reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{reflect.ValueOf(guid)})
	ScopeObjects = append(ScopeObjects, ScopeInjectedObject{
		Guid: guid,
		Objects: []InjectedObject{
			{
				IntefaceType: interfaceType,
				Object:       object,
			},
		},
	})
	return object.Interface().(resultType)
}
func GetTransiantObject[resultType any](interfaceType reflect.Type, objectType reflect.Type) resultType {
	object := reflect.New(objectType)
	reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{reflect.ValueOf("")})
	return object.Interface().(resultType)
}
func GetSinglletonObject[resultType any](interfaceType reflect.Type, objectType reflect.Type) resultType {
	for _, item := range SingletonObjects {
		if item.IntefaceType == interfaceType {
			return reflect.ValueOf(item.Object).Interface().(resultType)
		}
	}
	object := reflect.New(objectType)
	reflect.ValueOf(&object).MethodByName("Init").Call([]reflect.Value{reflect.ValueOf("")})
	return object.Interface().(resultType)
}
