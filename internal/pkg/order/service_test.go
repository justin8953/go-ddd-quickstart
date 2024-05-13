package order

import (
	"go-ddd-quickstart/internal/pkg/db"
	dbMock "go-ddd-quickstart/internal/pkg/db/mock"
	"go-ddd-quickstart/internal/pkg/dto"
	"go-ddd-quickstart/internal/pkg/events"
	handlerMock "go-ddd-quickstart/internal/pkg/events/mock"
	orderEvents "go-ddd-quickstart/internal/pkg/events/order"
	dbRecord "go-ddd-quickstart/internal/pkg/order/db"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type OrderServiceTestSuite struct {
	suite.Suite
	MockDb      *dbMock.MockDbRepo
	MockHandler *handlerMock.MockEventHandler
	Service     *OrderService
}

// before each test
func (suite *OrderServiceTestSuite) SetupTest() {
	t := suite.T()
	ctrl := gomock.NewController(t)
	mockDb := dbMock.NewMockDbRepo(ctrl)
	mockHandler := handlerMock.NewMockEventHandler(ctrl)
	orderRepo := OrderRepository{
		Repo: mockDb,
	}
	publisher := events.EventPublisher{
		Handlers: make(map[string][]events.EventHandler),
	}
	publisher.Subscribe(mockHandler, orderEvents.OrderCreated{}, orderEvents.OrderDeliveryAddressChanged{}, orderEvents.OrderDeliveryAddressChangeFailed{})

	orderService := &OrderService{
		repository: orderRepo,
		publisher:  publisher,
	}

	suite.Service = orderService
	suite.MockDb = mockDb
	suite.MockHandler = mockHandler
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *OrderServiceTestSuite) TestCreateOrder() {
	t := suite.T()
	uuid := uuid.New()

	suite.MockDb.EXPECT().Create(gomock.Any()).DoAndReturn(func(item db.IItem) (string, error) {
		orderItem := item.(dbRecord.OrderItem)
		orderItem.OrderID = uuid

		suite.MockDb.EXPECT().Retrieve(uuid.String()).Return(orderItem, nil)
		return uuid.String(), nil
	})

	payload := &dbRecord.OrderItem{
		IsDispatched: true,
		Address: dbRecord.Address{Address: dto.Address{
			Address1: "742 Evergreen Terrace",
			Address2: "Apt 123",
			City:     "Springfield",
			State:    "IL",
			ZipCode:  "12345",
			Country:  "USA",
		}},
	}
	suite.MockHandler.EXPECT().Notify(orderEvents.OrderCreated{
		OrderId: uuid,
	}).Times(1)
	expectedItem, err := suite.Service.Create(payload)
	if err != nil {
		t.Errorf("Error creating order: %v", err)
	}
	assert.Equal(t, expectedItem.OrderID, uuid)
	assert.Equal(t, expectedItem.IsDispatched, true)

}

func (suite *OrderServiceTestSuite) TestChangeOrderAddress() {
	id := uuid.New()
	updateAddress := dbRecord.Address{Address: dto.Address{
		Address1: "743 Evergreen Terrace",
		Address2: "Apt 124",
		City:     "Springfield",
		State:    "IL",
		ZipCode:  "12345",
		Country:  "USA",
	}}
	record := &dbRecord.OrderItem{
		OrderID:      id,
		IsDispatched: true,
		Address: dbRecord.Address{Address: dto.Address{
			Address1: "742 Evergreen Terrace",
			Address2: "Apt 123",
			City:     "Springfield",
			State:    "IL",
			ZipCode:  "12345",
			Country:  "USA",
		}},
	}
	suite.MockHandler.EXPECT().Notify(orderEvents.OrderDeliveryAddressChangeFailed{
		OrderId: id,
	}).Times(1)
	suite.Service.ChangeAddress(uuid.New(), record, updateAddress)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}
