package usecase

import (
	"context"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/domain/entity"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/domain/repository"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
)

type AddPlayerInput struct {
	ID           string
	Name         string
	InitialPrice float64
}

type AddPlayerUseCase struct {
	Uow uow.UowInterface
}

func NewAddPlayerUseCase(uow uow.UowInterface) *AddPlayerUseCase {
	return &AddPlayerUseCase{
		Uow: uow,
	}
}

func (uc *AddPlayerUseCase) Execute(ctx context.Context, input AddPlayerInput) error {
	playerRepository := uc.getPlayerRepository(ctx)
	player := entity.NewPlayer(input.ID, input.Name, input.InitialPrice)

	err := playerRepository.Create(ctx, player)

	if err != nil {
		return err
	}

	return uc.Uow.CommitOrRollback()
}

func (uc *AddPlayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := uc.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
