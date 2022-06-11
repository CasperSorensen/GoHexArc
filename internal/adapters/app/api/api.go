package api

import (
	"GoHexArc/internal/ports"
)

// the Adapter struct concains a arith
// which is a Arithmetic Interface
type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

// kinda like a constructor so we pass in a
// arithmetic interface, return a reference to a new
// adapter with the injected interface (arith: arith)
func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apiadapter Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiadapter.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apiadapter.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, err
}

func (apiadapter Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiadapter.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apiadapter.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, err
}

func (apiadapter Adapter) GetDivivion(a, b int32) (int32, error) {
	answer, err := apiadapter.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apiadapter.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, err
}

func (apiadapter Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apiadapter.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apiadapter.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, err
}
