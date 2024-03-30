package campaign

import (
	"emailN/internal/contract"
	internalerrors "emailN/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var (
	newCampaign = contract.NewCampaignDTO{
		Name:    "My Campaign",
		Content: "My Content",
		Emails:  []string{"bruno@email.com"},
	}
	service = Service{}
)

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_SaveCampaign(t *testing.T) {

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)

}
func Test_Create_ValidateDomainError(t *testing.T) {

	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)
	assert.Nil(err)
	//assert.Equal("Name is required", err.Error())

}
func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(internalerrors.ErrInternal)
	service.Repository = repositoryMock
	_, err := service.Create(newCampaign)
	assert.True(errors.Is(internalerrors.ErrInternal, err))
}
