package order

import (
	"go-ddd-quickstart/internal/pkg/db"
	dbMock "go-ddd-quickstart/internal/pkg/db/mock"
	"go-ddd-quickstart/internal/pkg/dto"
	dbRecord "go-ddd-quickstart/internal/pkg/order/db"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type OrderRepoTestSuite struct {
	suite.Suite
}

// before each test
func (suite *OrderRepoTestSuite) SetupTest() {
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *OrderRepoTestSuite) TestCreate() {
	t := suite.T()
	ctrl := gomock.NewController(t)
	mockDb := dbMock.NewMockDbRepo(ctrl)
	uuid := uuid.New()

	mockDb.EXPECT().Create(gomock.Any()).DoAndReturn(func(item db.IItem) (string, error) {
		orderItem := item.(dbRecord.OrderItem)
		orderItem.OrderID = uuid

		mockDb.EXPECT().Retrieve(uuid.String()).Return(orderItem, nil)
		return uuid.String(), nil
	})

	repo := OrderRepository{
		Repo: mockDb,
	}
	payload := dbRecord.OrderItem{
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
	newItem, err := repo.Create(payload)
	if err != nil {
		t.Errorf("Error creating order: %v", err)
	}
	assert.Equal(t, newItem.OrderID, uuid)
	assert.Equal(t, newItem.IsDispatched, true)

}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOrderRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepoTestSuite))
}
