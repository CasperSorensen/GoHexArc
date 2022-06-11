package api

import (
	"GoHexArc/internal/ports"
)

// The Adapter struct contains a arith and a db
// which is a Arithmetic Interface, and a Db interface
type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

// NewAdapter is kinda like a constructor, so we pass in the Arithmetic interface, and the Db interface
// then we return a reference to a new adapter reference with the injected interfaces
// db ports.DbPort, arith ports.ArithmeticPort
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
