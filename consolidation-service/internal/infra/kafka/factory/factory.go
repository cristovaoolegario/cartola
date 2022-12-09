package factory

import "github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/kafka/event"

func CreateProcessMessageStrategy(topic string) event.ProcessEventStrategy {
	switch topic {
	case "newPlayer":
		// {"id": "10","name": "Wesley","initial_price": 10.5}
		return event.ProcessNewPlayer{}
	}
	return nil
}
