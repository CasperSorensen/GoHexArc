package ports

type APIPort interface {
	GetAddition(a, b int32) (int32, error)
	GetMultiplication(a, b int32) (int32, error)
	GetDivivion(a, b int32) (int32, error)
	GetSubtraction(a, b int32) (int32, error)
}
