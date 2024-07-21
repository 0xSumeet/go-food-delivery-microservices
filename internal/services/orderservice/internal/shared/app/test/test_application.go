package test

import (
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/config/environment"
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/fxapp/test"
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/pkg/logger"
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/services/orderservice/internal/shared/app"
	"github.com/mehdihadeli/go-food-delivery-microservices/internal/services/orderservice/internal/shared/configurations/orders"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

type OrdersTestApplication struct {
	*app.OrdersApplication
	tb fxtest.TB
}

func NewOrdersTestApplication(
	tb fxtest.TB,
	providers []interface{},
	decorates []interface{},
	options []fx.Option,
	logger logger.Logger,
	environment environment.Environment,
) *OrdersTestApplication {
	testApp := test.NewTestApplication(
		tb,
		providers,
		decorates,
		options,
		logger,
		environment,
	)

	orderApplication := &app.OrdersApplication{
		OrdersServiceConfigurator: orders.NewOrdersServiceConfigurator(testApp),
	}

	return &OrdersTestApplication{
		OrdersApplication: orderApplication,
		tb:                tb,
	}
}
