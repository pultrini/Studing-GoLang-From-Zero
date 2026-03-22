package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getValidNewCampaign() contract.NewCampaign {
	return contract.NewCampaign{
		Name:    "Test y",
		Content: "Body valid",
		Emails:  []string{"test@test.com"},
	}
}

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	input := getValidNewCampaign()

	repo := &repositoryMock{}
	repo.On("Save", mock.Anything).Return(nil)
	service := &Service{Repository: repo}

	id, err := service.Create(input)
	assert.NotNil(id)
	assert.Nil(err)
	repo.AssertExpectations(t)
}

func Test_Create_ValidadeDomainError(t *testing.T) {
	assert := assert.New(t)
	input := getValidNewCampaign()

	input.Name = ""
	localService := Service{}

	_, err := localService.Create(input)
	assert.NotNil(err)
	assert.Equal("name is required with min 5", err.Error())
}

func Test_Create_SaveCampaign(t *testing.T) {
	input := getValidNewCampaign()
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != input.Name ||
			campaign.Content != input.Content ||
			len(campaign.Contacts) != len(input.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	service.Create(input)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)

	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	input := contract.NewCampaign{
		Name:    "Campaign Test",
		Content: "Content valid and long enough",
		Emails:  []string{"test@test.com"},
	}

	localService := Service{
		Repository: repositoryMock,
	}

	_, err := localService.Create(input)

	assert.NotNil(err)
	assert.True(errors.Is(err, internalErrors.ErrorInternal), "Deve retornar ErrorInternal quando o repositório falha")
	repositoryMock.AssertExpectations(t)
}
