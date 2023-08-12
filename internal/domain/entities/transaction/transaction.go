package transaction

import "github.com/Tobiska/avito-internship/internal/domain/entities/account"

type Type int

const (
	Withdraw Type = 0
	Add      Type = 1
)

type Transaction struct {
	Id      int //todo uuid
	Account account.Account
	Type    Type
	Amount  int
}
