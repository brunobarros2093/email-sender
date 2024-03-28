package campaign

import (
	"emailN/internal/contract"
	"testing"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := contract.NewCampaignDTO{
		Name:    "My Campaign",
		Content: "My Content",
		Emails:  []string{"bruno@email.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service := Service{repositoryMock}

	service.Create(newCampaign)
	repositoryMock.AssertExpectations(t)

}
