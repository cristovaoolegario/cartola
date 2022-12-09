package event

import (
	"context"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/usecase"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
)

type ProcessNewPlayer struct{}

func (p ProcessNewPlayer) Process(ctx context.Context, msg *kafka.Message, uow uow.UowInterface) error {
	var input usecase.AddPlayerInput
	err := json.Unmarshal(msg.Value, &input)
	if err != nil {
		return err
	}
	uc := usecase.NewAddPlayerUseCase(uow)

	err = uc.Execute(ctx, input)

	if err != nil {
		return err
	}
	return nil
}
