package report

import "context"

type Usecase struct {
	// todo repos
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) GetReport(ctx context.Context) {}
