package campaign_test

import (
	"emailN/internal/domain/campaign"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "My Campaign"
	content = "My Content"
	emails  = []string{"bruno@email.com", "email@email.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	// act - agir
	campaign, _ := campaign.NewCampaign(name, content, emails)

	// assert - assegurar
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}

func Test_NewCampaignIDIsNotNil(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	// act - agir
	campaign, _ := campaign.NewCampaign(name, content, emails)
	assert.NotNil(campaign.ID)
}
func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	now := time.Now().Add(-time.Second)

	// act - agir
	campaign, _ := campaign.NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	// act - agir
	_, err := campaign.NewCampaign("", content, emails)

	assert.Equal("name is required", err.Error())
}
func Test_NewCampaign_MustValidateContent(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	// act - agir
	_, err := campaign.NewCampaign(name, "", emails)

	assert.Equal("content is required", err.Error())
}
func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	// arrange - organizar o código
	assert := assert.New(t)

	// act - agir
	_, err := campaign.NewCampaign(name, content, []string{})

	assert.Equal("contacts are required", err.Error())
}
