package usecase

import (
	"context"
	"time"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/domain/entity"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/domain/repository"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
)

type MatchInput struct {
	ID      string
	Date    time.Time
	TeamAID string
	TeamBID string
}

type MatchUseCase struct {
	Uow uow.UowInterface
}

func NewMatchUseCase(uow uow.UowInterface) *MatchUseCase {
	return &MatchUseCase{
		Uow: uow,
	}
}

func (uc *MatchUseCase) Execute(ctx context.Context, input MatchInput) error {
	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		matchRepository := uc.getMatchRepository(ctx)
		teamRepository := uc.getTeamRepository(ctx)

		teamA, err := teamRepository.FindByID(ctx, input.TeamAID)
		if err != nil {
			return err
		}
		teamB, err := teamRepository.FindByID(ctx, input.TeamBID)
		if err != nil {
			return err
		}

		match := entity.NewMatch(input.ID, teamA, teamB, input.Date)

		// Create match
		err = matchRepository.Create(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (uc *MatchUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := uc.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (uc *MatchUseCase) getTeamRepository(ctx context.Context) repository.TeamRepositoryInterface {
	teamRepository, err := uc.Uow.GetRepository(ctx, "TeamRepository")
	if err != nil {
		panic(err)
	}
	return teamRepository.(repository.TeamRepositoryInterface)
}
