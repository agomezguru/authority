package authority

import (
	"testing"

	configmocks "github.com/goravel/framework/contracts/config/mocks"
	"github.com/stretchr/testify/suite"
)

type AuthorityTestSuite struct {
	suite.Suite
	authority  *Authority
	mockConfig *configmocks.Config
}

func TestAuthorityTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorityTestSuite))
}

func (s *AuthorityTestSuite) SetupTest() {
	s.mockConfig = &configmocks.Config{}
	s.authority = NewAuthority(s.mockConfig)
}

func (s *AuthorityTestSuite) TestWorld() {
	s.mockConfig.On("GetString", "hello.name").Return("Package").Once()

	s.Equal("Welcome To Goravel Package", s.authority.World())
	s.mockConfig.AssertExpectations(s.T())
}
