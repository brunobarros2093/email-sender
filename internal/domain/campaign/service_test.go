package campaign

import (
	"emailN/internal/contract"
	internalerrors "emailN/internal/internalErrors"
	"errors"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"teste1@test.com"},
	}
	service = Service{}
	fake    = faker.New()
)

func Test_Create_Campaign(t *testing.T) {

	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.NotNil(err)

}
func Test_Create_ValidateDomainError(t *testing.T) {

	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaignDTO{})
	assert.Nil(t, err)
	//assert.Equal("Name is required", err.Error())

}
func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(internalerrors.ErrInternal)
	service.Repository = repositoryMock
	_, err := service.Create(newCampaign)
	assert.True(errors.Is(internalerrors.ErrInternal, err))
}

func Test_NewCampaign_MustvalidateNameMin(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = fake.Lorem().Word()
	_, err := service.Create(newCampaign)
	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_NewCampaign_MustvalidateNameMax(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = fake.Lorem().Text(30)
	_, err := service.Create(newCampaign)
	assert.NotNil(err)
}
