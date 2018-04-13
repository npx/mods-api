package main

import (
	"github.com/npx/mods-api/module"
	"github.com/npx/mods-api/server"
)

var modules = []module.Module{
	module.DeciderModule{},
	module.DiceModule{},
	module.CoinModule{},
}

func main() {
	server.Setup(modules)
	server.Start()
}
