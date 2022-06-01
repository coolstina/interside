// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package interside

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestInterSideSuite(t *testing.T) {
	suite.Run(t, &InterSideSuite{})
}

type InterSideSuite struct {
	suite.Suite
	container Container
}

func (suite *InterSideSuite) BeforeTest(suiteName, testName string) {
	suite.container = NewContainer()
}

func (suite *InterSideSuite) Test_Uint_ToInt() {
	suite.container.Append(uint(1), uint(2), uint(3))
	assert.Equal(suite.T(), 3, suite.container.Length())

	actual := suite.container.IntSlice()
	assert.NotNil(suite.T(), actual)
	assert.Len(suite.T(), actual, 3)
}

func (suite *InterSideSuite) Test_UInt_ToString() {
	suite.container.Append(5, 6, 7)
	assert.Equal(suite.T(), 3, suite.container.Length())

	actual := suite.container.StringSlice()
	assert.NotNil(suite.T(), actual)
	assert.Len(suite.T(), actual, 3)
}

func (suite *InterSideSuite) Test_UInt_ToInterface() {
	suite.container.Append(5, 6, 7)
	assert.Equal(suite.T(), 3, suite.container.Length())

	actual := suite.container.InterfaceSlice()
	assert.NotNil(suite.T(), actual)
	assert.Len(suite.T(), actual, 3)
}
