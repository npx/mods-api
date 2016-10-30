package modsapi

import (
	"math/rand"
	"net/http"
)

// DiceModule is the implementation of a Dice
type DiceModule struct{}

// ID of the Dice Modul, which will be used in the URL
func (module DiceModule) ID() string {
	return "dice"
}

// Endpoints lists all endpoints of the module
func (module DiceModule) Endpoints() []Endpoint {
	return []Endpoint{
		Endpoint{"/roll", module.rollDie},
	}
}

// Response of the Dice Roll
type diceResponse struct {
	Result int `json:"result"`
}

// The RollDice function
<<<<<<< 0bebf8b3477d83dd97b3c2c79976c769df60bee0
func (module DiceModule) rollDie(r *http.Request) JSONResponse {
=======
func (module DiceModule) rollDie(r *http.Request) JSONResponse {
>>>>>>> changed the modules to support JSONRequest
	value := rand.Intn(6) + 1

	return JSONResponse{
		http.StatusOK,
		diceResponse{value},
	}
}
