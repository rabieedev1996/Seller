package InjectionConfig

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
