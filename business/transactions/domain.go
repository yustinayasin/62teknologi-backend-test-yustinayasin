package transactions

type Transaction struct {
	Id   int
	Name string
}

type TransactionUseCaseInterface interface {
	Add(transaction Transaction) (Transaction, error)
}

type TransactionRepoInterface interface {
	Add(transaction Transaction) (Transaction, error)
}
