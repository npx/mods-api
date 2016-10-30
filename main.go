package modsapi

var modules = []Module{
	DeciderModule{},
}

func init() {
	setupServer(modules)
}
