package integrationEvents

import (
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/types"
	uuid "github.com/satori/go.uuid"

	dtosV1 "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/dtos/v1"
)

type OrderCreatedV1 struct {
	*types.Message
	*dtosV1.OrderReadDto
}

func NewOrderCreatedV1(orderReadDto *dtosV1.OrderReadDto) *OrderCreatedV1 {
	return &OrderCreatedV1{OrderReadDto: orderReadDto, Message: types.NewMessage(uuid.NewV4().String())}
}