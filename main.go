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

func init() {
	server.Setup(modules)
}

func main() {
	server.Start()
}
