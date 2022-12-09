package event

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/usecase"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
)

type ProcessMatchUpdateResult struct{}

func (p ProcessMatchUpdateResult) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.MatchUpdateResultInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	fmt.Println("input: ", input)
	updateMatchResultUsecase := usecase.NewMatchUpdateResultUseCase(uow)
	err = updateMatchResultUsecase.Execute(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
