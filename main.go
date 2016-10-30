package modsapi

var modules = []Module{
	DeciderModule{},
	DiceModule{},
	CoinModule{},
}

func init() {
	setupServer(modules)
}
