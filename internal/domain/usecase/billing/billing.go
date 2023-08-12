package billing

import "context"

type Usecase struct {
	// todo repos
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) GetAccount(ctx context.Context)        {}
func (u *Usecase) CreateTransaction(ctx context.Context) {}
