package core

type DataLoader interface {
	Load() error
}

type DataGetter interface {
	Get(k string) map[string]interface{}
}

type DataMap interface {
	DataLoader
	DataGetter
}
