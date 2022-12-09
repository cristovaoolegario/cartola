package event

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
)

func TestNewPlayer(t *testing.T) {
	ctx, db := repository.SetupTestDb("../../../../")
	uow := repository.SetupTestUoW(ctx, db)

	msg := kafka.Message{
		Value: []byte(`{"id": "10","name": "TestNewPlayer","initial_price": 10.5}`),
	}

	malformedMsg := kafka.Message{
		Value: []byte(`{ice": 10.5}`),
	}

	p := ProcessNewPlayer{}

	t.Run("Should return an error When message json is malformed", func(t *testing.T) {
		err := p.Process(ctx, &malformedMsg, uow)

		if err == nil {
			t.Error("Expected an unmarshall error to show, got any")
		}
	})

	t.Run("Shouldn't return an error When message json OK", func(t *testing.T) {
		err := p.Process(ctx, &msg, uow)

		if err != nil {
			t.Errorf("Expected no error to show, got %s", err.Error())
		}
	})

}
