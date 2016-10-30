package modsapi

var modules = []Module{
	DeciderModule{},
	DiceModule{},
}

func init() {
	setupServer(modules)
}
