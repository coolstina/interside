package interside

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestCustomTagSuite(t *testing.T) {
	suite.Run(t, &CustomTagSuite{})
}

type CustomTagSuite struct {
	suite.Suite
	iMap *Map
}

func (suite *CustomTagSuite) BeforeTest(suiteName, testName string) {
	suite.iMap = NewMap()
}

func (suite *CustomTagSuite) Test_Append() {
	suite.iMap.Append("china", "中国").
		Append("usa", "美国")

	toMap := suite.iMap.ToMap()
	assert.NotNil(suite.T(), toMap)
	assert.Len(suite.T(), toMap, 2)
}
