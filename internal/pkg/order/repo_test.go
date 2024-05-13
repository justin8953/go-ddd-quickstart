package order

import (
	"go-ddd-quickstart/internal/pkg/db"
	dbMock "go-ddd-quickstart/internal/pkg/db/mock"
	"go-ddd-quickstart/internal/pkg/dto"
	dbRecord "go-ddd-quickstart/internal/pkg/order/db"
	"time"

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
	expectedItem, err := repo.Create(payload)
	if err != nil {
		t.Errorf("Error creating order: %v", err)
	}
	assert.Equal(t, expectedItem.OrderID, uuid)
	assert.Equal(t, expectedItem.IsDispatched, true)
}

func (suite *OrderRepoTestSuite) TestUpdate() {
	t := suite.T()
	ctrl := gomock.NewController(t)
	mockDb := dbMock.NewMockDbRepo(ctrl)
	uuid := uuid.New()
	var updatedTimestamp time.Time
	mockDb.EXPECT().Update(uuid.String(), gomock.Any()).DoAndReturn(func(id string, item db.IItem) error {
		orderItem := item.(dbRecord.OrderItem)
		orderItem.OrderID = uuid
		updatedTimestamp = orderItem.UpdatedTimestamp
		mockDb.EXPECT().Retrieve(uuid.String()).Return(orderItem, nil)
		return nil
	})

	repo := OrderRepository{
		Repo: mockDb,
	}
	payload := dbRecord.OrderItem{
		IsDispatched: false,
		Address: dbRecord.Address{Address: dto.Address{
			Address1: "742 Evergreen Terrace",
			Address2: "Apt 123",
			City:     "Springfield",
			State:    "IL",
			ZipCode:  "12345",
			Country:  "USA",
		}},
	}
	expectedItem, err := repo.Update(uuid.String(), payload)
	if err != nil {
		t.Errorf("Error creating order: %v", err)
	}
	assert.Equal(t, expectedItem.OrderID, uuid)
	assert.Equal(t, expectedItem.IsDispatched, false)
	assert.Equal(t, expectedItem.UpdatedTimestamp, updatedTimestamp)
}

func (suite *OrderRepoTestSuite) TestDelete() {
	t := suite.T()
	ctrl := gomock.NewController(t)
	mockDb := dbMock.NewMockDbRepo(ctrl)
	uuid := uuid.New()
	mockDb.EXPECT().Delete(uuid.String()).DoAndReturn(func(id string) error {
		return nil
	})

	repo := OrderRepository{
		Repo: mockDb,
	}

	err := repo.Delete(uuid.String())
	assert.Nil(t, err)
}

func (suite *OrderRepoTestSuite) TestRetrieve() {
	t := suite.T()
	ctrl := gomock.NewController(t)
	mockDb := dbMock.NewMockDbRepo(ctrl)
	uuid := uuid.New()
	actualItem := dbRecord.OrderItem{
		OrderID:      uuid,
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
	mockDb.EXPECT().Retrieve(uuid.String()).Return(actualItem, nil)

	repo := OrderRepository{
		Repo: mockDb,
	}

	expectedItem, err := repo.Retrieve(uuid.String())
	if err != nil {
		t.Errorf("Error creating order: %v", err)
	}
	assert.Equal(t, expectedItem.OrderID, uuid)
	assert.Equal(t, expectedItem.IsDispatched, true)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOrderRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepoTestSuite))
}
