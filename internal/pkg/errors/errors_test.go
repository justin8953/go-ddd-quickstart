package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomErrorTestSuite struct {
	suite.Suite
}

// before each test
func (suite *CustomErrorTestSuite) SetupTest() {
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *CustomErrorTestSuite) TestGeneralError() {
	error := NewEventError("test", errors.New("unavailable"))
	assert.Equal(suite.T(), "test", error.Name())
	assert.Equal(suite.T(), "event name test: err unavailable", error.Error())
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestCustomErrorTestSuite(t *testing.T) {
	suite.Run(t, new(CustomErrorTestSuite))
}
